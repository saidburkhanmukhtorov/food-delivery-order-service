package postgres

import (
	"context"
	"fmt"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type OrderRepo struct {
	db *pgx.Conn
}

func NewOrderRepo(db *pgx.Conn) *OrderRepo { // Now accepts *pgx.Conn
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.Order, error) {
	if req.Order.Id == "" {
		req.Order.Id = uuid.NewString()
	}
	query := `
		INSERT INTO orders (
			id,
			client_id,
			courier_id,
			office_id,
			delivery_latitude,
			delivery_longitude,
			total_price,
			status,
			basket_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		) RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		req.Order.Id,
		req.Order.ClientId,
		req.Order.CourierId,
		req.Order.OfficeId,
		req.Order.DeliveryLatitude,
		req.Order.DeliveryLongitude,
		req.Order.TotalPrice,
		req.Order.Status,
		req.Order.BasketId,
	).Scan(&req.Order.Id)

	if err != nil {
		return nil, err
	}

	return req.Order, nil
}

func (r *OrderRepo) GetOrderByID(ctx context.Context, id string) (*order.Order, error) {
	var (
		orderModel order.Order
	)
	query := `
		SELECT 
			id,
			client_id,
			courier_id,
			office_id,
			delivery_latitude,
			delivery_longitude,
			total_price,
			status,
			basket_id
		FROM orders
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&orderModel.Id,
		&orderModel.ClientId,
		&orderModel.CourierId,
		&orderModel.OfficeId,
		&orderModel.DeliveryLatitude,
		&orderModel.DeliveryLongitude,
		&orderModel.TotalPrice,
		&orderModel.Status,
		&orderModel.BasketId,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &orderModel, nil
}

func (r *OrderRepo) GetOrdersByClient(ctx context.Context, clientID string) (*order.GetOrdersResponse, error) {
	query := `
		SELECT 
			id,
			client_id,
			courier_id,
			office_id,
			delivery_latitude,
			delivery_longitude,
			total_price,
			status,
			basket_id
		FROM orders
		WHERE client_id = $1
	`

	rows, err := r.db.Query(ctx, query, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderList []*order.Order

	for rows.Next() {
		var (
			orderModel order.Order
		)
		err = rows.Scan(
			&orderModel.Id,
			&orderModel.ClientId,
			&orderModel.CourierId,
			&orderModel.OfficeId,
			&orderModel.DeliveryLatitude,
			&orderModel.DeliveryLongitude,
			&orderModel.TotalPrice,
			&orderModel.Status,
			&orderModel.BasketId,
		)
		if err != nil {
			return nil, err
		}
		orderList = append(orderList, &orderModel)
	}

	return &order.GetOrdersResponse{
		Orders: orderList,
	}, nil
}

func (r *OrderRepo) UpdateOrder(ctx context.Context, req *order.UpdateOrderRequest) (*order.Order, error) {
	query := `
		UPDATE orders
		SET 
			client_id = $1,
			courier_id = $2,
			office_id = $3,
			delivery_latitude = $4,
			delivery_longitude = $5,
			total_price = $6,
			status = $7,
			basket_id = $8
		WHERE id = $9
		RETURNING id, client_id, courier_id, office_id, delivery_latitude, delivery_longitude, total_price, status, basket_id
	`

	err := r.db.QueryRow(ctx, query,
		req.Order.ClientId,
		req.Order.CourierId,
		req.Order.OfficeId,
		req.Order.DeliveryLatitude,
		req.Order.DeliveryLongitude,
		req.Order.TotalPrice,
		req.Order.Status,
		req.Order.BasketId,
		req.Order.Id,
	).Scan(
		&req.Order.Id,
		&req.Order.ClientId,
		&req.Order.CourierId,
		&req.Order.OfficeId,
		&req.Order.DeliveryLatitude,
		&req.Order.DeliveryLongitude,
		&req.Order.TotalPrice,
		&req.Order.Status,
		&req.Order.BasketId,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return req.Order, nil
}

func (r *OrderRepo) PatchOrder(ctx context.Context, req *order.PatchOrderRequest) (*order.Order, error) {
	var args []interface{}
	count := 1
	query := `
		UPDATE orders
		SET 
	`

	filter := ""

	if req.ClientId != "" {
		filter += fmt.Sprintf(" client_id = $%d, ", count)
		args = append(args, req.ClientId)
		count++
	}

	if req.CourierId != "" {
		filter += fmt.Sprintf(" courier_id = $%d, ", count)
		args = append(args, req.CourierId)
		count++
	}

	if req.OfficeId != "" {
		filter += fmt.Sprintf(" office_id = $%d, ", count)
		args = append(args, req.OfficeId)
		count++
	}

	if req.DeliveryLatitude != 0 {
		filter += fmt.Sprintf(" delivery_latitude = $%d, ", count)
		args = append(args, req.DeliveryLatitude)
		count++
	}

	if req.DeliveryLongitude != 0 {
		filter += fmt.Sprintf(" delivery_longitude = $%d, ", count)
		args = append(args, req.DeliveryLongitude)
		count++
	}

	if req.TotalPrice != 0 {
		filter += fmt.Sprintf(" total_price = $%d, ", count)
		args = append(args, req.TotalPrice)
		count++
	}

	if req.Status != "" {
		filter += fmt.Sprintf(" status = $%d, ", count)
		args = append(args, req.Status)
		count++
	}

	if req.BasketId != "" {
		filter += fmt.Sprintf(" basket_id = $%d, ", count)
		args = append(args, req.BasketId)
		count++
	}

	if filter == "" {
		return nil, fmt.Errorf("at least one field to update is required")
	}

	filter = filter[:len(filter)-2] // Remove the trailing comma and space
	query += filter + fmt.Sprintf(" WHERE id = $%d RETURNING id, client_id, courier_id, office_id, delivery_latitude, delivery_longitude, total_price, status, basket_id", count)
	args = append(args, req.Id)

	var response order.Order
	err := r.db.QueryRow(ctx, query, args...).Scan(
		&response.Id,
		&response.ClientId,
		&response.CourierId,
		&response.OfficeId,
		&response.DeliveryLatitude,
		&response.DeliveryLongitude,
		&response.TotalPrice,
		&response.Status,
		&response.BasketId,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &response, nil
}

func (r *OrderRepo) DeleteOrder(ctx context.Context, req *order.DeleteOrderRequest) error {
	query := `
		DELETE FROM orders
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

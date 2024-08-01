package postgres

import (
	"context"
	"fmt"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type BasketRepo struct {
	db *pgx.Conn
}

func NewBasketRepo(db *pgx.Conn) *BasketRepo {
	return &BasketRepo{
		db: db,
	}
}

func (r *BasketRepo) CreateBasket(ctx context.Context, req *order.CreateBasketRequest) (*order.Basket, error) {
	if req.Basket.Id == "" {
		req.Basket.Id = uuid.NewString()
	}
	query := `
		INSERT INTO baskets (
			id,
			user_id,
			status
		) VALUES (
			$1, $2, $3
		) RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		req.Basket.Id,
		req.Basket.UserId,
		req.Basket.Status,
	).Scan(&req.Basket.Id)

	if err != nil {
		return nil, err
	}

	return req.Basket, nil
}

func (r *BasketRepo) GetBasketByID(ctx context.Context, id string) (*order.Basket, error) {
	var (
		basketModel order.Basket
	)
	query := `
		SELECT 
			id,
			user_id,
			status
		FROM baskets
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&basketModel.Id,
		&basketModel.UserId,
		&basketModel.Status,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &basketModel, nil
}

func (r *BasketRepo) GetBasketByUser(ctx context.Context, userID string) (*order.Basket, error) {
	var (
		basketModel order.Basket
	)
	query := `
		SELECT 
			id,
			user_id,
			status
		FROM baskets
		WHERE user_id = $1
	`

	err := r.db.QueryRow(ctx, query, userID).Scan(
		&basketModel.Id,
		&basketModel.UserId,
		&basketModel.Status,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &basketModel, nil
}

func (r *BasketRepo) UpdateBasket(ctx context.Context, req *order.UpdateBasketRequest) (*order.Basket, error) {
	query := `
		UPDATE baskets
		SET 
			user_id = $1,
			status = $2
		WHERE id = $3
		RETURNING id, user_id, status
	`

	err := r.db.QueryRow(ctx, query,
		req.Basket.UserId,
		req.Basket.Status,
		req.Basket.Id,
	).Scan(
		&req.Basket.Id,
		&req.Basket.UserId,
		&req.Basket.Status,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return req.Basket, nil
}

func (r *BasketRepo) PatchBasket(ctx context.Context, req *order.PatchBasketRequest) (*order.Basket, error) {
	var args []interface{}
	count := 1
	query := `
		UPDATE baskets
		SET 
	`

	filter := ""

	if req.UserId != "" {
		filter += fmt.Sprintf(" user_id = $%d, ", count)
		args = append(args, req.UserId)
		count++
	}

	if req.Status != "" {
		filter += fmt.Sprintf(" status = $%d, ", count)
		args = append(args, req.Status)
		count++
	}

	if filter == "" {
		return nil, fmt.Errorf("at least one field to update is required")
	}

	filter = filter[:len(filter)-2] // Remove the trailing comma and space
	query += filter + fmt.Sprintf(" WHERE id = $%d RETURNING id, user_id, status", count)
	args = append(args, req.Id)

	var response order.Basket
	err := r.db.QueryRow(ctx, query, args...).Scan(
		&response.Id,
		&response.UserId,
		&response.Status,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &response, nil
}

func (r *BasketRepo) DeleteBasket(ctx context.Context, id string) error {
	query := `
		DELETE FROM baskets
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *BasketRepo) CreateBasketItem(ctx context.Context, req *order.CreateBasketItemRequest) (*order.BasketItem, error) {
	if req.BasketItem.Id == "" {
		req.BasketItem.Id = uuid.NewString()
	}
	query := `
		INSERT INTO basket_items (
			id,
			basket_id,
			product_id,
			quantity
		) VALUES (
			$1, $2, $3, $4
		) RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		req.BasketItem.Id,
		req.BasketItem.BasketId,
		req.BasketItem.ProductId,
		req.BasketItem.Quantity,
	).Scan(&req.BasketItem.Id)

	if err != nil {
		return nil, err
	}

	return req.BasketItem, nil
}

func (r *BasketRepo) GetBasketItems(ctx context.Context, basketID string) (*order.GetBasketItemsResponse, error) {
	query := `
		SELECT 
			id,
			basket_id,
			product_id,
			quantity
		FROM basket_items
		WHERE basket_id = $1
	`

	rows, err := r.db.Query(ctx, query, basketID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var basketItemList []*order.BasketItem

	for rows.Next() {
		var (
			basketItemModel order.BasketItem
		)
		err = rows.Scan(
			&basketItemModel.Id,
			&basketItemModel.BasketId,
			&basketItemModel.ProductId,
			&basketItemModel.Quantity,
		)
		if err != nil {
			return nil, err
		}
		basketItemList = append(basketItemList, &basketItemModel)
	}

	return &order.GetBasketItemsResponse{
		BasketItems: basketItemList,
	}, nil
}

func (r *BasketRepo) UpdateBasketItemQuantity(ctx context.Context, basketItemID string, quantity int32) (*order.BasketItem, error) {
	query := `
		UPDATE basket_items
		SET 
			quantity = $1
		WHERE id = $2
		RETURNING id, basket_id, product_id, quantity
	`

	var response order.BasketItem
	err := r.db.QueryRow(ctx, query, quantity, basketItemID).Scan(
		&response.Id,
		&response.BasketId,
		&response.ProductId,
		&response.Quantity,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &response, nil
}

func (r *BasketRepo) DeleteBasketItem(ctx context.Context, id string) error {
	query := `
		DELETE FROM basket_items
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

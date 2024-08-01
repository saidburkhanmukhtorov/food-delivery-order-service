package postgres

import (
	"context"
	"fmt"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ProductRepo struct {
	db *pgx.Conn
}

func NewProductRepo(db *pgx.Conn) *ProductRepo { // Now accepts *pgx.Conn
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) CreateProduct(ctx context.Context, req *order.CreateProductRequest) (*order.Product, error) {
	if req.Product.Id == "" {
		req.Product.Id = uuid.NewString()
	}
	query := `
		INSERT INTO products (
			id,
			name,
			description,
			price,
			image_url
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		req.Product.Id,
		req.Product.Name,
		req.Product.Description,
		req.Product.Price,
		req.Product.ImageUrl,
	).Scan(&req.Product.Id)

	if err != nil {
		return nil, err
	}

	return req.Product, nil
}

func (r *ProductRepo) GetProductByID(ctx context.Context, id string) (*order.Product, error) {
	var (
		productModel order.Product
	)
	query := `
		SELECT 
			id,
			name,
			description,
			price,
			image_url
		FROM products
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&productModel.Id,
		&productModel.Name,
		&productModel.Description,
		&productModel.Price,
		&productModel.ImageUrl,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &productModel, nil
}
func (r *ProductRepo) GetAllProducts(ctx context.Context, req *order.GetProductsRequest) (*order.GetProductsResponse, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT 
			id,
			name,
			description,
			price,
			image_url
		FROM 
			products
		WHERE 1=1
	`

	filter := ""

	if req.Name != "" {
		filter += fmt.Sprintf(" AND name ILIKE $%d", count)
		args = append(args, "%"+req.Name+"%")
		count++
	}

	if req.Description != "" {
		filter += fmt.Sprintf(" AND description ILIKE $%d", count)
		args = append(args, "%"+req.Description+"%")
		count++
	}

	if req.MinPrice > 0 {
		filter += fmt.Sprintf(" AND price >= $%d", count)
		args = append(args, req.MinPrice)
		count++
	}

	if req.MaxPrice > 0 {
		filter += fmt.Sprintf(" AND price <= $%d", count)
		args = append(args, req.MaxPrice)
		count++
	}

	query += filter

	// Handle invalid page or limit values
	if req.Page <= 0 {
		req.Page = 1 // Default to page 1
	}
	if req.Limit <= 0 {
		req.Limit = 10 // Default to a limit of 10
	}

	// Add LIMIT and OFFSET for pagination using the proto fields
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", count, count+1)
	args = append(args, req.Limit, (req.Page-1)*req.Limit)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productList []*order.Product

	for rows.Next() {
		var (
			productModel order.Product
		)
		err = rows.Scan(
			&productModel.Id,
			&productModel.Name,
			&productModel.Description,
			&productModel.Price,
			&productModel.ImageUrl,
		)
		if err != nil {
			return nil, err
		}
		productList = append(productList, &productModel)
	}

	return &order.GetProductsResponse{
		Products: productList,
	}, nil
}
func (r *ProductRepo) UpdateProduct(ctx context.Context, req *order.UpdateProductRequest) (*order.Product, error) {
	query := `
		UPDATE products
		SET 
			name = $1,
			description = $2,
			price = $3,
			image_url = $4
		WHERE id = $5
		RETURNING id, name, description, price, image_url
	`

	err := r.db.QueryRow(ctx, query,
		req.Product.Name,
		req.Product.Description,
		req.Product.Price,
		req.Product.ImageUrl,
		req.Product.Id,
	).Scan(
		&req.Product.Id,
		&req.Product.Name,
		&req.Product.Description,
		&req.Product.Price,
		&req.Product.ImageUrl,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return req.Product, nil
}

func (r *ProductRepo) PatchProduct(ctx context.Context, req *order.PatchProductRequest) (*order.Product, error) {
	var args []interface{}
	count := 1
	query := `
		UPDATE products
		SET 
	`

	filter := ""

	if req.Name != "" {
		filter += fmt.Sprintf(" name = $%d, ", count)
		args = append(args, req.Name)
		count++
	}

	if req.Description != "" {
		filter += fmt.Sprintf(" description = $%d, ", count)
		args = append(args, req.Description)
		count++
	}

	if req.Price != 0 {
		filter += fmt.Sprintf(" price = $%d, ", count)
		args = append(args, req.Price)
		count++
	}

	if req.ImageUrl != "" {
		filter += fmt.Sprintf(" image_url = $%d, ", count)
		args = append(args, req.ImageUrl)
		count++
	}

	if filter == "" {
		return nil, fmt.Errorf("at least one field to update is required")
	}

	filter = filter[:len(filter)-2] // Remove the trailing comma and space
	query += filter + fmt.Sprintf(" WHERE id = $%d RETURNING id, name, description, price, image_url", count)
	args = append(args, req.Id)

	var response order.Product
	err := r.db.QueryRow(ctx, query, args...).Scan(
		&response.Id,
		&response.Name,
		&response.Description,
		&response.Price,
		&response.ImageUrl,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &response, nil
}

func (r *ProductRepo) DeleteProduct(ctx context.Context, id string) error {
	query := `
		DELETE FROM products
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

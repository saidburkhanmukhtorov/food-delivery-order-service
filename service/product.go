package service

import (
	"context"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/food-delivery/food-delivery-order-service/storage"
)

// ProductServiceI defines the interface for the product service.
type ProductServiceI interface {
	CreateProduct(ctx context.Context, req *order.CreateProductRequest) (*order.Product, error)
	GetProductByID(ctx context.Context, req *order.ProductRequest) (*order.Product, error)
	GetAllProducts(ctx context.Context, req *order.GetProductsRequest) (*order.GetProductsResponse, error)
	UpdateProduct(ctx context.Context, req *order.UpdateProductRequest) (*order.Product, error)
	PatchProduct(ctx context.Context, req *order.PatchProductRequest) (*order.Product, error)
	DeleteProduct(ctx context.Context, req *order.DeleteProductRequest) (*order.DeleteProductRes, error)
}

// ProductService implements the ProductServiceI interface.
type ProductService struct {
	storage storage.StorageI
	order.UnimplementedProductServiceServer
}

// NewProductService creates a new ProductService instance.
func NewProductService(storage storage.StorageI) *ProductService {
	return &ProductService{
		storage: storage,
	}
}

// CreateProduct creates a new product.
func (s *ProductService) CreateProduct(ctx context.Context, req *order.CreateProductRequest) (*order.Product, error) {
	product, err := s.storage.Product().CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetProductByID retrieves a product by its ID.
func (s *ProductService) GetProduct(ctx context.Context, req *order.ProductRequest) (*order.Product, error) {
	product, err := s.storage.Product().GetProductByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetAllProducts retrieves a list of products with optional filtering and pagination.
func (s *ProductService) GetAllProducts(ctx context.Context, req *order.GetProductsRequest) (*order.GetProductsResponse, error) {
	products, err := s.storage.Product().GetAllProducts(ctx, req)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct updates an existing product.
func (s *ProductService) UpdateProduct(ctx context.Context, req *order.UpdateProductRequest) (*order.Product, error) {
	product, err := s.storage.Product().UpdateProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// PatchProduct partially updates an existing product.
func (s *ProductService) PatchProduct(ctx context.Context, req *order.PatchProductRequest) (*order.Product, error) {
	product, err := s.storage.Product().PatchProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// DeleteProduct deletes a product.
func (s *ProductService) DeleteProduct(ctx context.Context, req *order.DeleteProductRequest) (*order.DeleteProductRes, error) {
	err := s.storage.Product().DeleteProduct(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &order.DeleteProductRes{Message: "Product deleted successfully"}, nil
}

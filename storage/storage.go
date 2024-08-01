package storage

import (
	"context"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
)

// StorageI defines the interface for interacting with the storage layer.
type StorageI interface {
	Product() ProductI
	Office() OfficeI
	Basket() BasketI
	Order() OrderI
}

// ProductI defines methods for interacting with product data.
type ProductI interface {
	CreateProduct(ctx context.Context, req *order.CreateProductRequest) (*order.Product, error)
	GetProductByID(ctx context.Context, id string) (*order.Product, error)
	GetAllProducts(ctx context.Context, req *order.GetProductsRequest) (*order.GetProductsResponse, error)
	UpdateProduct(ctx context.Context, req *order.UpdateProductRequest) (*order.Product, error)
	PatchProduct(ctx context.Context, req *order.PatchProductRequest) (*order.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

// OfficeI defines methods for interacting with office data.
type OfficeI interface {
	CreateOffice(ctx context.Context, req *order.CreateOfficeRequest) (*order.Office, error)
	GetOfficeByID(ctx context.Context, id string) (*order.Office, error)
	GetAllOffices(ctx context.Context, req *order.GetOfficesRequest) (*order.GetOfficesResponse, error)
	UpdateOffice(ctx context.Context, req *order.UpdateOfficeRequest) (*order.Office, error)
	PatchOffice(ctx context.Context, req *order.PatchOfficeRequest) (*order.Office, error)
	DeleteOffice(ctx context.Context, id string) error
}

// BasketI defines methods for interacting with basket data.
type BasketI interface {
	CreateBasket(ctx context.Context, req *order.CreateBasketRequest) (*order.Basket, error)
	GetBasketByID(ctx context.Context, id string) (*order.Basket, error)
	GetBasketByUser(ctx context.Context, userID string) (*order.Basket, error)
	UpdateBasket(ctx context.Context, req *order.UpdateBasketRequest) (*order.Basket, error)
	PatchBasket(ctx context.Context, req *order.PatchBasketRequest) (*order.Basket, error)
	DeleteBasket(ctx context.Context, id string) error
	CreateBasketItem(ctx context.Context, req *order.CreateBasketItemRequest) (*order.BasketItem, error)
	GetBasketItems(ctx context.Context, basketID string) (*order.GetBasketItemsResponse, error)
	UpdateBasketItemQuantity(ctx context.Context, basketItemID string, quantity int32) (*order.BasketItem, error)
	DeleteBasketItem(ctx context.Context, id string) error
}

// OrderI defines methods for interacting with order data.
type OrderI interface {
	CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.Order, error)
	GetOrderByID(ctx context.Context, id string) (*order.Order, error)
	GetOrdersByClient(ctx context.Context, clientID string) (*order.GetOrdersResponse, error)
	UpdateOrder(ctx context.Context, req *order.UpdateOrderRequest) (*order.Order, error)
	PatchOrder(ctx context.Context, req *order.PatchOrderRequest) (*order.Order, error)
	DeleteOrder(ctx context.Context, req *order.DeleteOrderRequest) error
}

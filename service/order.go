package service

import (
	"context"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/food-delivery/food-delivery-order-service/storage"
)

// OrderServiceI defines the interface for the order service.
type OrderServiceI interface {
	CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.Order, error)
	GetOrderByID(ctx context.Context, req *order.OrderRequest) (*order.Order, error)
	GetOrdersByClient(ctx context.Context, req *order.GetOrdersByClientRequest) (*order.GetOrdersResponse, error)
	UpdateOrder(ctx context.Context, req *order.UpdateOrderRequest) (*order.Order, error)
	PatchOrder(ctx context.Context, req *order.PatchOrderRequest) (*order.Order, error)
	DeleteOrder(ctx context.Context, req *order.DeleteOrderRequest) (*order.DeleteOrderRes, error)
}

// OrderService implements the OrderServiceI interface.
type OrderService struct {
	storage storage.StorageI
	order.UnimplementedOrderServiceServer
}

// NewOrderService creates a new OrderService instance.
func NewOrderService(storage storage.StorageI) *OrderService {
	return &OrderService{
		storage: storage,
	}
}

// CreateOrder creates a new order.
func (s *OrderService) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.Order, error) {
	order, err := s.storage.Order().CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetOrderByID retrieves an order by its ID.
func (s *OrderService) GetOrder(ctx context.Context, req *order.OrderRequest) (*order.Order, error) {
	order, err := s.storage.Order().GetOrderByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetOrdersByClient retrieves orders by client ID.
func (s *OrderService) GetOrdersByClient(ctx context.Context, req *order.GetOrdersByClientRequest) (*order.GetOrdersResponse, error) {
	orders, err := s.storage.Order().GetOrdersByClient(ctx, req.ClientId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrder updates an existing order.
func (s *OrderService) UpdateOrder(ctx context.Context, req *order.UpdateOrderRequest) (*order.Order, error) {
	order, err := s.storage.Order().UpdateOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// PatchOrder partially updates an existing order.
func (s *OrderService) PatchOrder(ctx context.Context, req *order.PatchOrderRequest) (*order.Order, error) {
	order, err := s.storage.Order().PatchOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// DeleteOrder deletes an order.
func (s *OrderService) DeleteOrder(ctx context.Context, req *order.DeleteOrderRequest) (*order.DeleteOrderRes, error) {
	err := s.storage.Order().DeleteOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return &order.DeleteOrderRes{Message: "Order deleted successfully"}, nil
}

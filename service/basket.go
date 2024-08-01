package service

import (
	"context"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/food-delivery/food-delivery-order-service/storage"
)

// BasketServiceI defines the interface for the basket service.
type BasketServiceI interface {
	CreateBasket(ctx context.Context, req *order.CreateBasketRequest) (*order.Basket, error)
	GetBasketByID(ctx context.Context, req *order.BasketRequest) (*order.Basket, error)
	GetBasketByUser(ctx context.Context, req *order.GetBasketByUserRequest) (*order.Basket, error)
	UpdateBasket(ctx context.Context, req *order.UpdateBasketRequest) (*order.Basket, error)
	PatchBasket(ctx context.Context, req *order.PatchBasketRequest) (*order.Basket, error)
	DeleteBasket(ctx context.Context, req *order.DeleteBasketRequest) (*order.DeleteBasketRes, error)
	CreateBasketItem(ctx context.Context, req *order.CreateBasketItemRequest) (*order.BasketItem, error)
	GetBasketItems(ctx context.Context, req *order.GetBasketItemsRequest) (*order.GetBasketItemsResponse, error)
	UpdateBasketItemQuantity(ctx context.Context, req *order.UpdateBasketItemQuantityRequest) (*order.BasketItem, error)
	DeleteBasketItem(ctx context.Context, req *order.DeleteBasketItemRequest) (*order.DeleteBasketItemRes, error)
}

// BasketService implements the BasketServiceI interface.
type BasketService struct {
	storage storage.StorageI
	order.UnimplementedBasketServiceServer
}

// NewBasketService creates a new BasketService instance.
func NewBasketService(storage storage.StorageI) *BasketService {
	return &BasketService{
		storage: storage,
	}
}

// CreateBasket creates a new basket.
func (s *BasketService) CreateBasket(ctx context.Context, req *order.CreateBasketRequest) (*order.Basket, error) {
	basket, err := s.storage.Basket().CreateBasket(ctx, req)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

// GetBasketByID retrieves a basket by its ID.
func (s *BasketService) GetBasket(ctx context.Context, req *order.BasketRequest) (*order.Basket, error) {
	basket, err := s.storage.Basket().GetBasketByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

// GetBasketByUser retrieves a basket by the user's ID.
func (s *BasketService) GetBasketByUser(ctx context.Context, req *order.GetBasketByUserRequest) (*order.Basket, error) {
	basket, err := s.storage.Basket().GetBasketByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

// UpdateBasket updates an existing basket.
func (s *BasketService) UpdateBasket(ctx context.Context, req *order.UpdateBasketRequest) (*order.Basket, error) {
	basket, err := s.storage.Basket().UpdateBasket(ctx, req)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

// PatchBasket partially updates an existing basket.
func (s *BasketService) PatchBasket(ctx context.Context, req *order.PatchBasketRequest) (*order.Basket, error) {
	basket, err := s.storage.Basket().PatchBasket(ctx, req)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

// DeleteBasket deletes a basket.
func (s *BasketService) DeleteBasket(ctx context.Context, req *order.DeleteBasketRequest) (*order.DeleteBasketRes, error) {
	err := s.storage.Basket().DeleteBasket(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &order.DeleteBasketRes{Message: "Basket deleted successfully"}, nil
}

// CreateBasketItem creates a new basket item.
func (s *BasketService) CreateBasketItem(ctx context.Context, req *order.CreateBasketItemRequest) (*order.BasketItem, error) {
	basketItem, err := s.storage.Basket().CreateBasketItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return basketItem, nil
}

// GetBasketItems retrieves all items in a basket.
func (s *BasketService) GetBasketItems(ctx context.Context, req *order.GetBasketItemsRequest) (*order.GetBasketItemsResponse, error) {
	basketItems, err := s.storage.Basket().GetBasketItems(ctx, req.BasketId)
	if err != nil {
		return nil, err
	}
	return basketItems, nil
}

// UpdateBasketItemQuantity updates the quantity of a product in a basket item.
func (s *BasketService) UpdateBasketItemQuantity(ctx context.Context, req *order.UpdateBasketItemQuantityRequest) (*order.BasketItem, error) {
	basketItem, err := s.storage.Basket().UpdateBasketItemQuantity(ctx, req.Id, req.Quantity)
	if err != nil {
		return nil, err
	}
	return basketItem, nil
}

// DeleteBasketItem deletes a basket item.
func (s *BasketService) DeleteBasketItem(ctx context.Context, req *order.DeleteBasketItemRequest) (*order.DeleteBasketItemRes, error) {
	err := s.storage.Basket().DeleteBasketItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &order.DeleteBasketItemRes{Message: "Basket item deleted successfully"}, nil
}

package service

import (
	"context"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/food-delivery/food-delivery-order-service/storage"
)

// OfficeServiceI defines the interface for the office service.
type OfficeServiceI interface {
	CreateOffice(ctx context.Context, req *order.CreateOfficeRequest) (*order.Office, error)
	GetOfficeByID(ctx context.Context, req *order.OfficeRequest) (*order.Office, error)
	GetAllOffices(ctx context.Context, req *order.GetOfficesRequest) (*order.GetOfficesResponse, error)
	UpdateOffice(ctx context.Context, req *order.UpdateOfficeRequest) (*order.Office, error)
	PatchOffice(ctx context.Context, req *order.PatchOfficeRequest) (*order.Office, error)
	DeleteOffice(ctx context.Context, req *order.DeleteOfficeRequest) (*order.DeleteOfficeRes, error)
}

// OfficeService implements the OfficeServiceI interface.
type OfficeService struct {
	storage storage.StorageI
	order.UnimplementedOfficeServiceServer
}

// NewOfficeService creates a new OfficeService instance.
func NewOfficeService(storage storage.StorageI) *OfficeService {
	return &OfficeService{
		storage: storage,
	}
}

// CreateOffice creates a new office.
func (s *OfficeService) CreateOffice(ctx context.Context, req *order.CreateOfficeRequest) (*order.Office, error) {
	office, err := s.storage.Office().CreateOffice(ctx, req)
	if err != nil {
		return nil, err
	}
	return office, nil
}

// GetOfficeByID retrieves an office by its ID.
func (s *OfficeService) GetOffice(ctx context.Context, req *order.OfficeRequest) (*order.Office, error) {
	office, err := s.storage.Office().GetOfficeByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return office, nil
}

// GetAllOffices retrieves a list of offices with optional filtering and pagination.
func (s *OfficeService) GetAllOffices(ctx context.Context, req *order.GetOfficesRequest) (*order.GetOfficesResponse, error) {
	offices, err := s.storage.Office().GetAllOffices(ctx, req)
	if err != nil {
		return nil, err
	}
	return offices, nil
}

// UpdateOffice updates an existing office.
func (s *OfficeService) UpdateOffice(ctx context.Context, req *order.UpdateOfficeRequest) (*order.Office, error) {
	office, err := s.storage.Office().UpdateOffice(ctx, req)
	if err != nil {
		return nil, err
	}
	return office, nil
}

// PatchOffice partially updates an existing office.
func (s *OfficeService) PatchOffice(ctx context.Context, req *order.PatchOfficeRequest) (*order.Office, error) {
	office, err := s.storage.Office().PatchOffice(ctx, req)
	if err != nil {
		return nil, err
	}
	return office, nil
}

// DeleteOffice deletes an office.
func (s *OfficeService) DeleteOffice(ctx context.Context, req *order.DeleteOfficeRequest) (*order.DeleteOfficeRes, error) {
	err := s.storage.Office().DeleteOffice(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &order.DeleteOfficeRes{Message: "Office deleted successfully"}, nil
}

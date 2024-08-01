package main

import (
	"log/slog"
	"net"

	"github.com/food-delivery/food-delivery-order-service/config"
	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/food-delivery/food-delivery-order-service/service"
	"github.com/food-delivery/food-delivery-order-service/storage/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to PostgreSQL database
	storageInstance, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		slog.Error("Failed to connect to database:", err)
		return
	}

	// Create service instances
	orderService := service.NewOrderService(storageInstance)
	productService := service.NewProductService(storageInstance)
	officeService := service.NewOfficeService(storageInstance)
	basketService := service.NewBasketService(storageInstance)

	// Create gRPC server
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		slog.Error("Failed to listen on port:" + err.Error())
		return
	}

	s := grpc.NewServer()

	// Register services
	order.RegisterOrderServiceServer(s, orderService)
	order.RegisterProductServiceServer(s, productService)
	order.RegisterOfficeServiceServer(s, officeService)
	order.RegisterBasketServiceServer(s, basketService)

	// Register reflection service (for introspection)
	reflection.Register(s)

	// Start gRPC server
	slog.Info("Starting gRPC server on port:8082")
	if err := s.Serve(lis); err != nil {
		slog.Error("Failed to serve gRPC server:", err)
	}
}

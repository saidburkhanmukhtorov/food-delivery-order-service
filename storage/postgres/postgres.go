package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/food-delivery/food-delivery-order-service/config"
	"github.com/food-delivery/food-delivery-order-service/storage"
	"github.com/jackc/pgx/v5"
)

// Storage implements the storage.StorageI interface for PostgreSQL.
type Storage struct {
	db       *pgx.Conn
	ProductS storage.ProductI
	OfficeS  storage.OfficeI
	BasketS  storage.BasketI
	OrderS   storage.OrderI
}

// NewPostgresStorage creates a new PostgreSQL storage instance.
func NewPostgresStorage(cfg config.Config) (storage.StorageI, error) {
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	db, err := pgx.Connect(context.Background(), dbCon)
	if err != nil {
		slog.Warn("Unable to connect to database:", err)
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil {
		slog.Warn("Unable to ping database:", err)
		return nil, err
	}

	return &Storage{
		db:       db,
		ProductS: NewProductRepo(db),
		OfficeS:  NewOfficeRepo(db),
		BasketS:  NewBasketRepo(db),
		OrderS:   NewOrderRepo(db),
	}, nil
}

// Product returns the ProductI implementation for PostgreSQL.
func (s *Storage) Product() storage.ProductI {
	return s.ProductS
}

// Office returns the OfficeI implementation for PostgreSQL.
func (s *Storage) Office() storage.OfficeI {
	return s.OfficeS
}

// Basket returns the BasketI implementation for PostgreSQL.
func (s *Storage) Basket() storage.BasketI {
	return s.BasketS
}

// Order returns the OrderI implementation for PostgreSQL.
func (s *Storage) Order() storage.OrderI {
	return s.OrderS
}

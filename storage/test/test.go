package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
)

func createDBConnection(t *testing.T) *pgx.Conn {
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad",
		"root",
		"localhost",
		5432,
		"postgres",
	)

	// Connecting to postgres
	db, err := pgx.Connect(context.Background(), dbCon)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	return db
}

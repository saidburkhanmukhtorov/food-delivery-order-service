package test

import (
	"context"
	"testing"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/food-delivery/food-delivery-order-service/storage/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestOfficeRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	officeRepo := postgres.NewOfficeRepo(db)

	t.Run("CreateOffice", func(t *testing.T) {
		req := &order.CreateOfficeRequest{
			Office: &order.Office{
				Name:      "Office A",
				Address:   "123 Main St",
				Latitude:  37.7749,
				Longitude: -122.4194,
			},
		}

		response, err := officeRepo.CreateOffice(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		// Cleanup
		defer deleteOffice(t, db, response.Id)
	})

	t.Run("GetOfficeByID", func(t *testing.T) {
		req := &order.CreateOfficeRequest{
			Office: &order.Office{
				Name:      "Office B",
				Address:   "456 Elm St",
				Latitude:  34.0522,
				Longitude: -118.2437,
			},
		}

		response, err := officeRepo.CreateOffice(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		office, err := officeRepo.GetOfficeByID(context.Background(), response.Id)
		assert.NoError(t, err)
		assert.NotNil(t, office)
		assert.Equal(t, response.Id, office.Id)

		// Cleanup
		defer deleteOffice(t, db, response.Id)
	})

	t.Run("GetAllOffices", func(t *testing.T) {
		// Create a few test offices
		officesToCreate := []*order.Office{
			{
				Name:      "Office C",
				Address:   "789 Oak St",
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			{
				Name:      "Office D",
				Address:   "1011 Pine St",
				Latitude:  33.4484,
				Longitude: -112.0740,
			},
		}

		for _, office := range officesToCreate {
			_, err := officeRepo.CreateOffice(context.Background(), &order.CreateOfficeRequest{Office: office})
			assert.NoError(t, err)
		}

		// Test GetAllOffices with no filters
		response, err := officeRepo.GetAllOffices(context.Background(), &order.GetOfficesRequest{})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(response.Offices), 2) // At least 2 offices should be returned

		// Test GetAllOffices with name filter
		response, err = officeRepo.GetAllOffices(context.Background(), &order.GetOfficesRequest{Name: "Office D"})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(response.Offices), 1) // At least 1 office with "Office A" in the name

		// Test GetAllOffices with address filter
		response, err = officeRepo.GetAllOffices(context.Background(), &order.GetOfficesRequest{Address: "789 Oak St"})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(response.Offices), 1) // At least 1 office with "123 Main St" in the address

		// Cleanup
		for _, office := range officesToCreate {
			defer deleteOffice(t, db, office.Id)
		}
	})

	t.Run("UpdateOffice", func(t *testing.T) {
		req := &order.CreateOfficeRequest{
			Office: &order.Office{
				Name:      "Office E",
				Address:   "567 Maple St",
				Latitude:  41.8781,
				Longitude: -87.6298,
			},
		}

		response, err := officeRepo.CreateOffice(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		// Update the office
		updateReq := &order.UpdateOfficeRequest{
			Office: &order.Office{
				Id:        response.Id,
				Name:      "Updated Office E",
				Address:   "890 Birch St",
				Latitude:  41.8781,
				Longitude: -87.6298,
			},
		}
		updatedOffice, err := officeRepo.UpdateOffice(context.Background(), updateReq)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Office E", updatedOffice.Name)
		assert.Equal(t, "890 Birch St", updatedOffice.Address)

		// Cleanup
		defer deleteOffice(t, db, response.Id)
	})

	t.Run("PatchOffice", func(t *testing.T) {
		req := &order.CreateOfficeRequest{
			Office: &order.Office{
				Name:      "Office F",
				Address:   "234 Pine St",
				Latitude:  34.0522,
				Longitude: -118.2437,
			},
		}

		response, err := officeRepo.CreateOffice(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		// Patch the office
		patchReq := &order.PatchOfficeRequest{
			Id:      response.Id,
			Name:    "Patched Office F",
			Address: "987 Cedar St",
		}
		patchedOffice, err := officeRepo.PatchOffice(context.Background(), patchReq)
		assert.NoError(t, err)
		assert.Equal(t, "Patched Office F", patchedOffice.Name)
		assert.Equal(t, "987 Cedar St", patchedOffice.Address)

		// Cleanup
		defer deleteOffice(t, db, response.Id)
	})

	t.Run("DeleteOffice", func(t *testing.T) {
		req := &order.CreateOfficeRequest{
			Office: &order.Office{
				Name:      "Office G",
				Address:   "123 Oak St",
				Latitude:  33.4484,
				Longitude: -112.0740,
			},
		}

		response, err := officeRepo.CreateOffice(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		err = officeRepo.DeleteOffice(context.Background(), response.Id)
		assert.NoError(t, err)

		_, err = officeRepo.GetOfficeByID(context.Background(), response.Id)
		assert.ErrorIs(t, err, pgx.ErrNoRows) // Office should not be found
	})
}

func deleteOffice(t *testing.T, db *pgx.Conn, officeID string) {
	_, err := db.Exec(context.Background(), "DELETE FROM offices WHERE id = $1", officeID)
	if err != nil {
		t.Fatalf("Error deleting office: %v", err)
	}
}

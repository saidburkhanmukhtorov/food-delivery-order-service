package test

import (
	"context"
	"testing"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/food-delivery/food-delivery-order-service/storage/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestBasketRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	basketRepo := postgres.NewBasketRepo(db)
	productRepo := postgres.NewProductRepo(db) // Add productRepo for creating products

	t.Run("CreateBasket", func(t *testing.T) {
		req := &order.CreateBasketRequest{
			Basket: &order.Basket{
				UserId: uuid.New().String(),
				Status: "OPEN",
			},
		}

		response, err := basketRepo.CreateBasket(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		// Cleanup
		defer deleteBasket(t, db, response.Id)
	})

	t.Run("GetBasketByID", func(t *testing.T) {
		req := &order.CreateBasketRequest{
			Basket: &order.Basket{
				UserId: uuid.New().String(),
				Status: "OPEN",
			},
		}

		response, err := basketRepo.CreateBasket(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		basket, err := basketRepo.GetBasketByID(context.Background(), response.Id)
		assert.NoError(t, err)
		assert.NotNil(t, basket)
		assert.Equal(t, response.Id, basket.Id)

		// Cleanup
		defer deleteBasket(t, db, response.Id)
	})

	t.Run("GetBasketByUser", func(t *testing.T) {
		userID := uuid.New().String()

		// Create a basket for the user
		_, err := basketRepo.CreateBasket(context.Background(), &order.CreateBasketRequest{
			Basket: &order.Basket{
				UserId: userID,
				Status: "OPEN",
			},
		})
		assert.NoError(t, err)

		// Get the basket by user ID
		basket, err := basketRepo.GetBasketByUser(context.Background(), userID)
		assert.NoError(t, err)
		assert.NotNil(t, basket)
		assert.Equal(t, userID, basket.UserId)

		// Cleanup
		defer deleteBasket(t, db, basket.Id)
	})

	t.Run("UpdateBasket", func(t *testing.T) {
		createReq := &order.CreateBasketRequest{
			Basket: &order.Basket{
				UserId: uuid.New().String(),
				Status: "OPEN",
			},
		}

		response, err := basketRepo.CreateBasket(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		// Update the basket
		updateReq := &order.UpdateBasketRequest{
			Basket: &order.Basket{
				Id:     response.Id,
				Status: "CLOSED",
			},
		}
		updatedBasket, err := basketRepo.UpdateBasket(context.Background(), updateReq)
		assert.NoError(t, err)
		assert.Equal(t, "CLOSED", updatedBasket.Status)

		// Cleanup
		defer deleteBasket(t, db, response.Id)
	})

	t.Run("PatchBasket", func(t *testing.T) {
		createReq := &order.CreateBasketRequest{
			Basket: &order.Basket{
				UserId: uuid.New().String(),
				Status: "OPEN",
			},
		}

		response, err := basketRepo.CreateBasket(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		// Patch the basket
		patchReq := &order.PatchBasketRequest{
			Id:     response.Id,
			Status: "COMPLETED",
		}
		patchedBasket, err := basketRepo.PatchBasket(context.Background(), patchReq)
		assert.NoError(t, err)
		assert.Equal(t, "COMPLETED", patchedBasket.Status)

		// Cleanup
		defer deleteBasket(t, db, response.Id)
	})

	t.Run("DeleteBasket", func(t *testing.T) {
		req := &order.CreateBasketRequest{
			Basket: &order.Basket{
				UserId: uuid.New().String(),
				Status: "OPEN",
			},
		}

		response, err := basketRepo.CreateBasket(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		err = basketRepo.DeleteBasket(context.Background(), response.Id)
		assert.NoError(t, err)

		_, err = basketRepo.GetBasketByID(context.Background(), response.Id)
		assert.ErrorIs(t, err, pgx.ErrNoRows) // Basket should not be found
	})

	t.Run("CreateBasketItem", func(t *testing.T) {
		// 1. Create a test basket
		createBasketReq := &order.CreateBasketRequest{
			Basket: &order.Basket{
				UserId: uuid.New().String(),
				Status: "OPEN",
			},
		}
		basket, err := basketRepo.CreateBasket(context.Background(), createBasketReq)
		assert.NoError(t, err)
		assert.NotEmpty(t, basket.Id)

		// 2. Create a test product
		createProductReq := &order.CreateProductRequest{
			Product: &order.Product{
				Name:        "Test Product",
				Description: "This is a test product",
				Price:       10.99,
				ImageUrl:    "https://example.com/test-product.jpg",
			},
		}
		product, err := productRepo.CreateProduct(context.Background(), createProductReq)
		assert.NoError(t, err)
		assert.NotEmpty(t, product.Id)

		// 3. Create a test basket item using the basket ID and product ID
		req := &order.CreateBasketItemRequest{
			BasketItem: &order.BasketItem{
				BasketId:  basket.Id,
				ProductId: product.Id,
				Quantity:  3,
			},
		}

		response, err := basketRepo.CreateBasketItem(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		// Cleanup
		defer deleteBasketItem(t, db, response.Id)
		defer deleteBasket(t, db, basket.Id)
		defer deleteProduct(t, db, product.Id)
	})
	t.Run("GetBasketItems", func(t *testing.T) {
		// 1. Create a test basket
		createBasketReq := &order.CreateBasketRequest{
			Basket: &order.Basket{
				UserId: uuid.New().String(),
				Status: "OPEN",
			},
		}
		basket, err := basketRepo.CreateBasket(context.Background(), createBasketReq)
		assert.NoError(t, err)
		assert.NotEmpty(t, basket.Id)

		// 2. Create a few test products
		productsToCreate := []*order.Product{
			{
				Name:        "Test Product 1",
				Description: "This is a test product",
				Price:       10.99,
				ImageUrl:    "https://example.com/test-product1.jpg",
			},
			{
				Name:        "Test Product 2",
				Description: "This is another test product",
				Price:       15.99,
				ImageUrl:    "https://example.com/test-product2.jpg",
			},
		}
		for _, product := range productsToCreate {
			_, err := productRepo.CreateProduct(context.Background(), &order.CreateProductRequest{Product: product})
			assert.NoError(t, err)
		}

		// 3. Create a few test basket items using the created products
		basketItemsToCreate := []*order.BasketItem{
			{
				BasketId:  basket.Id,
				ProductId: productsToCreate[0].Id, // Use the ID of the first created product
				Quantity:  2,
			},
			{
				BasketId:  basket.Id,
				ProductId: productsToCreate[1].Id, // Use the ID of the second created product
				Quantity:  1,
			},
		}

		for _, basketItem := range basketItemsToCreate {
			_, err := basketRepo.CreateBasketItem(context.Background(), &order.CreateBasketItemRequest{BasketItem: basketItem})
			assert.NoError(t, err)
		}

		// Test GetBasketItems
		response, err := basketRepo.GetBasketItems(context.Background(), basket.Id)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(response.BasketItems), 2) // At least 2 basket items should be returned

		// Cleanup
		for _, basketItem := range basketItemsToCreate {
			defer deleteBasketItem(t, db, basketItem.Id)
		}
		for _, product := range productsToCreate {
			defer deleteProduct(t, db, product.Id)
		}
		defer deleteBasket(t, db, basket.Id)
	})
	t.Run("UpdateBasketItemQuantity", func(t *testing.T) {
		// 1. Create a test basket
		createBasketReq := &order.CreateBasketRequest{
			Basket: &order.Basket{
				UserId: uuid.New().String(),
				Status: "OPEN",
			},
		}
		basket, err := basketRepo.CreateBasket(context.Background(), createBasketReq)
		assert.NoError(t, err)
		assert.NotEmpty(t, basket.Id)

		// 2. Create a test product
		createProductReq := &order.CreateProductRequest{
			Product: &order.Product{
				Name:        "Test Product",
				Description: "This is a test product",
				Price:       10.99,
				ImageUrl:    "https://example.com/test-product.jpg",
			},
		}
		product, err := productRepo.CreateProduct(context.Background(), createProductReq)
		assert.NoError(t, err)
		assert.NotEmpty(t, product.Id)

		// 3. Create a test basket item
		basketItem, err := basketRepo.CreateBasketItem(context.Background(), &order.CreateBasketItemRequest{
			BasketItem: &order.BasketItem{
				BasketId:  basket.Id,
				ProductId: product.Id,
				Quantity:  2,
			},
		})
		assert.NoError(t, err)

		// Update the basket item quantity
		newQuantity := int32(5)
		updatedBasketItem, err := basketRepo.UpdateBasketItemQuantity(context.Background(), basketItem.Id, newQuantity)
		assert.NoError(t, err)
		assert.Equal(t, newQuantity, updatedBasketItem.Quantity)

		// Cleanup
		defer deleteBasketItem(t, db, basketItem.Id)
		defer deleteBasket(t, db, basket.Id)
		defer deleteProduct(t, db, product.Id)
	})

	t.Run("DeleteBasketItem", func(t *testing.T) {
		// 1. Create a test basket
		createBasketReq := &order.CreateBasketRequest{
			Basket: &order.Basket{
				UserId: uuid.New().String(),
				Status: "OPEN",
			},
		}
		basket, err := basketRepo.CreateBasket(context.Background(), createBasketReq)
		assert.NoError(t, err)
		assert.NotEmpty(t, basket.Id)

		// 2. Create a test product
		createProductReq := &order.CreateProductRequest{
			Product: &order.Product{
				Name:        "Test Product",
				Description: "This is a test product",
				Price:       10.99,
				ImageUrl:    "https://example.com/test-product.jpg",
			},
		}
		product, err := productRepo.CreateProduct(context.Background(), createProductReq)
		assert.NoError(t, err)
		assert.NotEmpty(t, product.Id)

		// 3. Create a test basket item
		basketItem, err := basketRepo.CreateBasketItem(context.Background(), &order.CreateBasketItemRequest{
			BasketItem: &order.BasketItem{
				BasketId:  basket.Id,
				ProductId: product.Id,
				Quantity:  2,
			},
		})
		assert.NoError(t, err)

		// Delete the basket item
		err = basketRepo.DeleteBasketItem(context.Background(), basketItem.Id)
		assert.NoError(t, err)

		// Cleanup
		defer deleteBasket(t, db, basket.Id)
		defer deleteProduct(t, db, product.Id)
	})
}
func deleteBasket(t *testing.T, db *pgx.Conn, basketID string) {
	// 1. Delete all basket items associated with the basket
	_, err := db.Exec(context.Background(), "DELETE FROM basket_items WHERE basket_id = $1", basketID)
	if err != nil {
		t.Fatalf("Error deleting basket items: %v", err)
	}

	// 2. Delete the basket
	_, err = db.Exec(context.Background(), "DELETE FROM baskets WHERE id = $1", basketID)
	if err != nil {
		t.Fatalf("Error deleting basket: %v", err)
	}
}
func deleteBasketItem(t *testing.T, db *pgx.Conn, basketItemID string) {
	_, err := db.Exec(context.Background(), "DELETE FROM basket_items WHERE id = $1", basketItemID)
	if err != nil {
		t.Fatalf("Error deleting basket item: %v", err)
	}
}

package test

import (
	"context"
	"log"
	"testing"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/food-delivery/food-delivery-order-service/storage/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestProductRepo(t *testing.T) {
	db := createDBConnection(t) // Use the existing createDBConnection function
	defer db.Close(context.Background())

	productRepo := postgres.NewProductRepo(db)

	t.Run("CreateProduct", func(t *testing.T) {
		req := &order.CreateProductRequest{
			Product: &order.Product{
				Name:        "Pizza",
				Description: "Delicious pizza",
				Price:       10.99,
				ImageUrl:    "https://example.com/pizza.jpg",
			},
		}

		response, err := productRepo.CreateProduct(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		// Cleanup
		defer deleteProduct(t, db, response.Id)
	})

	t.Run("GetProductByID", func(t *testing.T) {
		req := &order.CreateProductRequest{
			Product: &order.Product{
				Name:        "Burger",
				Description: "Juicy burger",
				Price:       8.99,
				ImageUrl:    "https://example.com/burger.jpg",
			},
		}

		response, err := productRepo.CreateProduct(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		product, err := productRepo.GetProductByID(context.Background(), response.Id)
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, response.Id, product.Id)

		// Cleanup
		defer deleteProduct(t, db, response.Id)
	})

	t.Run("GetAllProducts", func(t *testing.T) {
		// Create a few test products
		productsToCreate := []*order.Product{
			{
				Name:        "Sushi",
				Description: "Fresh sushi",
				Price:       12.99,
				ImageUrl:    "https://example.com/sushi.jpg",
			},
			{
				Name:        "Tacos",
				Description: "Delicious tacos",
				Price:       7.99,
				ImageUrl:    "https://example.com/tacos.jpg",
			},
		}

		for _, product := range productsToCreate {
			_, err := productRepo.CreateProduct(context.Background(), &order.CreateProductRequest{Product: product})
			assert.NoError(t, err)
		}

		// Test GetAllProducts with no filters
		response, err := productRepo.GetAllProducts(context.Background(), &order.GetProductsRequest{})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(response.Products), 2) // At least 2 products should be returned
		log.Print()
		// Test GetAllProducts with name filter
		response, err = productRepo.GetAllProducts(context.Background(), &order.GetProductsRequest{Name: "Tacos"})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(response.Products), 1) // At least 1 product with "Pizza" in the name

		// Test GetAllProducts with price filter
		response, err = productRepo.GetAllProducts(context.Background(), &order.GetProductsRequest{MinPrice: 6, MaxPrice: 13})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(response.Products), 2) // At least 2 products within the price range

		// Cleanup
		for _, product := range productsToCreate {
			defer deleteProduct(t, db, product.Id)
		}
	})

	t.Run("UpdateProduct", func(t *testing.T) {
		req := &order.CreateProductRequest{
			Product: &order.Product{
				Name:        "Pasta",
				Description: "Italian pasta",
				Price:       9.99,
				ImageUrl:    "https://example.com/pasta.jpg",
			},
		}

		response, err := productRepo.CreateProduct(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		// Update the product
		updateReq := &order.UpdateProductRequest{
			Product: &order.Product{
				Id:          response.Id,
				Name:        "Spaghetti",
				Description: "Classic spaghetti",
				Price:       11.99,
				ImageUrl:    "https://example.com/spaghetti.jpg",
			},
		}
		updatedProduct, err := productRepo.UpdateProduct(context.Background(), updateReq)
		assert.NoError(t, err)
		assert.Equal(t, "Spaghetti", updatedProduct.Name)
		assert.Equal(t, 11.99, updatedProduct.Price)

		// Cleanup
		defer deleteProduct(t, db, response.Id)
	})

	t.Run("PatchProduct", func(t *testing.T) {
		req := &order.CreateProductRequest{
			Product: &order.Product{
				Name:        "Salad",
				Description: "Fresh salad",
				Price:       6.99,
				ImageUrl:    "https://example.com/salad.jpg",
			},
		}

		response, err := productRepo.CreateProduct(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		// Patch the product
		patchReq := &order.PatchProductRequest{
			Id:    response.Id,
			Name:  "Caesar Salad",
			Price: 7.99,
		}
		patchedProduct, err := productRepo.PatchProduct(context.Background(), patchReq)
		assert.NoError(t, err)
		assert.Equal(t, "Caesar Salad", patchedProduct.Name)
		assert.Equal(t, 7.99, patchedProduct.Price)

		// Cleanup
		defer deleteProduct(t, db, response.Id)
	})

	t.Run("DeleteProduct", func(t *testing.T) {
		req := &order.CreateProductRequest{
			Product: &order.Product{
				Name:        "Soup",
				Description: "Warm soup",
				Price:       5.99,
				ImageUrl:    "https://example.com/soup.jpg",
			},
		}

		response, err := productRepo.CreateProduct(context.Background(), req)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Id)

		err = productRepo.DeleteProduct(context.Background(), response.Id)
		assert.NoError(t, err)

		_, err = productRepo.GetProductByID(context.Background(), response.Id)
		assert.ErrorIs(t, err, pgx.ErrNoRows) // Product should not be found
	})
}
func deleteProduct(t *testing.T, db *pgx.Conn, productID string) {
	// 1. Delete all basket items associated with the product
	_, err := db.Exec(context.Background(), "DELETE FROM basket_items WHERE product_id = $1", productID)
	if err != nil {
		t.Fatalf("Error deleting basket items: %v", err)
	}

	// 2. Delete the product
	_, err = db.Exec(context.Background(), "DELETE FROM products WHERE id = $1", productID)
	if err != nil {
		t.Fatalf("Error deleting product: %v", err)
	}
}

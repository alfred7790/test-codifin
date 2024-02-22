package repository

import (
	"codifin-challenge/domain/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// Test_GetList tests the GetList function of the ProductRepository.
func Test_GetList(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		t.Fatalf("Error migrating table: %v", err)
	}

	repo := NewProductRepository(db)

	for i := 0; i < 10; i++ {
		err = db.Create(&model.Product{Name: fmt.Sprintf("Product %d", i)}).Error
		if err != nil {
			t.Fatalf("Error inserting product: %v", err)
		}
	}

	products, total, err := repo.GetList(1, 5, "", "id", true)
	if err != nil {
		t.Fatalf("Error getting product list: %v", err)
	}

	if len(products) != 5 {
		t.Errorf("Expected 5 products, found %d", len(products))
	}

	if total != 10 {
		t.Errorf("Expected 10 total products, found %d", total)
	}
}

// Test_Create tests the Create function of the ProductRepository.
func Test_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		t.Fatalf("Error migrating table: %v", err)
	}

	repo := NewProductRepository(db)

	newProduct := &model.Product{Name: "New Product"}
	err = repo.Create(newProduct)
	if err != nil {
		t.Fatalf("Error creating new product: %v", err)
	}

	product, err := repo.GetByID(newProduct.ID)
	if err != nil {
		t.Fatalf("Error getting created product: %v", err)
	}

	if product == nil {
		t.Fatalf("Expected a product, got nil")
	}

	if product.Name != newProduct.Name {
		t.Errorf("Incorrect product name. Expected '%s', got '%s'", newProduct.Name, product.Name)
	}
}

// Test_GetByID tests the GetByID function of the ProductRepository.
func Test_GetByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		t.Fatalf("Error migrating table: %v", err)
	}

	repo := NewProductRepository(db)

	exampleProduct := &model.Product{Name: "Example Product"}
	err = db.Create(exampleProduct).Error
	if err != nil {
		t.Fatalf("Error inserting example product: %v", err)
	}

	product, err := repo.GetByID(exampleProduct.ID)
	if err != nil {
		t.Fatalf("Error getting product by ID: %v", err)
	}

	if product == nil {
		t.Fatalf("Expected a product, got nil")
	}

	if product.Name != exampleProduct.Name {
		t.Errorf("Incorrect product name. Expected '%s', got '%s'", exampleProduct.Name, product.Name)
	}
}

// Test_Update tests the Update function of the ProductRepository.
func Test_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		t.Fatalf("Error migrating table: %v", err)
	}

	repo := NewProductRepository(db)

	productToUpdate := &model.Product{Name: "Product to Update"}
	err = repo.Create(productToUpdate)
	if err != nil {
		t.Fatalf("Error creating product to update: %v", err)
	}

	productToUpdate.Name = "Updated Product"
	err = repo.Update(productToUpdate)
	if err != nil {
		t.Fatalf("Error updating product: %v", err)
	}

	updatedProduct, err := repo.GetByID(productToUpdate.ID)
	if err != nil {
		t.Fatalf("Error getting updated product: %v", err)
	}

	if updatedProduct == nil {
		t.Fatalf("Expected an updated product, got nil")
	}

	if updatedProduct.Name != productToUpdate.Name {
		t.Errorf("Incorrect product name. Expected '%s', got '%s'", productToUpdate.Name, updatedProduct.Name)
	}
}

// Test_Delete tests the Delete function of the ProductRepository.
func Test_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error trying to open sqlite: %v", err)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		t.Fatalf("error trying to migrate product model: %v", err)
	}

	repo := NewProductRepository(db)

	productToDelete := &model.Product{Name: "Product to Delete"}
	err = repo.Create(productToDelete)
	if err != nil {
		t.Fatalf("error trying to create product: %v", err)
	}

	err = repo.Delete(productToDelete.ID)
	if err != nil {
		t.Fatalf("error trying to remove product: %v", err)
	}

	deletedProduct, err := repo.GetByID(productToDelete.ID)
	if err == nil && deletedProduct != nil {
		t.Fatalf("Expected the product to be deleted, but it still exists")
	}
}

package repository

import (
	"codifin-challenge/domain/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func Test_GetList(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir la base de datos: %v", err)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		t.Fatalf("Error al migrar la tabla: %v", err)
	}

	repo := NewProductRepository(db)

	for i := 0; i < 10; i++ {
		err := db.Create(&model.Product{Name: fmt.Sprintf("Product %d", i)}).Error
		if err != nil {
			t.Fatalf("Error al insertar producto: %v", err)
		}
	}

	products, total, err := repo.GetList(1, 5, "", "id", true)
	if err != nil {
		t.Fatalf("Error al obtener la lista de productos: %v", err)
	}

	if len(products) != 5 {
		t.Errorf("Se esperaban 5 productos, se encontraron %d", len(products))
	}

	if total != 10 {
		t.Errorf("Se esperaban 10 productos en total, se encontraron %d", total)
	}
}

func Test_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir la base de datos: %v", err)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		t.Fatalf("Error al migrar la tabla: %v", err)
	}

	repo := NewProductRepository(db)

	newProduct := &model.Product{Name: "New Product"}
	err = repo.Create(newProduct)
	if err != nil {
		t.Fatalf("Error al crear un nuevo producto: %v", err)
	}

	product, err := repo.GetByID(newProduct.ID)
	if err != nil {
		t.Fatalf("Error al obtener el producto creado: %v", err)
	}

	if product == nil {
		t.Fatalf("Se esperaba un producto, se obtuvo nil")
	}

	if product.Name != newProduct.Name {
		t.Errorf("Nombre del producto incorrecto. Se esperaba '%s', se obtuvo '%s'", newProduct.Name, product.Name)
	}
}

func Test_GetByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir la base de datos: %v", err)
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		t.Fatalf("Error al migrar la tabla: %v", err)
	}

	repo := NewProductRepository(db)

	exampleProduct := &model.Product{Name: "Example Product"}
	err = db.Create(exampleProduct).Error
	if err != nil {
		t.Fatalf("Error al insertar producto de ejemplo: %v", err)
	}

	product, err := repo.GetByID(exampleProduct.ID)
	if err != nil {
		t.Fatalf("Error al obtener el producto por ID: %v", err)
	}

	if product == nil {
		t.Fatalf("Se esperaba un producto, se obtuvo nil")
	}

	if product.Name != exampleProduct.Name {
		t.Errorf("Nombre del producto incorrecto. Se esperaba '%s', se obtuvo '%s'", exampleProduct.Name, product.Name)
	}
}

func Test_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir la base de datos: %v", err)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		t.Fatalf("Error al migrar la tabla: %v", err)
	}

	repo := NewProductRepository(db)

	productToUpdate := &model.Product{Name: "Product to Update"}
	err = repo.Create(productToUpdate)
	if err != nil {
		t.Fatalf("Error al crear el producto a actualizar: %v", err)
	}

	productToUpdate.Name = "Updated Product"
	err = repo.Update(productToUpdate)
	if err != nil {
		t.Fatalf("Error al actualizar el producto: %v", err)
	}

	updatedProduct, err := repo.GetByID(productToUpdate.ID)
	if err != nil {
		t.Fatalf("Error al obtener el producto actualizado: %v", err)
	}

	if updatedProduct == nil {
		t.Fatalf("Se esperaba un producto actualizado, se obtuvo nil")
	}

	if updatedProduct.Name != productToUpdate.Name {
		t.Errorf("Nombre del producto incorrecto. Se esperaba '%s', se obtuvo '%s'", productToUpdate.Name, updatedProduct.Name)
	}
}

func Test_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir la base de datos: %v", err)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		t.Fatalf("Error al migrar la tabla: %v", err)
	}

	repo := NewProductRepository(db)

	productToDelete := &model.Product{Name: "Product to Delete"}
	err = repo.Create(productToDelete)
	if err != nil {
		t.Fatalf("Error al crear el producto a eliminar: %v", err)
	}

	err = repo.Delete(productToDelete.ID)
	if err != nil {
		t.Fatalf("Error al eliminar el producto: %v", err)
	}

	deletedProduct, err := repo.GetByID(productToDelete.ID)
	if err == nil && deletedProduct != nil {
		t.Fatalf("Se esperaba que el producto se eliminara, pero aún existe")
	}
}

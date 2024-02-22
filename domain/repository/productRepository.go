// Package repository provides implementations for interacting with product data in the database.
package repository

import (
	"codifin-challenge/domain/model"
	"codifin-challenge/domain/utils"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

// ProductRepository defines methods for interacting with product data.
type ProductRepository interface {
	GetList(page, pageSize int, searchTerm, orderBy string, ascending bool) ([]*model.Product, uint, error)
	GetByID(productID uint) (*model.Product, error)
	Create(p *model.Product) error
	Update(p *model.Product) error
	Delete(productID uint) error
}

// ProductRepositoryImpl is an implementation of ProductRepository.
type ProductRepositoryImpl struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepositoryImpl.
func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

// GetList retrieves a list of products with pagination support.
func (r *ProductRepositoryImpl) GetList(page, pageSize int, searchTerm, orderBy string, ascending bool) ([]*model.Product, uint, error) {
	var products []*model.Product
	var total int64

	query := r.db.Model(&model.Product{})

	if len(searchTerm) > 0 {
		query = query.Where(r.db.Where("name ILIKE ?", "%"+searchTerm+"%").
			Or("name ILIKE ?", "%"+searchTerm+"%"))
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, utils.ToUserError(http.StatusInternalServerError, "No fue posible obtener el total productos debido a un error interno", err)
	}

	if orderBy != "" {
		orderDirection := "ASC"
		if !ascending {
			orderDirection = "DESC"
		}
		query = query.Order(orderBy + " " + orderDirection)
	}

	err := query.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&products).Error
	if err != nil {
		return nil, 0, utils.ToUserError(http.StatusInternalServerError, "No fue posible obtener la lista de productos debido a un error interno", err)
	}

	return products, uint(total), nil
}

// GetByID retrieves a product by its ID.
func (r *ProductRepositoryImpl) GetByID(productID uint) (*model.Product, error) {
	var product *model.Product

	err := r.db.Model(&model.Product{}).
		Where("id = ?", productID).
		First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ToUserError(http.StatusNotFound, "El producto solicitado no existe", err)
		}
		return nil, utils.ToUserError(http.StatusInternalServerError, "No fue posible obtener detalle de producto debido a un error interno", err)
	}

	return product, nil
}

// Create adds a new product to the database.
func (r *ProductRepositoryImpl) Create(p *model.Product) error {
	err := r.db.Create(p).Error
	if err != nil {
		return utils.ToUserError(http.StatusInternalServerError, "No fue posible registrar el producto debido a un error interno", err)
	}

	return nil
}

// Update updates an existing product in the database.
func (r *ProductRepositoryImpl) Update(p *model.Product) error {
	err := r.db.Save(p).Error
	if err != nil {
		return utils.ToUserError(http.StatusInternalServerError, "No fue posible actualizar el producto debido a un error interno", err)
	}

	return nil
}

// Delete deletes a product from the database by its ID.
func (r *ProductRepositoryImpl) Delete(id uint) error {
	err := r.db.Delete(&model.Product{}, id).Error
	if err != nil {
		return utils.ToUserError(http.StatusInternalServerError, "No fue posible eliminar el producto debido a un error interno", err)
	}

	return nil
}

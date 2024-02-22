// Package service provides implementations for interacting with product data.
package service

import (
	"codifin-challenge/domain/model"
	"codifin-challenge/domain/repository"
	"codifin-challenge/domain/utils"
	"fmt"
	"net/http"
)

// ProductService defines methods for interacting with product data.
type ProductService interface {
	ProductsList(page, pageSize int, searchTerm, orderBy string, ascending bool) ([]*model.Product, uint, error)
	ProductByID(productID uint) (*model.Product, error)
	CreateProduct(p *model.Product) error
	UpdateProduct(productID uint, updates map[string]interface{}) error
	DeleteProduct(productID uint) error
}

// ProductServiceImpl is an implementation of ProductService.
type ProductServiceImpl struct {
	productRepo repository.ProductRepository
}

// NewProductService creates a new instance of ProductServiceImpl.
func NewProductService(repo repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{productRepo: repo}
}

// ProductsList retrieves a list of products with pagination.
func (s *ProductServiceImpl) ProductsList(page, pageSize int, searchTerm, orderBy string, ascending bool) ([]*model.Product, uint, error) {
	return s.productRepo.GetList(page, pageSize, searchTerm, orderBy, ascending)
}

// ProductByID retrieves a product by its ID.
func (s *ProductServiceImpl) ProductByID(productID uint) (*model.Product, error) {
	return s.productRepo.GetByID(productID)
}

// CreateProduct creates a new product.
func (s *ProductServiceImpl) CreateProduct(p *model.Product) error {
	return s.productRepo.Create(p)
}

// DeleteProduct deletes a product by its ID.
func (s *ProductServiceImpl) DeleteProduct(productID uint) error {
	return s.productRepo.Delete(productID)
}

// UpdateProduct updates an existing product.
func (s *ProductServiceImpl) UpdateProduct(productID uint, updates map[string]interface{}) error {
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return err
	}

	for field, value := range updates {
		if err = s.assignUpdates(product, field, value); err != nil {
			return err
		}
	}

	return s.productRepo.Update(product)
}

// assignUpdates assign and validate updates the fields of a product.
func (s *ProductServiceImpl) assignUpdates(product *model.Product, field string, value interface{}) error {
	typeMsg := fmt.Sprintf("Field: %s, Value type: %T, Value: %v\n", field, value, value)
	var err error
	var message string
	switch field {
	case "code":
		if v, ok := value.(string); ok {
			product.Code = v
		} else {
			message = fmt.Sprintf("El valor para el codigo de producto es invalido")
			err = fmt.Errorf("invalid type for 'code' field: %s", typeMsg)
		}
	case "name":
		if v, ok := value.(string); ok {
			product.Name = v
		} else {
			message = fmt.Sprintf("El valor para el nombre de producto es invalido")
			err = fmt.Errorf("invalid type for 'name' field: %s", typeMsg)
		}
	case "price":
		if v, ok := value.(float64); ok {
			product.Price = v
		} else {
			message = fmt.Sprintf("El valor para el precio del producto es invalido")
			err = fmt.Errorf("invalid type for 'price', %s", typeMsg)
		}
	case "imageURL":
		if v, ok := value.(string); ok {
			product.ImageURL = v
		} else {
			message = fmt.Sprintf("El valor para la imagen de producto es invalido")
			err = fmt.Errorf("invalid type for 'imageURL' field: %s", typeMsg)
		}
	default:
		message = fmt.Sprintf("El campo %s no existe en el modelo producto", field)
		err = fmt.Errorf("'%s' field doesn't exist in product model", field)
	}

	if err != nil {
		return utils.ToUserError(http.StatusBadRequest, message, err)
	}

	return nil
}

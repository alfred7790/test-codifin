// Package repository provides implementations for interacting with shopping cart data in the database.
package repository

import (
	"codifin-challenge/domain/model"
	"codifin-challenge/domain/utils"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

// ShoppingCartRepository defines methods for interacting with shopping cart data.
type ShoppingCartRepository interface {
	Create(shoppingCart *model.ShoppingCart) error
	AddItems(items []*model.ItemCart) error
	GetByID(shoppingCartID uint) (*model.ShoppingCart, error)
	DeleteProducts(cartID uint, itemIds []uint) error
}

// ShoppingCartRepositoryImpl is an implementation of ShoppingCartRepository.
type ShoppingCartRepositoryImpl struct {
	db *gorm.DB
}

// NewShoppingCartRepository creates a new instance of ShoppingCartRepositoryImpl.
func NewShoppingCartRepository(db *gorm.DB) *ShoppingCartRepositoryImpl {
	return &ShoppingCartRepositoryImpl{db: db}
}

// GetByID retrieves a shopping cart by its ID.
func (r *ShoppingCartRepositoryImpl) GetByID(shoppingCartID uint) (*model.ShoppingCart, error) {
	var cart model.ShoppingCart

	err := r.db.Model(&model.ShoppingCart{}).
		Preload("Items.Product").
		Where("id = ?", shoppingCartID).
		First(&cart).Error
	if err != nil {
		return nil, utils.ToUserError(http.StatusInternalServerError, "No fue posible crear un nuevo carrito debido a un error interno", err)
	}
	return &cart, nil
}

// Create creates a new shopping cart in the database.
func (r *ShoppingCartRepositoryImpl) Create(shoppingCart *model.ShoppingCart) error {
	err := r.db.Model(&model.ShoppingCart{}).
		Create(&shoppingCart).Error
	if err != nil {
		return utils.ToUserError(http.StatusInternalServerError, "No fue posible crear un nuevo carrito debido a un error interno", err)
	}

	return nil
}

// AddItems adds items to a shopping cart in the database.
func (r *ShoppingCartRepositoryImpl) AddItems(items []*model.ItemCart) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return utils.ToUserError(http.StatusInternalServerError, "No fue posible iniciar transaccion debido a un error interno", tx.Error)
	}

	for _, v := range items {
		var existingItem model.ItemCart
		result := tx.Where("shopping_cart_id = ? AND product_id = ?", v.ShoppingCartID, v.ProductID).First(&existingItem)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return utils.ToUserError(http.StatusInternalServerError, "No fue posible agregar productos al carrito debido a un error interno", result.Error)
		}

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result = tx.Create(&v)
			if result.Error != nil {
				tx.Rollback()
				return utils.ToUserError(http.StatusInternalServerError, "No fue posible agregar productos al carrito debido a un error interno", result.Error)
			}
		} else {
			result = tx.Model(&existingItem).Update("count", gorm.Expr("count + ?", v.Count))
			if result.Error != nil {
				tx.Rollback()
				return utils.ToUserError(http.StatusInternalServerError, "No fue posible agregar productos al carrito debido a un error interno", result.Error)
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return utils.ToUserError(http.StatusInternalServerError, "No fue posible agregar productos al carrito debido a un error interno", err)
	}

	return nil
}

// DeleteProducts deletes products from a shopping cart.
func (r *ShoppingCartRepositoryImpl) DeleteProducts(cartID uint, itemIds []uint) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return utils.ToUserError(http.StatusInternalServerError, "No fue posible iniciar transaccion debido a un error interno", tx.Error)
	}

	for _, itemId := range itemIds {
		err := tx.Where("shopping_cart_id = ? AND product_id = ?", cartID, itemId).
			Delete(&model.ItemCart{}).Error
		if err != nil {
			tx.Rollback()
			return utils.ToUserError(http.StatusInternalServerError, "No fue posible eliminar productos del carrito debido a un error interno", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return utils.ToUserError(http.StatusInternalServerError, "No fue posible eliminar el carrito debido a un error interno", err)
	}

	return nil
}

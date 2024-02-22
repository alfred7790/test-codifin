// Package service provides implementations for interacting with shopping cart data.
package service

import (
	"codifin-challenge/domain/model"
	"codifin-challenge/domain/repository"
)

// ShoppingCartService defines methods for interacting with shopping cart data.
type ShoppingCartService interface {
	CreateShoppingCart(shoppingCart *model.ShoppingCart) error
	AddItemsToShoppingCart(items []*model.ItemCart) error
	FindCart(cartID uint) (*model.ShoppingCart, error)
	RemoveItemsFromShoppingCart(cartId uint, items []uint) error
}

// ShoppingCartServiceImpl is an implementation of ShoppingCartService.
type ShoppingCartServiceImpl struct {
	shoppingCartRepo repository.ShoppingCartRepository
}

// NewShoppingCartService creates a new instance of ShoppingCartServiceImpl.
func NewShoppingCartService(repo repository.ShoppingCartRepository) *ShoppingCartServiceImpl {
	return &ShoppingCartServiceImpl{shoppingCartRepo: repo}
}

// CreateShoppingCart creates a new shopping cart.
func (s *ShoppingCartServiceImpl) CreateShoppingCart(shoppingCart *model.ShoppingCart) error {
	return s.shoppingCartRepo.Create(shoppingCart)
}

// AddItemsToShoppingCart adds items to a shopping cart.
func (s *ShoppingCartServiceImpl) AddItemsToShoppingCart(items []*model.ItemCart) error {
	return s.shoppingCartRepo.AddItems(items)
}

// RemoveItemsFromShoppingCart removes items from a shopping cart.
func (s *ShoppingCartServiceImpl) RemoveItemsFromShoppingCart(cartId uint, items []uint) error {
	return s.shoppingCartRepo.DeleteProducts(cartId, items)
}

// FindCart finds a shopping cart by its ID.
func (s *ShoppingCartServiceImpl) FindCart(cartID uint) (*model.ShoppingCart, error) {
	return s.shoppingCartRepo.GetByID(cartID)
}

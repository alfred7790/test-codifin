package dto

import (
	"codifin-challenge/domain/model"
)

type ShoppingCartDTO struct {
	ID    uint           `json:"id"`
	Items []*ItemCartDTO `json:"items"`
}

type ItemData struct {
	ProductID uint `json:"productID"`
	Count     uint `json:"count"`
}

type ItemCartDTO struct {
	Product   *ProductDTO `json:"product"`
	ProductID uint        `json:"productID"`
	Count     uint        `json:"count"`
}

func ToItemsCart(shoppingCartID uint, items []*ItemData) []*model.ItemCart {
	itemsCart := make([]*model.ItemCart, 0)

	for _, v := range items {
		itemsCart = append(itemsCart, v.ToItemCart(shoppingCartID))
	}

	return itemsCart
}

func (i *ItemData) ToItemCart(shoppingCartID uint) *model.ItemCart {
	return &model.ItemCart{
		ShoppingCartID: shoppingCartID,
		ProductID:      i.ProductID,
		Count:          i.Count,
	}
}

func ToShoppingCartDTO(cart *model.ShoppingCart) *ShoppingCartDTO {
	if cart != nil {
		return &ShoppingCartDTO{
			ID:    cart.ID,
			Items: ToItemsCartDTO(cart.Items),
		}
	}

	return nil
}

func ToItemCartDTO(item *model.ItemCart) *ItemCartDTO {
	return &ItemCartDTO{
		ProductID: item.ProductID,
		Product:   ToProductDTO(item.Product),
		Count:     item.Count,
	}
}

func ToItemsCartDTO(items []*model.ItemCart) []*ItemCartDTO {
	itemsDTO := make([]*ItemCartDTO, 0)
	for _, v := range items {
		itemsDTO = append(itemsDTO, ToItemCartDTO(v))
	}
	return itemsDTO
}

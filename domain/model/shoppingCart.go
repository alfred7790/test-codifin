package model

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	Items []*ItemCart
}

type ItemCart struct {
	gorm.Model
	ShoppingCartID uint          `gorm:"not null"`
	ShoppingCart   *ShoppingCart `gorm:"foreignKey:ShoppingCartID"`
	ProductID      uint          `gorm:"not null"`
	Product        *Product      `gorm:"foreignKey:ProductID"`
	Count          uint          `gorm:"default:0"`
}

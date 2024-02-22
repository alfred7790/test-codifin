package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code     string
	Name     string
	Price    float64
	ImageURL string
}

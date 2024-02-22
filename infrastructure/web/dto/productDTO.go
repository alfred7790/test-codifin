package dto

import (
	"codifin-challenge/domain/model"
	"strings"
)

type ProductsListResp struct {
	Total    uint          `json:"total"`
	Products []*ProductDTO `json:"products"`
}

type ProductDTO struct {
	ID uint `json:"id"`
	ProductData
}

type ProductData struct {
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"imageURL"`
}

func (p *ProductData) ToProduct() *model.Product {
	return &model.Product{
		Code:     strings.TrimSpace(p.Code),
		Name:     strings.TrimSpace(strings.ToUpper(p.Name)),
		Price:    p.Price,
		ImageURL: strings.TrimSpace(p.ImageURL),
	}
}

func ToProductDTO(product *model.Product) *ProductDTO {
	if product != nil {
		return &ProductDTO{
			ID: product.ID,
			ProductData: ProductData{
				Code:     product.Code,
				Name:     product.Name,
				Price:    product.Price,
				ImageURL: product.ImageURL,
			},
		}
	}
	return nil
}

func ToProductsDTO(products []*model.Product) []*ProductDTO {
	productsDTO := make([]*ProductDTO, 0)
	for _, v := range products {
		productsDTO = append(productsDTO, ToProductDTO(v))
	}

	return productsDTO
}

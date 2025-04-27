package entity

import "github.com/google/uuid"

type Food struct {
	Id        string `json:"id"`
	VariantId string `json:"variant_id"`
	Name      string `json:"name"`
	ImageUrl  string `json:"image_url"`
	Price     string `json:"price"`
	Qty       string `json:"qty"`
	Ratings   string `json:"ratings"`
}

func NewFood(variantId, name, imageUrl, price, qty string) *Food {
	return &Food{
		Id:        uuid.New().String(),
		VariantId: variantId,
		Name:      name,
		ImageUrl:  imageUrl,
		Price:     price,
		Qty:       qty,
		Ratings:   "3",
	}
}

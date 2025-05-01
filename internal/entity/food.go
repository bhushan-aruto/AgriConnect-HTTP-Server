package entity

import "github.com/google/uuid"

type Food struct {
	Id        string `json:"id"`
	VariantId string `json:"variant_id"`
	Name      string `json:"name"`
	ImageUrl  string `json:"image_url"`
	Unit      string `json:"unit"`
	Price     string `json:"price"`
	Qty       string `json:"qty"`
	Ratings   string `json:"ratings"`
}

func NewFood(variantId, name, unit, imageUrl, price, qty string) *Food {
	return &Food{
		Id:        uuid.New().String(),
		VariantId: variantId,
		Name:      name,
		Unit:      unit,
		ImageUrl:  imageUrl,
		Price:     price,
		Qty:       qty,
		Ratings:   "3",
	}
}

package models

type CreateFoodRequest struct {
	VariantId string `form:"variant_id"`
	Name      string `form:"name"`
	Price     string `form:"price"`
	Qty       string `form:"qty"`
}

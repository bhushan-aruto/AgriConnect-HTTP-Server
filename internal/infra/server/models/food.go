package models

type CreateFoodRequest struct {
	VariantId string `form:"variant_id"`
	Name      string `form:"name"`
	Unit      string `form:"unit"`
	Price     string `form:"price"`
	Qty       string `form:"qty"`
}

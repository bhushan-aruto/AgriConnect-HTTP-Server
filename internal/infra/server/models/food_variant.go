package models

type CreateFoodVariantRequest struct {
	FarmerId string `form:"farmer_id"`
	Name     string `form:"name"`
}

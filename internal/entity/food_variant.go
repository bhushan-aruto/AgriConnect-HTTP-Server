package entity

import "github.com/google/uuid"

type FoodVariant struct {
	Id             string `json:"id"`
	FarmerId       string `json:"farmer_id"`
	Name           string `json:"name"`
	BannerImageUrl string `json:"banner_image_url"`
}

func NewFoodVariant(name, farmerId, bannerImageUrl string) *FoodVariant {
	return &FoodVariant{
		Id:             uuid.New().String(),
		FarmerId:       farmerId,
		Name:           name,
		BannerImageUrl: bannerImageUrl,
	}
}

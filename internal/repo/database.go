package repo

import "github.com/bhushn-aruto/krushi-sayak-http-server/internal/entity"

type DatabaseRepo interface {
	CheckFamerEmailExists(email string) (bool, error)
	CheckFarmerPhoneNumberExists(phNum string) (bool, error)
	CreateFarmer(farmer *entity.Farmer) error
	GetFarmerForLogin(email string) (*entity.Farmer, error)
	GetFarmerPhoneNumberByFoodId(foodId string) (string, error)

	CheckBuyerEmailExists(email string) (bool, error)
	CheckBuyerPhoneNumberExists(phNum string) (bool, error)
	CreateBuyer(buyer *entity.Buyer) error
	GetBuyerForLogin(email string) (*entity.Buyer, error)

	CheckFoodVariantExists(farmerId string, name string) (bool, error)
	CreateFoodVariant(variant *entity.FoodVariant) error
	GetFoodVariantsByFormerId(farmerId string) ([]*entity.FoodVariant, error)
	GetFoodVariantImageUrl(id string) (string, error)
	DeleteFoodVariant(id string) error

	CheckFoodExists(variantId string, foodName string) (bool, error)
	CreateFood(food *entity.Food) error
	GetFoodsByVariantId(variantId string) ([]*entity.Food, error)
	GetFoodImageUrl(id string) (string, error)
	DeleteFood(foodId string) error

	GetAllFoods() ([]*entity.Food, error)

	GetBuyerDetails(buyerId string) (*entity.Buyer, error)
	GetFoodQty(itemId string) (string, error)
	CreateBuyerOrder(order *entity.Order, qty string) error
	GetOrdersByFarmerId(farmerId string) ([]*entity.OrderResponse, error)
	DeleteOrder(orderId string) error
}

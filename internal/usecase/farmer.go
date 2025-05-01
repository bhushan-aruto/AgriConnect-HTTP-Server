package usecase

import (
	"errors"
	"log"
	"sync"

	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/entity"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/repo"
	"github.com/bhushn-aruto/krushi-sayak-http-server/utils"
)

type item struct {
	name           string
	bannerImageUrl string
}

var itemsSections = []item{
	item{
		name:           "Fruits",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/fruits.jpeg",
	},
	item{
		name:           "Grains and Cereals",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/grains_and_cerals.jpeg",
	},
	item{
		name:           "Pulses and Legumes",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/pulses_and_legumes.jpeg",
	},
	item{
		name:           "Spices and Herbs",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/spices_and_herbs.jpeg",
	},
	item{
		name:           "Flowers",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/flowers.jpeg",
	},
	item{
		name:           "Dairy Products",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/dairy_products.jpeg",
	},
	item{
		name:           "Poultry and Eggs",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/poultry_and_eggs.jpeg",
	},
	item{
		name:           "Meat and Fish",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/meat_and_fish.jpeg",
	},
	item{
		name:           "Organic Produce",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/organic_produce.jpeg",
	},
	item{
		name:           "Nuts and Dry Fruits",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/nuts_and_dry_fruits.jpeg",
	},
	item{
		name:           "Medicinal Plants",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/medicinal_plants.jpeg",
	},
	item{
		name:           "Honey and Beekeeping Products",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/honey_and_beekeeping_products.jpeg",
	},
	item{
		name:           "Handmade / Farm Processed Goods",
		bannerImageUrl: "https://agriconnect.vsensetech.in/public/item_category_sec/handmade.jpeg",
	},
}

type FarmerUseCase struct {
	dbRepo      repo.DatabaseRepo
	storageRepo repo.StorageRepo
}

func NewFormerUseCase(dbRepo repo.DatabaseRepo, storageRepo repo.StorageRepo) *FarmerUseCase {
	return &FarmerUseCase{
		dbRepo:      dbRepo,
		storageRepo: storageRepo,
	}
}

func (u *FarmerUseCase) SignUp(email, fullName, phoneNumber, password string) (int32, error) {

	exists, err := u.dbRepo.CheckFamerEmailExists(email)
	if err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	if exists {
		return 400, errors.New("farmer email already exists")
	}

	exists, err = u.dbRepo.CheckFarmerPhoneNumberExists(phoneNumber)

	if err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	if exists {
		return 400, errors.New("farmer phone number already exists")
	}

	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		log.Println(err)
		return 500, errors.New("error occurred while hashing password")
	}

	farmer := entity.NewFarmer(
		email,
		fullName,
		phoneNumber,
		hashedPassword,
	)

	if err := u.dbRepo.CreateFarmer(farmer); err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	var isErrorOccurred bool

	var wg sync.WaitGroup
	for _, itemSec := range itemsSections {
		wg.Add(1)
		go func() {

			defer wg.Done()

			fvu := NewFoodVariantUseCase(
				u.dbRepo,
				u.storageRepo,
			)

			_, err := fvu.CreateFoodVariant(
				farmer.FarmerId,
				itemSec.name,
				itemSec.bannerImageUrl,
			)

			if err != nil {
				isErrorOccurred = true
				log.Println("error occurred while creating the food variant section, Err: ", err.Error())
				return
			}

		}()
	}

	wg.Wait()

	if isErrorOccurred {
		return 500, errors.New("error occurred while creating the food variant section")
	}

	return 201, nil
}

func (u *FarmerUseCase) Login(email, password string) (string, int32, error) {
	exists, err := u.dbRepo.CheckFamerEmailExists(email)
	if err != nil {
		log.Println(err)
		return "", 500, errors.New("error occurred with database")
	}

	if !exists {
		return "", 400, errors.New("farmer email not exists")
	}

	farmer, err := u.dbRepo.GetFarmerForLogin(email)

	if err != nil {
		log.Println(err)
		return "", 500, errors.New("error occurred with database")
	}

	if err := utils.CheckPassword(farmer.Password, password); err != nil {
		return "", 401, errors.New("incorrect password")
	}

	token, err := utils.GenerateToken(
		farmer.FarmerId,
		farmer.FullName,
		farmer.Email,
	)

	if err != nil {
		log.Println(err)
		return "", 500, errors.New("error occurred while generating token")
	}

	return token, 200, nil
}

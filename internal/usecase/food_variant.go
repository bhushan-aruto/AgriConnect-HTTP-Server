package usecase

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/entity"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/repo"
)

type FoodVariant struct {
	dbRepo      repo.DatabaseRepo
	storageRepo repo.StorageRepo
}

func NewFoodVariantUseCase(dbRepo repo.DatabaseRepo, stRepo repo.StorageRepo) *FoodVariant {
	return &FoodVariant{
		dbRepo:      dbRepo,
		storageRepo: stRepo,
	}
}

func (u *FoodVariant) CreateFoodVariant(farmerId, name, bannerImageUrl string) (int32, error) {

	fv := entity.NewFoodVariant(
		name,
		farmerId,
		bannerImageUrl,
	)

	if err := u.dbRepo.CreateFoodVariant(fv); err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	return 201, nil
}

func (u *FoodVariant) GetFoodVariants(farmerId string) ([]*entity.FoodVariant, int32, error) {
	fvs, err := u.dbRepo.GetFoodVariantsByFormerId(farmerId)

	if err != nil {
		log.Println(err)
		return nil, 500, errors.New("error occurred with database")
	}

	return fvs, 200, nil
}

func (u *FoodVariant) DeleteFoodVariant(id string) (int32, error) {

	imageUrl, err := u.dbRepo.GetFoodVariantImageUrl(id)

	if err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	imageUrlArr := strings.Split(imageUrl, ".")

	imageType := imageUrlArr[len(imageUrlArr)-1]

	fileName := fmt.Sprintf("%s.%s", id, imageType)

	if err := u.storageRepo.DeleteFoodVariantImage(fileName); err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with local storage")
	}

	if err := u.dbRepo.DeleteFoodVariant(id); err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	return 200, nil
}

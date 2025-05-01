package usecase

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/entity"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/repo"
)

type FoodUseCase struct {
	dbRepo      repo.DatabaseRepo
	storageRepo repo.StorageRepo
}

func NewFoodUseCase(dbRepo repo.DatabaseRepo, stRepo repo.StorageRepo) *FoodUseCase {
	return &FoodUseCase{
		dbRepo:      dbRepo,
		storageRepo: stRepo,
	}
}

func (u *FoodUseCase) CreateFood(variantId, name, unit, qty, price, fileType string, fileSrc io.Reader) (int32, error) {
	exists, err := u.dbRepo.CheckFoodExists(variantId, name)

	if err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	if exists {
		return 400, errors.New("food already exists")
	}

	f := entity.NewFood(
		variantId,
		name,
		unit,
		"",
		price,
		qty,
	)

	fileName := fmt.Sprintf("%s.%s", f.Id, fileType)

	filePath, err := u.storageRepo.SaveFoodImage(
		fileName,
		fileSrc,
	)

	if err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with file storage")
	}

	f.ImageUrl = fmt.Sprintf("https://agriconnect.vsensetech.in/%s", filePath)

	if err := u.dbRepo.CreateFood(f); err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	return 201, nil
}

func (u *FoodUseCase) GetFoods(variantId string) ([]*entity.Food, int32, error) {
	fs, err := u.dbRepo.GetFoodsByVariantId(variantId)

	if err != nil {
		log.Println(err)
		return nil, 500, errors.New("error occurred with database")
	}

	return fs, 200, nil
}

func (u *FoodUseCase) DeleteFood(id string) (int32, error) {

	imageUrl, err := u.dbRepo.GetFoodImageUrl(id)

	if err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	imageUrlArr := strings.Split(imageUrl, ".")

	imageType := imageUrlArr[len(imageUrlArr)-1]

	fileName := fmt.Sprintf("%s.%s", id, imageType)

	if err := u.storageRepo.DeleteFoodImage(fileName); err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with local storage")
	}

	if err := u.dbRepo.DeleteFood(id); err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}
	return 200, nil
}

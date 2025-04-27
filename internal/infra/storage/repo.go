package storage

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type storageRepo struct {
	foodVariantsDirPath string
	foodsDirPath        string
}

func NewStorageRepo(foodVariantsDirPath, foodsDirPath string) *storageRepo {
	return &storageRepo{
		foodVariantsDirPath: foodVariantsDirPath,
		foodsDirPath:        foodsDirPath,
	}
}

func (repo *storageRepo) Init() {
	if err := os.MkdirAll(repo.foodVariantsDirPath, os.ModePerm); err != nil {
		log.Fatalln("error occurred while creating the dir for food variants, Err: ", err.Error())
	}
	if err := os.MkdirAll(repo.foodsDirPath, os.ModePerm); err != nil {
		log.Fatalln("error occurred while creating the dir for foods, Err: ", err.Error())
	}
}

func (repo *storageRepo) SaveFoodVariantImage(imageName string, file io.Reader) (string, error) {
	imagePath := filepath.Join(repo.foodVariantsDirPath, imageName)
	dst, err := os.Create(imagePath)

	if err != nil {
		return "", err
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)

	return imagePath, err

}

func (repo *storageRepo) DeleteFoodVariantImage(imageName string) error {
	pth := filepath.Join(repo.foodVariantsDirPath, imageName)
	err := os.Remove(pth)
	return err
}

func (repo *storageRepo) SaveFoodImage(imageName string, file io.Reader) (string, error) {
	imagePath := filepath.Join(repo.foodsDirPath, imageName)
	dst, err := os.Create(imagePath)

	if err != nil {
		return "", err
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)

	return imagePath, err
}

func (repo *storageRepo) DeleteFoodImage(imageName string) error {
	pth := filepath.Join(repo.foodsDirPath, imageName)
	err := os.Remove(pth)
	return err
}

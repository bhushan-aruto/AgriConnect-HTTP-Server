package repo

import "io"

type StorageRepo interface {
	SaveFoodVariantImage(imageName string, file io.Reader) (string, error)
	DeleteFoodVariantImage(imageName string) error

	SaveFoodImage(imageName string, file io.Reader) (string, error)
	DeleteFoodImage(imageName string) error
}

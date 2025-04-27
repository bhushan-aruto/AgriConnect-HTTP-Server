package usecase

import (
	"errors"
	"log"

	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/entity"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/repo"
	"github.com/bhushn-aruto/krushi-sayak-http-server/utils"
)

type FarmerUseCase struct {
	dbRepo repo.DatabaseRepo
}

func NewFormerUseCase(dbRepo repo.DatabaseRepo) *FarmerUseCase {
	return &FarmerUseCase{
		dbRepo: dbRepo,
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

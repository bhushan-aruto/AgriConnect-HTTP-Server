package usecase

import (
	"errors"
	"log"

	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/entity"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/repo"
	"github.com/bhushn-aruto/krushi-sayak-http-server/utils"
)

type BuyerUseCase struct {
	dbRepo repo.DatabaseRepo
}

func NewBuyerUseCase(dbRepo repo.DatabaseRepo) *BuyerUseCase {
	return &BuyerUseCase{
		dbRepo: dbRepo,
	}
}

func (u *BuyerUseCase) SignUp(email, fullName, phoneNumber, password string) (int32, error) {
	exists, err := u.dbRepo.CheckBuyerEmailExists(email)
	if err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	if exists {
		return 400, errors.New("buyer email already exists")
	}

	exists, err = u.dbRepo.CheckBuyerPhoneNumberExists(phoneNumber)

	if err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	if exists {
		return 400, errors.New("buyer phone number already exists")
	}

	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		log.Println(err)
		return 500, errors.New("error occurred while hashing password")
	}

	buyer := entity.NewBuyer(
		email,
		fullName,
		phoneNumber,
		hashedPassword,
	)

	if err := u.dbRepo.CreateBuyer(buyer); err != nil {
		log.Println(err)
		return 500, errors.New("error occurred with database")
	}

	return 201, nil
}

func (u *BuyerUseCase) Login(email, password string) (string, int32, error) {
	exists, err := u.dbRepo.CheckBuyerEmailExists(email)
	if err != nil {
		log.Println(err)
		return "", 500, errors.New("error occurred with database")
	}

	if !exists {
		return "", 400, errors.New("buyer email not exists")
	}

	buyer, err := u.dbRepo.GetBuyerForLogin(email)

	if err != nil {
		log.Println(err)
		return "", 500, errors.New("error occurred with database")
	}

	if err := utils.CheckPassword(buyer.Password, password); err != nil {
		return "", 401, errors.New("incorrect password")
	}

	token, err := utils.GenerateToken(
		buyer.BuyerId,
		buyer.FullName,
		buyer.Email,
	)

	if err != nil {
		log.Println(err)
		return "", 500, errors.New("error occurred while generating token")
	}

	return token, 200, nil
}

func (u *BuyerUseCase) GetAllFoods() ([]*entity.Food, int32, error) {
	fs, err := u.dbRepo.GetAllFoods()

	if err != nil {
		log.Println(err)
		return nil, 500, errors.New("error occurred with database")
	}

	return fs, 200, nil
}

package usecase

import (
	"errors"
	"log"
	"strconv"

	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/entity"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/repo"
)

type OrderUseCase struct {
	dbRepo        repo.DatabaseRepo
	CallAnswerApi string
	CallFrom      string
	twilioRepo    repo.TwilioRepo
}

func NewOrderUseCase(dbRepo repo.DatabaseRepo, CallAnswerApi, CallFrom string, twilioRepo repo.TwilioRepo) *OrderUseCase {
	return &OrderUseCase{
		dbRepo:        dbRepo,
		twilioRepo:    twilioRepo,
		CallAnswerApi: CallAnswerApi,
		CallFrom:      CallFrom,
	}
}

func (u *OrderUseCase) CreateOrder(
	buyerId,
	itemId,
	qty,
	address string,
) (int32, error) {

	details, err := u.dbRepo.GetBuyerDetails(
		buyerId,
	)

	if err != nil {
		log.Println("error occurred with database, Err: ", err.Error())
		return 500, errors.New("error occurred with database")
	}

	itemQty, err := u.dbRepo.GetFoodQty(
		itemId,
	)

	if err != nil {
		log.Println("error occurred with database, Err: ", err.Error())
		return 500, errors.New("error occurred with database")
	}

	itemQtyNum, err := strconv.Atoi(itemQty)

	if err != nil {
		log.Println("error occurred while converting string to int, Err: ", err.Error())
		return 500, errors.New("conversion error occurred in server")

	}

	qtyNum, err := strconv.Atoi(qty)

	if err != nil {
		log.Println("error occurred while converting string to int, Err: ", err.Error())
		return 500, errors.New("conversion error occurred in server")
	}

	if qtyNum > itemQtyNum {
		return 400, errors.New("quantity limit exceeded")
	}

	order := entity.NewOrder(
		itemId,
		details.FullName,
		details.PhoneNumber,
		details.Email,
		address,
		qty,
	)

	itemQtyNum -= qtyNum

	itemQty = strconv.Itoa(itemQtyNum)

	if err := u.dbRepo.CreateBuyerOrder(
		order,
		itemQty,
	); err != nil {
		log.Println("error occurred with database, Err: ", err.Error())
		return 500, errors.New("error occurred with database")
	}

	phoneNumber, err := u.dbRepo.GetFarmerPhoneNumberByFoodId(itemId)

	if err != nil {
		log.Println("error occurred with database, Err: ", err.Error())
		return 500, errors.New("error occurred with database")
	}

	u.twilioRepo.MakeOrderCall(
		u.CallAnswerApi,
		u.CallFrom,
		phoneNumber,
	)

	return 201, nil
}

func (u *OrderUseCase) GetOrdersByFarmerId(farmerId string) ([]*entity.OrderResponse, int32, error) {
	orders, err := u.dbRepo.GetOrdersByFarmerId(farmerId)
	if err != nil {
		log.Println("error occurred with database, Err: ", err.Error())
		return nil, 500, errors.New("error occurred with database")
	}

	for _, order := range orders {
		totalQtyNum, err := strconv.Atoi(order.TotalQty)

		if err != nil {
			log.Println("error occurred while converting string to integer, Error: ", err.Error())
			return nil, 500, errors.New("conversion error occurred in server")
		}

		itemPriceNum, err := strconv.Atoi(order.ItemPrice)

		if err != nil {
			log.Println("error occurred while converting string to integer, Error: ", err.Error())
			return nil, 500, errors.New("conversion error occurred in server")
		}

		totalPrice := totalQtyNum * itemPriceNum

		order.TotalPrice = strconv.Itoa(totalPrice)
	}

	if orders == nil {
		return []*entity.OrderResponse{}, 200, nil
	}

	return orders, 200, nil
}

func (u *OrderUseCase) DeleteOrder(orderId string) (int32, error) {
	if err := u.dbRepo.DeleteOrder(orderId); err != nil {
		log.Println("error occurred with database, Err: ", err.Error())
		return 500, errors.New("error occurred with database")
	}
	return 200, nil
}

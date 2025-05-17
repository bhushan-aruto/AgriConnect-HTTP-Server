package handlers

import (
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/infra/server/models"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/repo"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/usecase"
	"github.com/labstack/echo"
)

type BuyerHandler struct {
	dbRepo        repo.DatabaseRepo
	callAnswerApi string
	callFrom      string
	twilioRepo    repo.TwilioRepo
}

func NewBuyerHandler(dbRepo repo.DatabaseRepo, callAnswerApi, callFrom string, twilioRepo repo.TwilioRepo) *BuyerHandler {
	return &BuyerHandler{
		dbRepo:        dbRepo,
		callAnswerApi: callAnswerApi,
		callFrom:      callFrom,
		twilioRepo:    twilioRepo,
	}
}

func (h *BuyerHandler) SignUpHandler(ctx echo.Context) error {
	req := new(models.FarmerSignUpRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(400, echo.Map{
			"message": "invalid json request body",
		})
	}

	u := usecase.NewBuyerUseCase(h.dbRepo)

	statusCode, err := u.SignUp(
		req.Email,
		req.FullName,
		req.PhoneNumber,
		req.Password,
	)

	if err != nil {
		return ctx.JSON(int(statusCode), echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(
		int(statusCode),
		echo.Map{
			"message": "buyer signup successfull",
		},
	)

}

func (h *BuyerHandler) LoginHandler(ctx echo.Context) error {
	req := new(models.BuyerLoginRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(400, echo.Map{
			"message": "invalid json request body",
		})
	}

	u := usecase.NewBuyerUseCase(h.dbRepo)

	token, statusCode, err := u.Login(req.Email, req.Password)

	if err != nil {
		return ctx.JSON(int(statusCode), echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(
		int(statusCode),
		echo.Map{
			"token": token,
		},
	)
}

func (h *BuyerHandler) GetAllFoodsHandler(ctx echo.Context) error {
	u := usecase.NewBuyerUseCase(h.dbRepo)

	fs, statusCode, err := u.GetAllFoods()

	if err != nil {
		return ctx.JSON(int(statusCode), echo.Map{
			"message": err.Error(),
		})
	}

	if fs == nil {
		return ctx.JSON(
			int(statusCode),
			[]interface{}{},
		)
	}

	return ctx.JSON(
		int(statusCode),
		fs,
	)
}

func (h *BuyerHandler) CreateOrderHandler(ctx echo.Context) error {

	req := new(models.OrderRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(400, echo.Map{
			"message": "invalid json request body",
		})
	}

	u := usecase.NewOrderUseCase(
		h.dbRepo,
		h.callAnswerApi,
		h.callFrom,
		h.twilioRepo,
	)

	statusCode, err := u.CreateOrder(
		req.BuyerId,
		req.ItemId,
		req.Qty,
		req.Address,
	)

	if err != nil {
		return ctx.JSON(int(statusCode), echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(int(statusCode), echo.Map{
		"message": "order placed successfully",
	})
}

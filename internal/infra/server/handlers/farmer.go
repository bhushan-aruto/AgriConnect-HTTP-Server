package handlers

import (
	"strings"

	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/infra/server/models"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/repo"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/usecase"
	"github.com/labstack/echo"
)

type FarmerHandler struct {
	dbRepo        repo.DatabaseRepo
	storageRepo   repo.StorageRepo
	callAnswerApi string
	callFrom      string
	twilioRepo    repo.TwilioRepo
}

func NewFormerHandler(dbRepo repo.DatabaseRepo, stRepo repo.StorageRepo, callAnswerApi, callFrom string, twilioRepo repo.TwilioRepo) *FarmerHandler {
	return &FarmerHandler{
		dbRepo:        dbRepo,
		storageRepo:   stRepo,
		callAnswerApi: callAnswerApi,
		callFrom:      callFrom,
		twilioRepo:    twilioRepo,
	}
}

func (h *FarmerHandler) SignUpHandler(ctx echo.Context) error {
	req := new(models.FarmerSignUpRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(400, echo.Map{
			"message": "invalid json request body",
		})
	}

	u := usecase.NewFormerUseCase(h.dbRepo, h.storageRepo)

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
			"message": "farmer signup successfull",
		},
	)

}

func (h *FarmerHandler) LoginHandler(ctx echo.Context) error {
	req := new(models.FarmerLoginRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(400, echo.Map{
			"message": "invalid json request body",
		})
	}

	u := usecase.NewFormerUseCase(h.dbRepo, h.storageRepo)

	token, statusCode, err := u.Login(
		req.Email,
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
			"token": token,
		},
	)

}

func (h *FarmerHandler) GetFoodVariantsHandler(ctx echo.Context) error {
	farmerId := ctx.Param("id")

	u := usecase.NewFoodVariantUseCase(
		h.dbRepo,
		h.storageRepo,
	)

	fvs, statusCode, err := u.GetFoodVariants(farmerId)

	if err != nil {
		return ctx.JSON(
			int(statusCode),
			echo.Map{
				"message": err.Error(),
			},
		)
	}

	if fvs == nil {
		return ctx.JSON(
			int(statusCode),
			[]interface{}{},
		)
	}

	return ctx.JSON(
		int(statusCode),
		fvs,
	)
}

func (h *FarmerHandler) DeleteFoodVariantHandler(ctx echo.Context) error {
	id := ctx.Param("id")

	u := usecase.NewFoodVariantUseCase(
		h.dbRepo,
		h.storageRepo,
	)

	statusCode, err := u.DeleteFoodVariant(id)

	if err != nil {
		return ctx.JSON(
			int(statusCode),
			echo.Map{
				"message": err.Error(),
			},
		)
	}

	return ctx.JSON(
		int(statusCode),
		echo.Map{
			"message": "food variant deleted successfully",
		},
	)
}

func (h *FarmerHandler) CreateFoodHandler(ctx echo.Context) error {
	req := new(models.CreateFoodRequest)

	req.Name = ctx.FormValue("name")
	req.Unit = ctx.FormValue("unit")
	req.Price = ctx.FormValue("price")
	req.Qty = ctx.FormValue("qty")
	req.VariantId = ctx.FormValue("variant_id")

	file, err := ctx.FormFile("file")

	if err != nil {
		return ctx.JSON(
			400,
			echo.Map{
				"message": "file upload failed",
			},
		)
	}

	src, err := file.Open()

	if err != nil {
		return ctx.JSON(
			500,
			echo.Map{
				"message": "filed to open the file",
			},
		)
	}

	defer src.Close()

	fileNameArr := strings.Split(file.Filename, ".")

	inputFileType := fileNameArr[len(fileNameArr)-1]

	u := usecase.NewFoodUseCase(h.dbRepo, h.storageRepo)

	statusCode, err := u.CreateFood(
		req.VariantId,
		req.Name,
		req.Unit,
		req.Qty,
		req.Price,
		inputFileType,
		src,
	)

	if err != nil {
		return ctx.JSON(int(statusCode), echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(
		int(statusCode),
		echo.Map{
			"message": "item added successfully",
		},
	)

}

func (h *FarmerHandler) GetFoodsHandler(ctx echo.Context) error {
	variantId := ctx.Param("id")

	u := usecase.NewFoodUseCase(h.dbRepo, h.storageRepo)

	fs, statusCode, err := u.GetFoods(variantId)

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

	return ctx.JSON(int(statusCode), fs)
}

func (h *FarmerHandler) DeleteFoodHandler(ctx echo.Context) error {
	foodId := ctx.Param("id")

	u := usecase.NewFoodUseCase(
		h.dbRepo,
		h.storageRepo,
	)

	statusCode, err := u.DeleteFood(foodId)

	if err != nil {
		return ctx.JSON(int(statusCode), echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(int(statusCode), echo.Map{
		"message": "item deleted successfully",
	})
}

func (h *FarmerHandler) GetOrdersHandler(ctx echo.Context) error {
	farmerId := ctx.Param("farmerId")

	u := usecase.NewOrderUseCase(
		h.dbRepo,
		h.callAnswerApi,
		h.callFrom,
		h.twilioRepo,
	)

	orders, statusCode, err := u.GetOrdersByFarmerId(farmerId)

	if err != nil {
		return ctx.JSON(int(statusCode), echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(int(statusCode), orders)
}

func (h *FarmerHandler) DeleteOrderHandler(ctx echo.Context) error {
	orderId := ctx.Param("orderId")
	u := usecase.NewOrderUseCase(
		h.dbRepo,
		h.callAnswerApi,
		h.callFrom,
		h.twilioRepo,
	)
	statusCode, err := u.DeleteOrder(orderId)

	if err != nil {
		return ctx.JSON(int(statusCode), echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(int(statusCode), echo.Map{
		"message": "order deleted successfully",
	})

}

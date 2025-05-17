package routes

import (
	"github.com/bhushn-aruto/krushi-sayak-http-server/config"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/infra/server/handlers"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/repo"
	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo, config *config.Config, dbRepo repo.DatabaseRepo, storageRepo repo.StorageRepo, twilioRepo repo.TwilioRepo) {
	farmerHandler := handlers.NewFormerHandler(dbRepo, storageRepo, config.CallAnswerApi, config.CallFrom, twilioRepo)
	buyerHandler := handlers.NewBuyerHandler(dbRepo, config.CallAnswerApi, config.CallFrom, twilioRepo)
	alertHandler := handlers.NewAlertHandler()
	orderNotifyHandler := handlers.NewOrderNotifyHandler()

	e.Static("/public", "public")

	farmer := e.Group("/farmer")
	buyer := e.Group("/buyer")
	alert := e.Group("/alert")
	notification := e.Group("/notify")

	farmer.POST("/signup", farmerHandler.SignUpHandler)
	farmer.POST("/login", farmerHandler.LoginHandler)

	buyer.POST("/signup", buyerHandler.SignUpHandler)
	buyer.POST("/login", buyerHandler.LoginHandler)

	farmer.GET("/get/category/:id", farmerHandler.GetFoodVariantsHandler)
	farmer.DELETE("/delete/category/:id", farmerHandler.DeleteFoodVariantHandler)

	farmer.POST("/create/item", farmerHandler.CreateFoodHandler)
	farmer.GET("/get/item/:id", farmerHandler.GetFoodsHandler)
	farmer.DELETE("/delete/item/:id", farmerHandler.DeleteFoodHandler)

	farmer.GET("/get/orders/:farmerId", farmerHandler.GetOrdersHandler)
	farmer.DELETE("/delete/order/:orderId", farmerHandler.DeleteOrderHandler)

	alert.POST("/moisture/high", alertHandler.MoistureHighAlertHandler)
	alert.POST("/moisture/low", alertHandler.MoistureLowAlertHandler)

	buyer.GET("/get/items", buyerHandler.GetAllFoodsHandler)
	buyer.POST("/place/order", buyerHandler.CreateOrderHandler)

	notification.POST("/order", orderNotifyHandler.OrderNotifyHandler)

}

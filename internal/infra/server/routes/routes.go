package routes

import (
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/infra/server/handlers"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/repo"
	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo, dbRepo repo.DatabaseRepo, storageRepo repo.StorageRepo) {
	farmerHandler := handlers.NewFormerHandler(dbRepo, storageRepo)
	buyerHandler := handlers.NewBuyerHandler(dbRepo)
	alertHandler := handlers.NewAlertHandler()

	e.Static("/public", "public")

	farmer := e.Group("/farmer")
	buyer := e.Group("/buyer")
	alert := e.Group("/alert")

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

}

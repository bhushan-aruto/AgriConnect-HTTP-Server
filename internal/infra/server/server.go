package server

import (
	"github.com/bhushn-aruto/krushi-sayak-http-server/config"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/infra/postgres"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/infra/server/routes"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/infra/storage"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/infra/twilio_app"
	"github.com/labstack/echo"
)

func StartApp(conf *config.Config) {
	e := echo.New()

	dbConn := postgres.NewDatabase(conf.DatabaseUrl)

	dbConn.CheckDatabase()

	defer dbConn.CloseConnection()

	dbRepo := postgres.NewPostgresRepo(dbConn.Pool)

	dbRepo.Init()

	storageRepo := storage.NewStorageRepo(
		"./public/variants",
		"./public/items",
	)

	storageRepo.Init()

	twilioClient := twilio_app.NewTwilioClient()

	trepo := twilio_app.NewTwilioRepo(twilioClient)

	routes.InitRoutes(e, conf, dbRepo, storageRepo, trepo)

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}

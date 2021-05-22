package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"time"
	useCase "wrk-connector/src/product/application/use-case"
	"wrk-connector/src/register/infrastructure/persistence/postgres"
	log "wrk-connector/src/shared/infrastructure/config"
	"wrk-connector/src/shared/infrastructure/config/persistence/gorm"
	version "wrk-connector/src/version/application/use_case"
)

func main() {
	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())
	e.HideBanner = true
	version.NewHealthHandler(e)

	postgresConnect := gorm.NewPostgresConnection()
	registerRepo := postgres.NewRegisterRepository(postgresConnect)

	sendToCreateProduct := useCase.NewSendToCreate(registerRepo)
	_ = sendToCreateProduct.Do()

	log.Info("Starting server...")
	portServer := os.Getenv("PORT_SERVER")
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", portServer),
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(server))

}

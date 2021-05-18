package main

import (
  "fmt"
  "github.com/labstack/echo/v4"
  echoMiddleware "github.com/labstack/echo/v4/middleware"
  "net/http"
  "os"
  "time"
  log "wrk-connector/src/shared/infrastructure/config"
  version "wrk-connector/src/version/application/use_case"
)

func main() {
  e := echo.New()
  e.Use(echoMiddleware.Recover())
  e.Use(echoMiddleware.CORS())
  e.HideBanner = true
  version.NewHealthHandler(e)

  log.Info("Starting server...")
  portServer := os.Getenv("PORT_SERVER")
  server := &http.Server{
    Addr:         fmt.Sprintf(":%s", portServer),
    ReadTimeout:  3 * time.Minute,
    WriteTimeout: 3 * time.Minute,
  }

  e.Logger.Fatal(e.StartServer(server))
}
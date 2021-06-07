package main

import (
  "fmt"
  "github.com/labstack/echo/v4"
  echoMiddleware "github.com/labstack/echo/v4/middleware"
  "net/http"
  "os"
  "time"
  useCase "wrk-connector/src/product/application/use-case"
  "wrk-connector/src/product/infrastructure/persistence/oracle"
  "wrk-connector/src/register/infrastructure/persistence/postgres"
  "wrk-connector/src/shared/domain/constant"
  "wrk-connector/src/shared/infrastructure/config/persistence/gorm"
  "wrk-connector/src/shared/infrastructure/queue/kafka"

  log "wrk-connector/src/shared/infrastructure/config"
  databaseSql "wrk-connector/src/shared/infrastructure/config/persistence/godror"
  _interface "wrk-connector/src/shared/infrastructure/config/persistence/interface"
  version "wrk-connector/src/version/application/use_case"
)

const (
  POSTGRES = "postgres"
  ORACLE   = "godror"
)

func main() {
  e := echo.New()
  e.Use(echoMiddleware.Recover())
  e.Use(echoMiddleware.CORS())
  e.HideBanner = true
  version.NewHealthHandler(e)

  postgresDbParameters := _interface.DbParameters{
    Host:     os.Getenv("POSTGRES_HOST"),
    Port:     os.Getenv("POSTGRES_PORT"),
    User:     os.Getenv("POSTGRES_USER"),
    Password: os.Getenv("POSTGRES_PASSWORD"),
    DbName:   os.Getenv("POSTGRES_DATABASE_NAME"),
  }

  oracleDbParameters := _interface.DbParameters{
    Host:     os.Getenv("ORACLE_DB_CONNECTOR_HOST"),
    Port:     os.Getenv("ORACLE_DB_CONNECTOR_PORT"),
    User:     os.Getenv("ORACLE_DB_CONNECTOR_USER"),
    Password: os.Getenv("ORACLE_DB_CONNECTOR_PASSWORD"),
    DbName:   os.Getenv("ORACLE_DB_CONNECTOR_DATABASE_NAME"),
  }

  oracleConnect := databaseSql.NewOracleConnection(ORACLE, oracleDbParameters)
  postgresConnect := gorm.NewConnection(POSTGRES, postgresDbParameters)

  producerMassage := kafka.NewKafkaProducer(
    kafka.GetDialer(),
    constant.KafkaBrokers...
    )

  registerRepo := postgres.NewRegisterRepository(postgresConnect)
  productRepo := oracle.NewProductRepository(oracleConnect)
  sendToCreateProduct := useCase.NewProductSendToCreate(
    registerRepo, productRepo,
    )
  sendToCreateProductQueue := useCase.NewProductSendToCreateQueue(
    sendToCreateProduct,
    producerMassage,
    )
  _, _ = sendToCreateProductQueue.Do()

  log.Info("Starting server...")
  portServer := os.Getenv("PORT_SERVER")
  server := &http.Server{
    Addr:         fmt.Sprintf(":%s", portServer),
    ReadTimeout:  3 * time.Minute,
    WriteTimeout: 3 * time.Minute,
  }
  e.Logger.Fatal(e.StartServer(server))

}

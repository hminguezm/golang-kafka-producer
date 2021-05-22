package gorm

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "os"
  log "wrk-connector/src/shared/infrastructure/config"
)

var connection *gorm.DB

type PostgresConnection struct {
  url  string
}

func NewPostgresConnection() *PostgresConnection {
  const (
    // UrlFormat "postgres://user:password@localhost:26257/dbName?sslmode=disable"
    UrlFormat = "postgresql://%v:%v@%v:%v/%v?sslmode=disable"
  )
  url := fmt.Sprintf(
    UrlFormat,
    os.Getenv("POSTGRES_USER"),
    os.Getenv("POSTGRES_PASSWORD"),
    os.Getenv("POSTGRES_HOST"),
    os.Getenv("POSTGRES_PORT"),
    os.Getenv("POSTGRES_DATABASE_NAME"),
    )
  if url == "" {
    panic("Error: url cannot be empty")
  }
  return &PostgresConnection{
    url:  url,
  }
}

func (r *PostgresConnection) OpenConnection() *gorm.DB {
  var err error
  if connection == nil || !isAlive() {
    connection, err = gorm.Open("postgres", r.url)
    if err != nil {
      log.WithError(err).Error("error trying to connect to DB")
    } else {
       log.Debug("connected to DB")
    }
  }

  // connection.AutoMigrate(&model.Register{})

  return connection
}

func (r *PostgresConnection) CloseConnection() {
  if err := connection.Close(); err != nil {
    log.WithError(err).Error("error trying to close connection")
  } else {
    log.Debug("connection closed")
  }
}

func isAlive() bool {
  if err := connection.DB().Ping(); err != nil {
     log.WithError(err).Error("error trying to ping DB")
    return false
  }
  return true
}



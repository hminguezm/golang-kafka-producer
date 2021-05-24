package database_sql

import (
  "database/sql"
  "fmt"
  _ "github.com/godror/godror"
  log "wrk-connector/src/shared/infrastructure/config"
  _interface "wrk-connector/src/shared/infrastructure/config/persistence/interface"
)

var connection *sql.DB

type OracleConnection struct {
  dbParameters _interface.DbParameters
  dialect      string
  url          string
}

func NewOracleConnection(dialect string, dbParameters _interface.DbParameters) *OracleConnection {
  url := fmt.Sprintf(
    "%s/%s@%s:%s/%s",
    dbParameters.User,
    dbParameters.Password,
    dbParameters.Host,
    dbParameters.Port,
    dbParameters.DbName,
  )
  if url == "" {
    panic("Error: url cannot be empty")
  }
  return &OracleConnection{
    dialect: dialect,
    url:     url,
  }
}

func (r *OracleConnection) OpenConnection() *sql.DB {
  var err error
  if connection == nil || !isAlive() {
    connection, err = sql.Open(r.dialect, r.url)

    if err != nil {
      log.WithError(err).Error("error trying to connect to %s DB", r.dialect)
      panic("shutdown app")
    } else {
      log.Debug("connected to %s DB", r.dialect)
    }
  }

  return connection
}

func (r *OracleConnection) CloseConnection() {
  if err := connection.Close(); err != nil {
    log.WithError(err).Error("error trying to close connection to database-sql DB")
  } else {
    log.Debug("connection closed to database-sql DB")
  }
}

func isAlive() bool {
  if err := connection.Ping(); err != nil {
    log.WithError(err).Error("error trying to ping database-sql DB")

    return false
  }

  return true
}

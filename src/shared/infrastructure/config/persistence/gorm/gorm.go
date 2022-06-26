package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "wrk-connector/src/shared/infrastructure/config"
	_interface "wrk-connector/src/shared/infrastructure/config/persistence/interface"
)

var connection *gorm.DB

type Connection struct {
	dbParameters _interface.DbParameters
	dialect      string
	url          string
}

func NewConnection(dialect string, dbParameters _interface.DbParameters) *Connection {
	url := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbParameters.Host,
		dbParameters.User,
		dbParameters.DbName,
		dbParameters.Password,
	)
	if url == "" {
		panic("Error: url cannot be empty")
	}
	return &Connection{
		dialect: dialect,
		url:     url,
	}
}

func (r *Connection) OpenConnection() *gorm.DB {
	var err error
	if connection == nil || !isAlive() {
		connection, err = gorm.Open(r.dialect, r.url)
		if err != nil {
			log.WithError(err).Error("error trying to connect to %s DB", r.dialect)
			panic("shutdown app")
		} else {
			log.Debug("connected to %s DB", r.dialect)
		}
	}

	return connection
}

func (r *Connection) CloseConnection() {
	if err := connection.Close(); err != nil {
		log.WithError(err).Error("error trying to close connection to %s DB", r.dialect)
	} else {
		log.Debug("connection closed to %s DB", r.dialect)
	}
}

func isAlive() bool {
	if err := connection.DB().Ping(); err != nil {
		log.WithError(err).Error("error trying to ping to DB")

		return false
	}

	return true
}

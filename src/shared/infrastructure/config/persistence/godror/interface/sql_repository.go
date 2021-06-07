package _interface

import (
	"database/sql"
)

type SqlRepository interface {
	OpenConnection() *sql.DB
	CloseConnection()
}

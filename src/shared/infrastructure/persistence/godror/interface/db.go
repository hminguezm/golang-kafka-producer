package _interface

import "database/sql"

type DbQuery interface {
	ExecuteQuery(conn *sql.DB, query string) ([][]interface{}, error)
}

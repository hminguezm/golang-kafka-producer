package oracle

import (
  "database/sql"
  "fmt"
  "wrk-connector/src/product/infrastructure/persistence/oracle/model"
  log "wrk-connector/src/shared/infrastructure/config"
  _interface "wrk-connector/src/shared/infrastructure/config/persistence/database-sql/interface"
  oracleDatabase "wrk-connector/src/shared/infrastructure/persistence/database-sql"
)

type ProductRepository struct {
	connection _interface.SqlRepository
}

func NewProductRepository(connection _interface.SqlRepository) *ProductRepository {
	return &ProductRepository{
		connection: connection,
	}
}

func (r *ProductRepository) FindAll() ([]*model.Product, error) {
	// product := model.Product{}
	conn := r.connection.OpenConnection()
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close database: %v", err))
		}
	}(conn)

  var result [][]interface{}
  query := "SELECT * FROM CON_PRODUCTS"
  result, _ = oracleDatabase.ExecuteQuery(conn, query)

	log.Debug("result %s", result)

	return nil, nil
}

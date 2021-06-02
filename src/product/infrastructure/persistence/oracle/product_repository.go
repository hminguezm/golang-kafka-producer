package oracle

import (
  "database/sql"
  "fmt"
  "wrk-connector/src/product/domain/entity"
  s "wrk-connector/src/product/domain/service"
  "wrk-connector/src/product/infrastructure/mapper"
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

func (r *ProductRepository) FindToCreate(dateExec string) ([]*entity.Product, error) {
  conn := r.connection.OpenConnection()
  defer func(conn *sql.DB) {
    err := conn.Close()
    if err != nil {
      panic(fmt.Sprintf("failed to close database: %v", err))
    }
  }(conn)

  query := s.GetProductToSendSql(dateExec)
  dbProducts, err := oracleDatabase.ExecuteQuery(conn, query)
  if nil != err {
    log.WithError(err).Error("error executing products query")
    return nil, err
  }

  products := make([]*entity.Product, 0, len(dbProducts))
  for _, dbProduct := range dbProducts {
    product := new(entity.Product)
    product = mapper.MapProductFromDBResponse(dbProduct...)
    products = append(products, product)
  }

  return products, nil
}

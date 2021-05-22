package postgres

import (
  "wrk-connector/src/register/infrastructure/persistence/postgres/model"
  log "wrk-connector/src/shared/infrastructure/config"
  _interface "wrk-connector/src/shared/infrastructure/config/persistence/gorm/interface"
)

type registerRepository struct {
  connection _interface.GormRepository
}

func NewRegisterRepository(connection _interface.GormRepository) *registerRepository {
  return &registerRepository{
    connection: connection,
  }
}

func (r *registerRepository) CreateRegister() error {
  register := model.Register{}
  conn := r.connection.OpenConnection()
  conn = conn.Create(&register)
  if conn.Error != nil {
    log.WithError(conn.Error).Error("error inserting record")
    return conn.Error
  }

  return nil
}

func (r *registerRepository) GetLastRegister() (*model.Register, error) {
  var register model.Register
  conn := r.connection.OpenConnection()

  result := conn.Last(&register)
  if result.Error != nil {
    log.WithError(result.Error).Error("error executing find query")
    return nil, result.Error
  }

  registerResult := model.Register{
    ID: register.ID,
    CreatedAt: register.CreatedAt,
  }
  log.Debug("query finished successfully,", result.RowsAffected, "records were found")

  return &registerResult, nil
}


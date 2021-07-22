package postgres

import (
  "fmt"
  "github.com/jinzhu/gorm"
  uuid "github.com/satori/go.uuid"
  "time"
  "wrk-connector/src/register/infrastructure/mapper"
  "wrk-connector/src/register/infrastructure/persistence/dto"
  "wrk-connector/src/register/infrastructure/persistence/postgres/model"
  "wrk-connector/src/shared/domain/service"
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
  var register model.Register
  conn := r.connection.OpenConnection()
  conn = conn.AutoMigrate(&register)
  defer func(conn *gorm.DB) {
    err := conn.Close()
    if err != nil {
      panic(fmt.Sprintf("failed to close database: %v", err))
    }
  }(conn)

  conn = conn.Create(&register)
  if conn.Error != nil {
    log.WithError(conn.Error).Error("error inserting record")
    return conn.Error
  }

  log.Info("query finished successfully inserting record")

  return nil
}

func (r *registerRepository) GetLastRegister() *dto.RegisterGetLastedDTO {
  var register model.Register
  conn := r.connection.OpenConnection()
  conn = conn.AutoMigrate(&register)
  defer func(conn *gorm.DB) {
    err := conn.Close()
    if err != nil {
      panic(fmt.Sprintf("failed to close database: %v", err))
    }
  }(conn)

  result := conn.Last(&register)

  if result.RowsAffected == 0 {
    register.ID = uuid.NewV4()
    register.CreatedAt = service.ConvertToTime(time.Now())
    register.UpdatedAt = service.ConvertToTime(time.Now())
  }

  registerResult := mapper.MapRegisterFromDBResponse(register)

  return registerResult
}

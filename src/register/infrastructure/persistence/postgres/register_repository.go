package postgres

import (
  "fmt"
  "github.com/jinzhu/gorm"
  "wrk-connector/src/register/infrastructure/mapper"
  "wrk-connector/src/register/infrastructure/persistence/dto"
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
	conn := r.connection.OpenConnection()
	defer func(conn *gorm.DB) {
		err := conn.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close database: %v", err))
		}
	}(conn)

	register := model.Register{}
	conn = conn.Create(&register)
	if conn.Error != nil {
		log.WithError(conn.Error).Error("error inserting record")
		return conn.Error
	}

	log.Debug("query finished successfully inserting record")

	return nil
}

func (r *registerRepository) GetLastRegister() (*dto.RegisterGetLastedDTO, error) {
	conn := r.connection.OpenConnection()
	defer func(conn *gorm.DB) {
		err := conn.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close database: %v", err))
		}
	}(conn)

	var register model.Register
	result := conn.Last(&register)
	if result.Error != nil {
		log.WithError(result.Error).Error("error executing find query")
		return nil, result.Error
	}

	registerResult := mapper.MapRegisterFromDBResponse(register)
	log.Debug("query finished successfully,", result.RowsAffected, "records were found")

	return registerResult, nil
}

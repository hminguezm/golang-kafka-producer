package repository

import (
  "wrk-connector/src/register/infrastructure/persistence/dto"
)

type Register interface {
	CreateRegister() error
	GetLastRegister() *dto.RegisterGetLastedDTO
}

package repository

import "wrk-connector/src/register/infrastructure/persistence/postgres/model"

type RegisterPersistence interface {
	CreateRegister() error
	GetLastRegister() (*model.Register, error)
}

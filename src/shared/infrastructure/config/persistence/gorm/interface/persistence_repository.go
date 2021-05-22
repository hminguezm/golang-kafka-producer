package _interface

import (
	"github.com/jinzhu/gorm"
)

type GormRepository interface {
	OpenConnection() *gorm.DB
	CloseConnection()
}

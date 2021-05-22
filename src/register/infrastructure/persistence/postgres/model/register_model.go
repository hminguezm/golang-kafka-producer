package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type Register struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;type:uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (base *Register) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

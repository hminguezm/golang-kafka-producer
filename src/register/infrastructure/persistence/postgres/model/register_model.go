package model

import (
  "github.com/satori/go.uuid"
  "time"
  "wrk-connector/src/shared/domain/constant"
  "wrk-connector/src/shared/domain/service"
)

type Register struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;type:uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (base *Register) BeforeCreate() (err error) {
  base.ID = uuid.NewV4()
  base.UpdatedAt = service.ConvertToTime(time.Now().Format(constant.TimeFormat))

  return
}

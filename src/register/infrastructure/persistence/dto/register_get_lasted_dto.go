package dto

import (
  "github.com/google/uuid"
)

type RegisterGetLastedDTO struct {
  ID        uuid.UUID `json:"id"`
  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"update_at"`
}

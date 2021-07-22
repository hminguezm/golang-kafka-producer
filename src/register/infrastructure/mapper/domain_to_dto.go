package mapper

import (
  "github.com/google/uuid"
  "wrk-connector/src/register/infrastructure/persistence/dto"
  "wrk-connector/src/register/infrastructure/persistence/postgres/model"
  "wrk-connector/src/shared/domain/constant"
)

func MapRegisterFromDBResponse(data model.Register) *dto.RegisterGetLastedDTO {
  result := dto.RegisterGetLastedDTO{
    ID: uuid.UUID(data.ID),
    CreatedAt: data.CreatedAt.Format(constant.TimeFormat),
  }

  return &result
}

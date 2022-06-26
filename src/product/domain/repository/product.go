package repository

import (
  "wrk-connector/src/product/domain/entity"
)

type Product interface {
	FindToCreate(dateExec string) ([]*entity.Product, error)
}

package _interface

import "wrk-connector/src/product/domain/entity"

type ProductSendToCreate interface {
  Do() ([]*entity.Product, error)
}

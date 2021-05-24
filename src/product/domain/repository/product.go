package repository

import "wrk-connector/src/product/infrastructure/persistence/oracle/model"

type Product interface {
	FindAll() ([]*model.Product, error)
}

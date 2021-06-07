package use_case

import (
  "os"
  "wrk-connector/src/product/domain/entity"
  repoP "wrk-connector/src/product/domain/repository"
  repoR "wrk-connector/src/register/domain/repository"
  "wrk-connector/src/register/infrastructure/persistence/dto"
  log "wrk-connector/src/shared/infrastructure/config"
)

type ProductSendToCreate struct {
  registerRepository repoR.Register
  productRepository  repoP.Product
}

func NewProductSendToCreate(
  registerRepository repoR.Register,
  productRepository repoP.Product,
) *ProductSendToCreate {
  return &ProductSendToCreate{
    registerRepository: registerRepository,
    productRepository:  productRepository,
  }
}

func (u *ProductSendToCreate) Do() ([]*entity.Product, error) {
  log.Debug("SendToCreate!")

  var register *dto.RegisterGetLastedDTO
  register, err := u.registerRepository.GetLastRegister()
  if err != nil {
    return nil, err
  }

  log.Info("register CreatedAt: %s: ", register.CreatedAt)
  productsDB, err := u.productRepository.FindToCreate(register.CreatedAt)
  if nil != err {
    log.WithError(err).Fatal("error could not get product from db")
    os.Exit(0)
  }

  log.Info("create new register")
  if len(productsDB) == 0 {
    log.Info("Could not found products")
    os.Exit(0)
  }

  err = u.registerRepository.CreateRegister()
  if err != nil {
    return nil, err
  }

  return productsDB, nil
}

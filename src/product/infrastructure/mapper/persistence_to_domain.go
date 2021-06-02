package mapper

import (
  "wrk-connector/src/product/domain/entity"
  "wrk-connector/src/product/infrastructure/mapper/constant"
  "wrk-connector/src/shared/domain/service"
)

func MapProductFromDBResponse(data ...interface{}) *entity.Product {
  product := entity.Product{}
  product.ProductSku = service.ConvertToString(data[constant.ProductSku])
  product.ProductSkuParent = service.ConvertToString(data[constant.ProductSkuParent])
  product.VariantId = service.ConvertToString(data[constant.VariantId])
  product.CategoryCode = service.ConvertToString(data[constant.CategoryCode])
  product.Tittle = service.ConvertToString(data[constant.Tittle])
  product.ShortDescription = service.ConvertToString(data[constant.ShortDescription])
  product.LongDescription = service.ConvertToString(data[constant.LongDescription])
  product.Brand = service.ConvertToString(data[constant.Brand])
  product.SellerId = service.ConvertToInt(data[constant.SellerId])
  product.OffersId = service.ConvertToInt(data[constant.OffersId])
  product.PartNumber = service.ConvertToString(data[constant.PartNumber])
  product.Attributes = service.ConvertToString(data[constant.Attributes])
  product.Media = service.ConvertToString(data[constant.Media])
  product.CreatedAt = service.ConvertToTime(data[constant.CreatedAt])
  product.UpdatedAt = service.ConvertToTime(data[constant.UpdatedAt])

  return &product
}

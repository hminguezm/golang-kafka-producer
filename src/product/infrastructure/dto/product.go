package dto

import "time"

type ProductDTO struct {
  ProductSku       string    `json:"product_sku"`
  ProductSkuParent string    `json:"product_sku_parent"`
  VariantId        string    `json:"variant_id"`
  Categories       string    `json:"categories"`
  Tittle           string    `json:"tittle"`
  ShortDescription string    `json:"short_description"`
  LongDescription  string    `json:"long_description"`
  BRAND            string    `json:"brand"`
  SellerId         uint      `json:"seller_id"`
  OffersId         uint      `json:"offers_id"`
  PartNumber       string    `json:"part_number"`
  Attributes       string    `json:"attributes"`
  Media            string    `json:"media"`
  CreatedAt        time.Time `json:"created_at"`
  UpdatedAt        time.Time `json:"updated_at"`
}

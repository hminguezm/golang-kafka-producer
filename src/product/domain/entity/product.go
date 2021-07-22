package entity

import "time"

type Product struct {
  ProductSku       string    `json:"product_sku"`
  ProductSkuParent string    `json:"product_sku_parent"`
  VariantId        string    `json:"variant_id"`
  CategoryCode     string    `json:"category_code"`
  Tittle           string    `json:"tittle"`
  ShortDescription string    `json:"short_description"`
  LongDescription  string    `json:"long_description"`
  Brand            string    `json:"brand"`
  SellerId         int       `json:"seller_id"`
  OffersId         int       `json:"offers_id"`
  PartNumber       string    `json:"part_number"`
  Attributes       string    `json:"attributes"`
  Media            string    `json:"media"`
  Sync             int       `json:"sync"`
  CreatedAt        time.Time `json:"created_at"`
  UpdatedAt        time.Time `json:"updated_at"`
}

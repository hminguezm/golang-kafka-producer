package model

import (
	"time"
)

type Product struct {
	ProductsId       uint      `json:"PRODUCTS_ID"`
	FileId           uint      `json:"FILE_ID"`
	VariantId        string    `json:"VARIANT_ID"`
	VariantCorr      uint      `json:"VARIANT_CORR"`
	PARTNUMBER       string    `json:"PARTNUMBER"`
	THUMBNAIL        string    `json:"THUMBNAIL"`
	FULLIMAGE        string    `json:"FULLIMAGE"`
	NAME             string    `json:"NAME"`
	SHORTDESCRIPTION string    `json:"SHORTDESCRIPTION"`
	LONGDESCRIPTION  string    `json:"LONGDESCRIPTION"`
	CategoryCode     string    `json:"CATEGORY_CODE"`
	CATEGORY         string    `json:"CATEGORY"`
	BRAND            string    `json:"BRAND"`
	ACTIVATED        string    `json:"ACTIVATED"`
	ItemCode         string    `json:"ITEM_CODE"`
	CURRENCY         string    `json:"CURRENCY"`
	STATUS           string    `json:"STATUS"`
	CreateDate       time.Time `json:"CREATE_DATE"`
	LastDate         time.Time `json:"LAST_DATE"`
	PRICE            float64   `json:"PRICE"`
	OriginPrice      float64   `json:"ORIGIN_PRICE"`
	SYNC             string    `json:"SYNC"`
	SellerId         uint      `json:"SELLER_ID"`
	ProductSku       string    `json:"PRODUCT_SKU"`
	ProductSkuParent string    `json:"PRODUCT_SKU_PARENT"`
}

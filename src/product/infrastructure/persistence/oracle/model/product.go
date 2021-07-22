package model

import (
	"time"
)

type Product struct {
	ProductsId       uint      `json:"PRODUCTS_ID"`
	FileId           uint      `json:"FILE_ID"`
	VariantId        string    `json:"VARIANT_ID"`
	VariantCorr      uint      `json:"VARIANT_CORR"`
	Partnumber       string    `json:"PARTNUMBER"`
	Thumbnail        string    `json:"THUMBNAIL"`
	Fullimage        string    `json:"FULLIMAGE"`
	Name             string    `json:"NAME"`
	ShortDescription string    `json:"SHORTDESCRIPTION"`
	LongDescription  string    `json:"LONGDESCRIPTION"`
	CategoryCode     string    `json:"CATEGORY_CODE"`
	Category         string    `json:"CATEGORY"`
	Brand            string    `json:"BRAND"`
	Activated        string    `json:"ACTIVATED"`
	ItemCode         string    `json:"ITEM_CODE"`
	Currency         string    `json:"CURRENCY"`
	Status           string    `json:"STATUS"`
	CreateDate       time.Time `json:"CREATE_DATE"`
	LastDate         time.Time `json:"LAST_DATE"`
	Price            float64   `json:"PRICE"`
	OriginPrice      float64   `json:"ORIGIN_PRICE"`
	Sync             string    `json:"SYNC"`
	SellerId         uint      `json:"SELLER_ID"`
	ProductAttribute string    `json:"ProductAttribute"`
	ProductSkuParent string    `json:"PRODUCT_SKU_PARENT"`
}

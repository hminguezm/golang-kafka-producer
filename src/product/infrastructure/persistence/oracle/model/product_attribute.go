package model

import "time"

type ProductAttribute struct {
	Partnumber    string    `json:"PARTNUMBER"`
	Name          string    `json:"NAME"`
	Value         string    `json:"VALUE"`
	AttributeType string    `json:"ATTRIBUTE_TYPE"`
	ListValue     string    `json:"LIST_VALUE"`
	Status        string    `json:"STATUS"`
	CreateDate    time.Time `json:"CREATE_DATE"`
	FileId        uint32    `json:"FILE_ID"`
	ProductSku    Product   `json:"PRODUCT_SKU"`
}

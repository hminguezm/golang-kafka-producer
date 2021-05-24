package model

import "time"

type ProductAttribute struct {
	PARTNUMBER    string    `json:"PARTNUMBER"`
	NAME          string    `json:"NAME"`
	VALUE         string    `json:"VALUE"`
	AttributeType string    `json:"ATTRIBUTE_TYPE"`
	ListValue     string    `json:"LIST_VALUE"`
	STATUS        string    `json:"STATUS"`
	CreateDate    time.Time `json:"CREATE_DATE"`
	FileId        uint32    `json:"FILE_ID"`
	ProductSku    string    `json:"PRODUCT_SKU"`
}

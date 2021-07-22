package model

import "time"

type Offer struct {
	OffersId                   uint      `json:"OffersId"`
	ProductSku                 string    `json:"ProductSku"`
	ParentPartnumber           string    `json:"ParentPartnumber"`
	MinShippingPrice           uint      `json:"MinShippingPrice"`
	MinShippingPriceAdditional uint      `json:"MinShippingPriceAdditional"`
	MinShippingZone            string    `json:"MinShippingZone"`
	MinShippingType            string    `json:"MinShippingType"`
	Price                      float64   `json:"Price"`
	TotalPrice                 float64   `json:"TotalPrice"`
	PriceAdditionalInfo        string    `json:"PriceAdditionalInfo"`
	Quantity                   uint16    `json:"Quantity"`
	Description                string    `json:"Description"`
	StateCode                  uint8     `json:"StateCode"`
	ShopId                     uint16    `json:"ShopId"`
	ShopName                   string    `json:"SHOP_NAME"`
	Professional               string    `json:"Professional"`
	Premium                    string    `json:"Premium"`
	LogisticClass              string    `json:"LogisticClass"`
	Activated                  string    `json:"Activated"`
	FavoriteRank               string    `json:"FavoriteRank"`
	Channels                   string    `json:"Channels"`
	Deleted                    string    `json:"Deleted"`
	OriginPrice                uint      `json:"OriginPrice"`
	DiscountStartDate          time.Time `json:"DiscountStartDate"`
	DiscountEndDate            time.Time `json:"DiscountEndDate"`
	AvailableStartDate         time.Time `json:"AvailableStartDate"`
	AvailableEndDate           time.Time `json:"AvailableEndDate"`
	DiscountPrice              float64   `json:"DiscountPrice"`
	CurrencyIsoCode            string    `json:"CurrencyIsoCode"`
	DiscountRanges             string    `json:"DiscountRanges"`
	LeadtimeToShip             string    `json:"LeadtimeToShip"`
	Timpominimumdespacho       string    `json:"Timpominimumdespacho"`
	Status                     string    `json:"Status"`
	CreateDate                 time.Time `json:"CreateDate"`
	LastDate                   time.Time `json:"LastDate"`
	ConProductProductSku       string    `json:"ConProductProductSku"`
}

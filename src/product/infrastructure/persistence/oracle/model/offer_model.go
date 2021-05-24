package model

import "time"

type Offer struct {
	OffersId                   uint      `json:"OFFERS_ID"`
	ProductSku                 string    `json:"PRODUCT_SKU"`
	ParentPartnumber           string    `json:"PARENT_PARTNUMBER"`
	MinShippingPrice           uint      `json:"MIN_SHIPPING_PRICE"`
	MinShippingPriceAdditional uint      `json:"MIN_SHIPPING_PRICE_ADDITIONAL"`
	MinShippingZone            string    `json:"MIN_SHIPPING_ZONE"`
	MinShippingType            string    `json:"MIN_SHIPPING_TYPE"`
	Price                      float64   `json:"PRICE"`
	TotalPrice                 float64   `json:"TOTAL_PRICE"`
	PriceAdditionalInfo        string    `json:"PRICE_ADDITIONAL_INFO"`
	Quantity                   uint16    `json:"QUANTITY"`
	Description                string    `json:"DESCRIPTION"`
	StateCode                  uint8     `json:"STATE_CODE"`
	ShopId                     uint16    `json:"SHOP_ID"`
	ShopName                   string    `json:"SHOP_NAME"`
	Professional               string    `json:"PROFESSIONAL"`
	Premium                    string    `json:"PREMIUM"`
	LogisticClass              string    `json:"LOGISTIC_CLASS"`
	Activated                  string    `json:"ACTIVATED"`
	FavoriteRank               string    `json:"FAVORITE_RANK"`
	Channels                   string    `json:"CHANNELS"`
	Deleted                    string    `json:"DELETED"`
	OriginPrice                uint      `json:"ORIGIN_PRICE"`
	DiscountStartDate          time.Time `json:"DISCOUNT_START_DATE"`
	DiscountEndDate            time.Time `json:"DISCOUNT_END_DATE"`
	AvailableStartDate         time.Time `json:"AVAILABLE_START_DATE"`
	AvailableEndDate           time.Time `json:"AVAILABLE_END_DATE"`
	DiscountPrice              float64   `json:"DISCOUNT_PRICE"`
	CurrencyIsoCode            string    `json:"CURRENCY_ISO_CODE"`
	DiscountRanges             string    `json:"DISCOUNT_RANGES"`
	LeadtimeToShip             string    `json:"LEADTIME_TO_SHIP"`
	TIMPOMINIMUMDESPACHO       string    `json:"TIMPOMINIMUMDESPACHO"`
	Status                     string    `json:"STATUS"`
	CreateDate                 time.Time `json:"CREATE_DATE"`
	LastDate                   time.Time `json:"LAST_DATE"`
	ConProductProductSku       string    `json:"CON_PRODUCT_PRODUCT_SKU"`
}

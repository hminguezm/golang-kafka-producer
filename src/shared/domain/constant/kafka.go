package constant

import (
	"os"
	"strings"
)

var (
	KafkaBrokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
)

var KafkaTopics = map[string]string{
	"TOPIC_PRODUCTS_CREATE": os.Getenv("TOPIC_PRODUCTS_CREATE"),
	"TOPIC_PRODUCTS_UPDATE": os.Getenv("TOPIC_PRODUCTS_UPDATE"),
}

var EventName = map[string]string{
	"PRODUCTS_CREATED": "PRODUCTS_CREATED",
	"PRODUCTS_UPDATED": "PRODUCTS_UPDATED",
}

var EventDataFormat = map[string]string{
	"JSON": "JSON",
	"XML":  "XML",
	"TXT":  "TXT",
}

var EventType = map[string]string{
	"CREATE": "CREATE",
	"READ":   "READ",
	"UPDATE": "UPDATE",
	"DELETE": "DELETE",
}

var EventCountry = map[string]string{
	"CL": "CL",
	"PE": "PE",
}

var EventOrigin = "wrk-connector-producer"

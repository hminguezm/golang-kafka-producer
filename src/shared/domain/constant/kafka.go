package constant

import (
	"os"
	"strings"
)

var (
	KafkaBrokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
)

var KafkaTopics = map[string]string{
	"TOPIC_PRODUCT_CREATE": os.Getenv("KAFKA_TOPIC_PRODUCT_CREATE"),
	"TOPIC_PRODUCT_UPDATE": os.Getenv("KAFKA_TOPIC_PRODUCT_UPDATE"),
}

var EventName = map[string]string{
	"PRODUCT_CREATED": "PRODUCT_CREATED",
	"PRODUCT_UPDATED": "PRODUCT_UPDATED",
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

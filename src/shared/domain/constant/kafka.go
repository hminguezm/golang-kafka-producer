package constant

import (
  "os"
  "strings"
)

var (
  KafkaBrokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
)

var KafkaTopics = map[string] string{
  "TOPIC_OF03_INPUT": os.Getenv("TOPIC_OF03_INPUT"),
  "TOPIC_P21": os.Getenv("TOPIC_P21"),
}

var EventName = map[string] string{
  "MIRAKL_OF02": "MIRAKL_OF02",
}

var EventDataFormat = map[string] string{
  "JSON": "JSON",
  "XML": "XML",
  "TXT": "TXT",
}

var EventType = map[string] string{
  "CREATE": "CREATE",
  "READ": "READ",
  "UPDATE": "UPDATE",
  "DELETE": "DELETE",
}

var EventCountry = map[string] string{
  "CL": "CL",
  "PE": "PE",
}

var EventOrigin = "wrk-connector-producer"

package service

import "encoding/json"

func EntityToJson(entity interface{}) string {
	str, err := json.Marshal(entity)
	if err != nil {
		return "{}"
	}
	return string(str)
}

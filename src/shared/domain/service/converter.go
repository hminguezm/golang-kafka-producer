package service

import (
  "encoding/json"
  "reflect"
  "strconv"
  "time"
)

func ConvertToString(field interface{}) string {
	if field == nil {
		return ""
	}
	if data, ok := field.(string); ok {
		return data
	} else if data, ok := field.(int); ok {
		return strconv.Itoa(data)
	} else if data, ok := field.(int64); ok {
		return strconv.Itoa(int(data))
	}

	return ""
}

func ConvertNumberToBool(field interface{}) bool {
	if 1 == ConvertToInt64(field) {
		return true
	}

	return false
}

func ConvertToInt64(field interface{}) int64 {
	if field != nil {
		if s, ok := field.(string); ok {
			if i, err := strconv.ParseInt(s, 10, 64); err == nil {
				return i
			}
		}
		if i, ok := field.(int); ok {
			return int64(i)
		} else if i, ok := field.(int64); ok {
			return i
		}
	}

	return 0
}

func ConvertToInt(field interface{}) int {
	return int(ConvertToInt64(field))
}

func ConvertToFloat64(field interface{}) float64 {
	if field != nil {
		if f, ok := field.(float64); ok {
			return f
		} else if s, ok := field.(string); ok {
			f, _ := strconv.ParseFloat(s, 64)

			return f
		}
	}

	return 0.0
}

func ConvertToTime(field interface{}) time.Time {
	if field == nil {
		return time.Time{}
	}
	if data, ok := field.(time.Time); ok {
		return data
	}

	return time.Time{}
}

func IsNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func EntityToJson(entity interface{}) string {
  str, err := json.Marshal(entity)
  if err != nil {
    return "{}"
  }

  return string(str)
}

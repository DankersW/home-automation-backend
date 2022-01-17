package utils

import "reflect"

func ToString(item interface{}) string {
	if item == nil || reflect.TypeOf(item).Kind() != reflect.String {
		return ""
	}
	return item.(string)
}

func ToFloat32(value interface{}) float32 {
	if reflect.TypeOf(value).Kind() == reflect.Float64 {
		return float32(value.(float64))
	} else if reflect.TypeOf(value).Kind() == reflect.Int32 {
		return float32(value.(int32))
	} else {
		return 0.0
	}
}

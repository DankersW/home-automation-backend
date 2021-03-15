package main

import (
	"reflect"
)

func cast_to_float32(value interface{}) float32 {
	if reflect.TypeOf(value).Kind() == reflect.Float64 {
		return float32(value.(float64))
	} else if reflect.TypeOf(value).Kind() == reflect.Int32 {
		return float32(value.(int32))
	} else {
		return 0.0
	}
}

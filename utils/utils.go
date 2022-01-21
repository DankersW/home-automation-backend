package utils

import (
	"reflect"
	"unicode/utf8"
)

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

func RmFirstChar(str string) string {
	_, i := utf8.DecodeRuneInString(str)
	return str[i:]
}

func StrInSlice(str string, slice []string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

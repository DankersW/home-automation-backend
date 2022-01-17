package utils

import "reflect"

func ToString(item interface{}) string {
	if item == nil || reflect.TypeOf(item).Kind() != reflect.String {
		return ""
	}
	return item.(string)
}

package utils

import (
	"log"
	"reflect"
)

func ArrayMap(arr interface{}, cb func(item interface{}) interface{}) interface{} {
	parsedArr := ToInterfaceArray(arr)
	result := make([]interface{}, len(parsedArr))

	for i, item := range parsedArr {
		result[i] = cb(item)
	}

	return result
}

func ToInterfaceArray(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)

	if v.Kind() != reflect.Slice {
		log.Fatal("The arg must be a slice")
	}

	len := v.Len()
	result := make([]interface{}, len)

	for i := 0; i < len; i++ {
		result[i] = v.Index(i).Interface()
	}

	return result
}

package utils

import (
	"log"
	"reflect"
)

func ArrayMap(arr interface{}, cb func(item interface{}) interface{}, result interface{}) {
	parsedArr := ToInterfaceArray(arr)
	v := reflect.ValueOf(result).Elem()

	for _, item := range parsedArr {
		v.Set(reflect.Append(v, reflect.ValueOf(cb(item))))
	}
}

func ArrayIncludes(arr interface{}, item interface{}) bool {
	parsedArr := ToInterfaceArray(arr)

	// validate type of item of arr
	if reflect.ValueOf(parsedArr[0]).Kind() != reflect.ValueOf(item).Kind() {
		log.Fatal("ArrayIncludes: The type does not compatible")
	}

	for _, e := range parsedArr {
		if e == item {
			return true
		}
	}

	return false
}

func ToInterfaceArray(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)

	if v.Kind() != reflect.Slice {
		log.Fatal("ToInterfaceArray: The arg must be a slice")
	}

	arrLen := v.Len()
	result := make([]interface{}, arrLen)

	for i := 0; i < arrLen; i++ {
		result[i] = v.Index(i).Interface()
	}

	return result
}

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

func ArrayFilter(arr interface{}, cb func(item interface{}) bool, result interface{}) {
	parsedArr := ToInterfaceArray(arr)
	v := reflect.ValueOf(result).Elem()

	for _, item := range parsedArr {
		if !cb(item) {
			continue
		}

		v.Set(reflect.Append(v, reflect.ValueOf(item)))
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

func ArrayFind(arr interface{}, cb func(item interface{}) bool, item interface{}) bool {
	result := ToInterfaceArray([]interface{}{})

	ArrayFilter(arr, func(item interface{}) bool {
		return cb(item)
	}, &result)

	if len(result) == 0 {
		return false
	}

	v := reflect.ValueOf(item).Elem()
	v.Set(reflect.ValueOf(result[0]))

	return true
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

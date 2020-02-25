package utils

import (
	"fmt"
	"log"
	"reflect"
)

func MapKeys(m interface{}, keys interface{}) {
	mapKeys := ValidMap(m).MapKeys()
	keysVal := reflect.ValueOf(keys).Elem()

	for _, k := range mapKeys {
		keysVal.Set(reflect.Append(keysVal, reflect.ValueOf(k.Interface())))
	}
}

func MapValues(m interface{}, values interface{}) {
	var keys []string
	MapKeys(m, &keys)

	ArrayMap(keys, func(k interface{}) interface{} {
		//iter := v.MapRange()
		//values := make([]interface{}, v.Len())

		//fmt.Println(k)

		fmt.Println(reflect.ValueOf(m).MapIndex(reflect.ValueOf(k.(string))))

		return ""
	}, values)
}

func ValidMap(m interface{}) reflect.Value {
	v := reflect.ValueOf(m)

	if v.Kind() != reflect.Map {
		log.Fatal("The arg must be a Map")
	}

	return v
}

package utils

import (
	"reflect"
)

func ValidField(origin interface{}) {
	updates := make(map[string]interface{})

	vUser := reflect.ValueOf(origin)

	for i := 0; i < vUser.NumField(); i++ {
		val := vUser.Field(i)

		if val.IsValid() {
			updates[vUser.Type().Field(i).Name] = val
		}
	}
}

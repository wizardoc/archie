package utils

import (
	"reflect"
)

func CpStruct(origin interface{}, target interface{}) {
	originRealVal := reflect.ValueOf(origin).Elem()
	targetRealVal := reflect.ValueOf(target).Elem()

	for i := 0; i < originRealVal.NumField(); i++ {
		fVal := originRealVal.Field(i)
		fName := originRealVal.Type().Field(i).Name
		targetFVal := targetRealVal.FieldByName(fName)

		// 判断目标结构体是否有该元素
		if !targetFVal.IsValid() {
			continue
		}

		// 判断类型
		if fVal.Type() != targetFVal.Type() {
			continue
		}

		targetFVal.Set(fVal)
	}
}

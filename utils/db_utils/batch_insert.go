package db_utils

import (
	"archie/connection/postgres_conn"
	"archie/utils"
	"fmt"
	"reflect"
	"strings"
)

func BatchInsert(table string, heads []string, values interface{}) error {
	ifaceValues := utils.ToInterfaceArray(values)

	var strValues []string

	for _, val := range ifaceValues {
		resultVal := val
		reflectVal := reflect.ValueOf(val)
		var parsedVals []string
		var attachQuoteItem []string

		// 转 struct 到 slice
		if reflectVal.Kind() == reflect.Struct {
			for i := 0; i < reflectVal.NumField(); i++ {
				parsedVals = append(parsedVals, fmt.Sprintf("%v", reflectVal.Field(i).Interface()))
			}

			resultVal = parsedVals
		}

		// 统一处理 slice 拼接 SQL
		utils.ArrayMap(resultVal, func(item interface{}) interface{} {
			return fmt.Sprintf("'%s'", item)
		}, &attachQuoteItem)

		tupleStr := strings.Join(attachQuoteItem, ",")
		strValues = append(strValues, fmt.Sprintf("(%s)", tupleStr))
	}

	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES %s;",
		table,
		strings.Join(heads, ","),
		strings.Join(strValues, ","),
	)

	return postgres_conn.DB.Instance().Exec(sql).Error
}

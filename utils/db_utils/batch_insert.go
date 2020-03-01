package db_utils

import (
	"archie/connection/postgres_conn"
	"archie/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

func BatchInsert(table string, heads []string, values interface{}) error {
	ifaceValues := utils.ToInterfaceArray(values)
	var strValues []string
	var attachQuoteItem []string

	for _, val := range ifaceValues {
		utils.ArrayMap(val, func(item interface{}) interface{} {
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

	return postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		return db.Exec(sql).Error
	})
}

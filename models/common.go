package models

import (
	"archie/connection/postgres_conn"
)

func updateSig(model interface{}, name string, value interface{}) error {
	return postgres_conn.DB.Instance().Model(model).Update(name, value).Error
}

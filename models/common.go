package models

import (
	"archie/connection/postgres_conn"
)

type AssociationModel interface {
}

func updateSig(model interface{}, name string, value interface{}) error {
	return postgres_conn.DB.Instance().Model(model).Update(name, value).Error
}

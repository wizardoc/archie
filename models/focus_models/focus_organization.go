package focus_models

import (
	"archie/connection/postgres_conn"
)

type FocusOrganization struct {
	UserID         string `gorm:"type:uuid;primary_key"`
	OrganizationID string `gorm:"type:uuid;primary_key"`
}

func (fo *FocusOrganization) New() error {
	return postgres_conn.DB.Instance().Create(fo).Error
}

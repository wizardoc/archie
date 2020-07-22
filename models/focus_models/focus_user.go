package focus_models

import "archie/connection/postgres_conn"

type FocusUser struct {
	UserID      string `gorm:"type:uuid;primary_key"`
	FocusUserID string `gorm:"type:uuid;primary_key"`
}

func (fu *FocusUser) New() error {
	return postgres_conn.DB.Instance().Create(fu).Error
}

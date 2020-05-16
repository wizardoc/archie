package models

import (
	"archie/connection/postgres_conn"
	"fmt"
	"github.com/jinzhu/gorm"
)

type OrganizationPermission struct {
	PermissionID   string `gorm:"primary_key"`
	UserID         string `gorm:"primary_key"`
	OrganizationID string `gorm:"primary_key"`
}

func (op *OrganizationPermission) New(permission int) error {
	return postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		p := Permission{Value: permission}
		findPermission := Permission{}

		if err := p.Find(&findPermission); err != nil {
			return err
		}

		fmt.Println(findPermission)

		if findPermission.ID == "" {
			return fmt.Errorf("Cannot find permission that value is %d. ", permission)
		}

		op.PermissionID = findPermission.ID

		return db.Create(op).Error
	})
}

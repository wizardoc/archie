package models

import (
	"archie/connection/postgres_conn"
	"archie/utils/db_utils"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	CANNOT_FIND_PERMISSION_ERROR_MESSAGE = "Cannot find the permission. "
)

type OrganizationPermission struct {
	PermissionID   string `gorm:"primary_key"`
	UserID         string `gorm:"primary_key"`
	OrganizationID string `gorm:"primary_key"`
}

func (op *OrganizationPermission) TableName() string {
	return "organization_permissions"
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
			return errors.New(CANNOT_FIND_PERMISSION_ERROR_MESSAGE)
		}

		op.PermissionID = findPermission.ID

		return db.Create(op).Error
	})
}

func (op *OrganizationPermission) NewMulti(permissions []int) error {
	return postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		var findPermissions []Permission
		p := Permission{}

		if err := p.FindMulti(&findPermissions, permissions); err != nil {
			return err
		}

		if len(findPermissions) == 0 {
			return errors.New(CANNOT_FIND_PERMISSION_ERROR_MESSAGE)
		}

		var records []OrganizationPermission

		for _, p := range findPermissions {
			records = append(records, OrganizationPermission{PermissionID: p.ID, UserID: op.UserID, OrganizationID: op.OrganizationID})
		}

		return db_utils.BatchInsert(op.TableName(), []string{"permission_id", "user_id", "organization_id"}, records)
	})
}

// 覆盖之前的权限
// 比较脏的做法，先全部删除，然后再批量写入
func (op *OrganizationPermission) CoverPermission(permissions []int) error {
	return postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		if err := db.Where("user_id = ? AND organization_id = ?", op.UserID, op.OrganizationID).Delete(OrganizationPermission{}).Error; err != nil {
			return err
		}

		return op.NewMulti(permissions)
	})
}

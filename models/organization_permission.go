package models

import (
	"archie/connection/postgres_conn"
	"archie/utils"
	"archie/utils/db_utils"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	CANNOT_FIND_PERMISSION_ERROR_MESSAGE = "Cannot find the permission. "
)

type OrganizationPermission struct {
	PermissionID   string `gorm:"type:uuid;primary_key"`
	UserID         string `gorm:"type:uuid;primary_key"`
	OrganizationID string `gorm:"type:uuid;primary_key"`
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
		if err := op.specifyPermission(db).Delete(OrganizationPermission{}).Error; err != nil {
			return err
		}

		return op.NewMulti(permissions)
	})
}

func (op *OrganizationPermission) All(results *[]Permission) error {
	return postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		p := Permission{}

		return db.Table(p.TableName()).
			Joins(fmt.Sprintf("INNER JOIN %s ON permissions.id = CAST(organization_permissions.permission_id AS UUID)", op.TableName())).
			Where("organization_id = ? AND user_id = ?", op.OrganizationID, op.UserID).
			Find(results).
			Error
	})
}

func (op *OrganizationPermission) AllAsValue(results *[]int) error {
	var permissions []Permission

	err := op.All(&permissions)

	utils.ArrayMap(permissions, func(item interface{}) interface{} {
		return item.(Permission).Value
	}, results)

	return err
}

// 指定用户是否有此权限
func (op *OrganizationPermission) Has(limitPermissionValue []int) (bool, error) {
	var permissionValues []int
	hasResult := true
	err := op.AllAsValue(&permissionValues)

	for _, lp := range limitPermissionValue {
		hasResult = utils.ArrayIncludes(permissionValues, lp) && hasResult
	}

	return hasResult, err
}

func (op *OrganizationPermission) specifyPermission(db *gorm.DB) *gorm.DB {
	return db.Where("user_id = ? AND organization_id = ?", op.UserID, op.OrganizationID)
}

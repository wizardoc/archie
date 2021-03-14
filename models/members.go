package models

import (
	"archie/connection/postgres_conn"
	"time"
)

type Member struct {
	OrganizationID string `gorm:"primaryKey;type:uuid" json:"organizationId"`
	UserID         string `gorm:"primaryKey;type:uuid" json:"userId"`
	Role           int    `gorm:"type:int" json:"role"`
	JoinTime       string `gorm:"type:varchar(200)" json:"createTime"`
}

type UserWithRole struct {
	User
	Role int32 `json:"role"`
}

func (m *Member) FindUserWithRoleByOrgID(userWithRoles *[]UserWithRole) error {
	return postgres_conn.DB.Instance().
		Raw("SELECT * FROM members JOIN users ON users.id = user_id").
		Where("organization_id = ?", m.OrganizationID).
		Scan(userWithRoles).Error
}

func (m *Member) Create() error {
	m.JoinTime = time.Now().String()

	return postgres_conn.DB.Instance().Create(m).Find(m).Error
}

package models

type DocumentPermission struct {
	PermissionID string `gorm:"primary_key"`
	UserID       string `gorm:"primary_key"`
	DocumentID   string `gorm:"primary_key"`
}

package models

type DocumentContributor struct {
	UserID     string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	DocumentID string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UpdateTime int32  `gorm:"type:bigint"`
}

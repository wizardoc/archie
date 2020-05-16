package models

type Category struct {
	ID             string     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"id"`
	Name           string     `gorm:"type:varchar(50)"json:"name"`
	Description    string     `gorm:"type:varchar(200)";json:"description"`
	Cover          string     `gorm:"type:varchar(200)"json:"cover"`
	UserID         string     `gorm:"type:uuid;"` // 分类创建者
	CreateTime     int64      `gorm:"type:bigint"json:"createTime"`
	LastModifyTime int64      `gorm:"type:bigint"json:"lastModifyTime"`
	OrganizationID string     `gorm:"type:uuid;"` // 隶属的组织
	Documents      []Document `gorm:"foreign_key:CategoryId"`
}

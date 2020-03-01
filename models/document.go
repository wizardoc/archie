package models

type Document struct {
	ID             string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Path           string `gorm:"type:varchar(200);unique"json:"path"`
	CategoryId     string // 隶属的分类
	CreateUser     string // 文档创建者
	CreateTime     int64  `gorm:"type:bigint"json:"createTime"`
	LastModifyTime int64  `gorm:"type:bigint"json:"lastModifyTime"`
	Cover          string `gorm:"type:varchar(200)"json:"cover"`
	Up             int    `gorm:"type:int"json:"up"`
	Down           int    `gorm:"type:int"json:"down"`
}

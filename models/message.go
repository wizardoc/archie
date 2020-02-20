package models

type Message struct {
	ID          string   `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"-"`
	Owner       string   `gorm:"type:uuid";json:"-"`
	Type        int      `gorm:"type:int";json:"-"`
	From        string   `gorm:"type:uuid";json:"-"`
	To          []string `gorm:"type:uuid";json:"-"`
	SendTime    int64    `gorm:"type:bigint"`
	MessageType int      `gorm:"type:int";json:"type"`
	IsRead      bool     `gorm:"bool;default:FALSE";json:"isRead"`
	Main        string   `gorm:"type:varchar(1000)";json:"-"` // contain body and title (marshal channel_message)
}

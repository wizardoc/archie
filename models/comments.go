package models

import (
	"archie/connection/postgres_conn"
	"gorm.io/gorm"
	"time"
)

const (
	UP = iota
	DOWN
	NONE
)

type Comment struct {
	ID            string          `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	DocumentID    string          `json:"documentID" gorm:"type:uuid"`
	UserID        string          `json:"-" gorm:"type:uuid"`
	Content       string          `json:"content" gorm:"type:varchar(1000)"`
	Reply         string          `json:"reply" gorm:"type:char(36)"` // 被回复的用户 ID
	CreateTime    string          `json:"createTime" gorm:"type:varchar(200)"`
	User          User            `json:"user"`
	CommentStatus []CommentStatus `json:"-"`
	Up            int             `json:"up" gorm:"-"`
	Down          int             `json:"down" gorm:"-"`
	Status        int             `json:"status" gorm:"-"`
}

func (comment *Comment) New() error {
	comment.CreateTime = time.Now().String()
	comment.Status = NONE

	return postgres_conn.DB.Instance().Create(comment).Preload("User").Find(comment).Error
}

func (comment *Comment) FindAll(page int, pageSize int, comments *[]Comment) error {
	return postgres_conn.DB.Transaction(func(db *gorm.DB) error {
		return db.Offset(page).
			Limit(pageSize).
			Where("document_id = ?", comment.DocumentID).
			Preload("User").
			Preload("CommentStatus").
			Find(comments).
			Error
	})
}

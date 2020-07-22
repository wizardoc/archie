package models

import (
	"archie/connection/postgres_conn"
	"gorm.io/gorm"
)

type CommentStatus struct {
	UserID    string `gorm:"type:uuid;primary_key;"`
	CommentID string `gorm:"type:uuid;primary_key;"`
	IsUp      bool   `gorm:"type:bool;"`
}

func (cs *CommentStatus) New() error {
	return postgres_conn.DB.Transaction(func(db *gorm.DB) error {
		var records []CommentStatus
		searchCondition := func() *gorm.DB {
			return db.Model(CommentStatus{}).Where("user_id = ? AND comment_id = ?", cs.UserID, cs.CommentID)
		}

		if err := searchCondition().Find(&records).Error; err != nil {
			return err
		}

		if len(records) == 0 {
			if err := db.Create(cs).Error; err != nil {
				return err
			}

			return nil
		}

		return searchCondition().Update("is_up", cs.IsUp).Error

		//fmt.Printf("%+v", cs)

		//return db.Clauses(clause.OnConflict{
		//	Columns:   []clause.Column{{Name: "user_id"}, {Name: "comment_id"}},
		//	DoUpdates: clause.AssignmentColumns([]string{"is_up"}),
		//}).Create(cs).Error

		//return db.Exec(`
		//	INSERT INTO comment_statuses (user_id, comment_id, is_up)
		//	VALUES (?, ?, ?)
		//	ON CONFLICT (user_id, comment_id) DO
		//	UPDATE SET is_up = ?
		//`, cs.UserID, cs.CommentID, cs.IsUp, cs.IsUp).Error
	})
}

func (cs *CommentStatus) Delete() error {
	return postgres_conn.DB.Instance().Delete(cs).Error
}

package models

import (
	"archie/connection"
	"github.com/jinzhu/gorm"
)

type UserTodo struct {
	UserID      string `gorm:"type:uuid"`
	Name        string `gorm:"type:varchar(15);primary_key"`
	Description string `gorm:"type:varchar(30)"`
	Route       string `gorm:"type:varchar(20)"`
}

func (todo *UserTodo) AddUserTodoItem() error {
	return connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Create(todo).Error
	})
}

func (todo *UserTodo) RemoveUserTodoItem() error {
	return connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Delete(todo).Error
	})
}

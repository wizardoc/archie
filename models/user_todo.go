package models

import (
	"archie/connection"
	"github.com/jinzhu/gorm"
)

type UserTodo struct {
	UserID      string `gorm:"type:uuid;primary_key"json:"-"`
	Name        string `gorm:"type:varchar(15);primary_key"json:"name"`
	Description string `gorm:"type:varchar(30)"json:"description"`
	Route       string `gorm:"type:varchar(20)"json:"route"`
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

func (todo *UserTodo) GetAllTodoItemsByID() (todoItems []UserTodo, err error) {
	todoItems = []UserTodo{}
	err = connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Where("user_id = ?", todo.UserID).Find(&todoItems).Error
	})

	return
}

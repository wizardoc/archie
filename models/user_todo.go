package models

import (
	"archie/connection"
	"archie/utils"
)

type UserTodo struct {
	UserID      string `gorm:"type:uuid"`
	Name        string `gorm:"type:varchar(15);primary_key"`
	Description string `gorm:"type:varchar(30)"`
	Route       string `gorm:"type:varchar(20)"`
}

func (todo *UserTodo) AddUserTodoItem() {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	db.Create(todo)
}

func (todo *UserTodo) RemoveUserTodoItem() {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	db.Delete(todo)
}

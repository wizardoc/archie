package main

import (
	"archie/connection"
	"archie/models"
	"archie/routes"
	"archie/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createDataTable(db *gorm.DB, model interface{}) {
	if !db.HasTable(model) {
		db.CreateTable(model)
	}
}

func InitTable() {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	// init table
	createDataTable(db, models.UserOrganization{})
	createDataTable(db, models.User{})
	createDataTable(db, models.Organization{})
	createDataTable(db, models.UserTodo{})
	createDataTable(db, models.RolePermission{})
	createDataTable(db, models.Role{})
	createDataTable(db, models.UserRole{})
	createDataTable(db, models.Permission{})
	createDataTable(db, models.Category{})
	createDataTable(db, models.DocumentContribute{})
}

func main() {
	InitTable()
	routes.Serve()
}

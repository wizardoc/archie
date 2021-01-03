package main

import (
	"archie/connection/postgres_conn"
	"archie/models"
	"archie/models/focus_models"
	"archie/routes"
	"archie/services"
	"archie/utils/db_utils/db_migrate_utils"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func initTable() {
	db_migrate_utils.InitTable(func(db *gorm.DB) error {
		return db.AutoMigrate(
			models.UserOrganization{},
			models.User{},
			models.Message{},
			models.Organization{},
			models.RolePermission{},
			models.Role{},
			models.Permission{},
			models.Category{},
			models.DocumentContributor{},
			models.Document{},
			models.DocumentPermission{},
			models.OrganizationPermission{},
			models.Comment{},
			models.CommentStatus{},
			focus_models.FocusOrganization{},
			focus_models.FocusUser{},
		)
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//// init database
	postgres_conn.DB.InitDB()
	go services.Receiver.Run()

	initTable()

	routes.Serve()
}

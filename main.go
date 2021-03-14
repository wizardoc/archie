package main

import (
	"archie/connection/postgres_conn"
	"archie/models"
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
		if err := db.AutoMigrate(
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
			models.Member{},
		); err != nil {
			return err
		}

		return db.SetupJoinTable(&models.Organization{}, "Members", &models.Member{})
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// init database
	postgres_conn.DB.InitDB()
	go services.Receiver.Run()

	initTable()

	routes.Serve()
}

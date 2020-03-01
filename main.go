package main

import (
	"archie/connection/postgres_conn"
	"archie/models"
	"archie/routes"
	"archie/services"
	"archie/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

func createDataTable(db *gorm.DB, model ...interface{}) {
	if !db.HasTable(model) {
		db.AutoMigrate(model...)
	}
}

func InitTable() {
	db, err := postgres_conn.GetDB()

	utils.Check(err)
	defer db.Close()

	// init table
	createDataTable(db, models.UserOrganization{})
	createDataTable(db, models.User{}, models.Message{})
	createDataTable(db, models.Organization{})
	createDataTable(db, models.UserTodo{})
	createDataTable(db, models.RolePermission{})
	createDataTable(db, models.Role{})
	createDataTable(db, models.UserRole{})
	createDataTable(db, models.Permission{})
	createDataTable(db, models.Category{})
	createDataTable(db, models.DocumentContribute{})
}

var a = map[string]int{
	"aa": 1,
	"bb": 2,
}

func main() {
	go services.Receiver.Run()

	t := time.NewTicker(time.Duration(time.Second * 1))

	go func() {
		for {
			<-t.C
			_, err := services.NewChannelMessage("471bd6a4-15bf-4097-8b75-39987e06c1d7", "471bd6a4-15bf-4097-8b75-39987e06c1d7", []string{"471bd6a4-15bf-4097-8b75-39987e06c1d7"}, services.BROADCAST, services.NOTIFY, "ssss", "sddasdasdas")

			if err != nil {
				log.Fatal(err)
			}
			//services.NewPublisher(data).Publish()
		}
	}()

	InitTable()
	routes.Serve()
}

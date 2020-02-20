package main

import (
	"archie/connection"
	"archie/models"
	"archie/routes"
	"archie/services"
	"archie/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
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
}

var a = map[string]int{
	"aa": 1,
	"bb": 2,
}

func main() {
	go services.Receiver.Run()

	data, err := services.NewChannelMessage("aaaa", "aaaa", []string{"aaaa"}, 0, services.BROADCAST, "ssss", "sddasdasdas")

	if err != nil {
		log.Fatal(err)
	}

	time.AfterFunc(4*time.Second, func() {
		services.NewPublisher(data).Publish()
	})

	InitTable()
	routes.Serve()
}

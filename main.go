package main

import (
	"archie/models"
	"archie/routes"
	"archie/services"
	"archie/utils/db_utils/db_migrate_utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
		).Error
	})
}

func main() {
	go services.Receiver.Run()

	//t := time.NewTicker(time.Duration(time.Second * 1))

	//go func() {
	//	for {
	//		<-t.C
	//		data, err := services.NewChannelMessage("471bd6a4-15bf-4097-8b75-39987e06c1d7", "471bd6a4-15bf-4097-8b75-39987e06c1d7", []string{"471bd6a4-15bf-4097-8b75-39987e06c1d7"}, services.BROADCAST, services.NOTIFY, "ssss", "sddasdasdas")
	//
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		services.NewPublisher(data).Publish()
	//	}
	//}()

	initTable()
	routes.Serve()
	//
	//op := models.OrganizationPermission{OrganizationID: "8663e9b7-d6aa-462c-8fed-d714437a37a6", UserID: "645ee6c0-4551-42ce-a450-2dc2900f170f"}
	////var results []models.Permission
	//
	//if err := op.NewMulti([]int{
	//	models.DOCUMENT_VIEW,
	//	models.DOCUMENT_READ,
	//	models.ORG_INVITE,
	//}); err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(results)
}

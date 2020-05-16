package main

import (
	"archie/models"
	"archie/utils/db_utils/db_migrate_utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
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
	//go services.Receiver.Run()

	//t := time.NewTicker(time.Duration(time.Second * 1))
	//
	//go func() {
	//	for {
	//		<-t.C
	//		_, err := services.NewChannelMessage("471bd6a4-15bf-4097-8b75-39987e06c1d7", "471bd6a4-15bf-4097-8b75-39987e06c1d7", []string{"471bd6a4-15bf-4097-8b75-39987e06c1d7"}, services.BROADCAST, services.NOTIFY, "ssss", "sddasdasdas")
	//
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		//services.NewPublisher(data).Publish()
	//	}
	//}()

	initTable()
	//routes.Serve()

	p := models.OrganizationPermission{UserID: "6f937ed0-a56a-4b62-bbc3-a477542fe3ce", OrganizationID: "www"}

	if err := p.New(models.ORG_DELETE); err != nil {
		log.Fatal(err)
	}
}

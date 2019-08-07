package main

import (
	"archie/connection"
	"archie/models"
	"archie/routes"
	"archie/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createDataTable(db *gorm.DB, model interface{}) (hasTable bool) {
	hasTable = db.HasTable(model)

	if !hasTable {
		db.CreateTable(model)
	}

	return
}

func InitTable() {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	// init table
	createDataTable(db, models.UserOrganization{})
	createDataTable(db, models.User{})
	createDataTable(db, models.Organization{})
}

func main() {
	InitTable()

	routes.Serve()

	//claim := utils.Claims{
	//	"younccat",
	//	uuid.NewV4().String(),
	//	time.Now().Unix(),
	//	time.Now().Add(time.Hour).Unix(),
	//	"122312312312",
	//}
	//
	//jwtStr := claim.SignJWT()
	//
	//utils.ParseToken(jwtStr)
}

package connection

import (
	"archie/models"
	"archie/utils"
	"github.com/jinzhu/gorm"
)

func crearteDataTable(db *gorm.DB, model interface{}) (hasTable bool) {
	hasTable = db.HasTable(model)

	if !hasTable {
		db.CreateTable(model)
	}

	return
}

func InitTable() {
	db, err := GetDB()

	utils.Check(err)
	defer db.Close()

	// init table
	crearteDataTable(db, models.User{})
}

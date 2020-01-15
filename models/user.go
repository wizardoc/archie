package models

import (
	"archie/connection"
	"archie/utils"
	"fmt"
)

type LoginInfo struct {
	Username string `gorm:"type:varchar(20);unique;" json:"username" form:"username" validate:"gt=4,lt=20,required"`
	Password string `gorm:"type:char(62)" json:"-" form:"password" validate:"required,gt=4,lt=20"`
}

type RegisterInfo struct {
	LoginInfo
	Email        string `gorm:"type:varchar(64)" json:"email" form:"email" validate:"email,required"`
	DisplayName  string `gorm:"type:varchar(12)" json:"displayName" form:"displayName" validate:"required,gt=2,lt=10"`
	RegisterTime int64  `gorm:"type:bigint"json:"registerTime"`
	IsValidEmail bool   `gorm:"type:boolean"json:"-"`
}

type User struct {
	ID            string          `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"-"`
	Avatar        string          `gorm:"type:varchar(200)"json:"avatar"`
	Organizations *[]Organization `gorm:"many2many:user_organizations;"json:"-"`
	LoginTime     int64           `gorm:"type:bigint"json:"loginTime"`
	RegisterInfo
}

func (user *User) Register() (ok bool) {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	user.RegisterTime = utils.Now()
	// make more security password
	user.Password = utils.Hash(user.Password)
	user.IsValidEmail = false

	db.Create(user)

	return true
}

func (user *User) UpdateLoginTime() {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	db.Model(&user).Where("id = ?", user.ID).Update("login_time", utils.Now())
}

func (user *User) GetUserInfoByID() User {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	userID := user.ID
	result := User{}

	db.Find(&result, "id = ?", userID)

	return result
}

func findUser(queryKey string, queryBody string) User {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()
	user := User{}

	db.Find(&user, fmt.Sprintf("%s = ?", queryKey), queryBody)

	return user
}

func FindOneByUsername(username string) User {
	return findUser("username", username)
}

func FindOneByEmail(email string) User {
	return findUser("email", email)
}

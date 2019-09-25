package models

import (
	"archie/connection"
	"archie/utils"
	"fmt"
	"time"
)

type User struct {
	ID            string          `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"-"`
	DisplayName   string          `gorm:"type:varchar(12)"json:"displayName"`
	Username      string          `gorm:"type:varchar(20);unique;"json:"username"`
	Email         string          `gorm:"type:varchar(64)"json:"email"`
	Avatar        string          `gorm:"type:varchar(200)"json:"avatar"`
	Organizations *[]Organization `gorm:"many2many:user_organizations;"json:"-"`
	RegisterTime  int64           `gorm:"type:bigint"json:"registerTime"`
	LoginTime     int64           `gorm:"type:bigint"json:"loginTime"`
	Password      string          `gorm:"type:char(62)"json:"-"`
	IsValidEmail  bool            `gorm:"type:boolean"json:"-"`
}

func (user *User) Register() (ok bool) {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	user.RegisterTime = time.Now().Unix()
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

	db.Model(&user).Where("id = ?", user.ID).Update("login_time", time.Now().Unix())
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

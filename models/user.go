package models

import (
	"archie/connection"
	"archie/utils"
	"fmt"
	"github.com/jinzhu/gorm"
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

func (user *User) Register() error {
	return connection.WithPostgreConn(func(db *gorm.DB) error {
		user.RegisterTime = utils.Now()
		// make more security password
		user.Password = utils.Hash(user.Password)
		user.IsValidEmail = false

		return db.Create(user).Error
	})
}

func (user *User) UpdateLoginTime() error {
	return connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Model(&user).Where("id = ?", user.ID).Update("login_time", utils.Now()).Error
	})
}

func (user *User) GetUserInfoByID() (result User, err error) {
	userID := user.ID
	result = User{}

	err = connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Find(&result, "id = ?", userID).Error
	})

	return
}

// 更新 user model 里有值的字段
func (user *User) UpdateAvatar() error {
	return updateSig(user, "avatar", user.Avatar)
}

func findUser(queryKey string, queryBody string) (user User, err error) {
	user = User{}
	err = connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Find(&user, fmt.Sprintf("%s = ?", queryKey), queryBody).Error
	})

	return
}

func FindOneByUsername(username string) (User, error) {
	return findUser("username", username)
}

func FindOneByEmail(email string) (User, error) {
	return findUser("email", email)
}

package models

import (
	"archie/connection"
	"archie/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	Username      string          `gorm:"type:varchar(20);unique;" json:"username" `
	Password      string          `gorm:"type:char(62)" json:"-" `
	ID            string          `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"-"`
	Avatar        string          `gorm:"type:varchar(200)"json:"avatar"`
	Organizations *[]Organization `gorm:"many2many:user_organizations;"json:"-"`
	LoginTime     int64           `gorm:"type:bigint"json:"loginTime"`
	Categories    []Category      `gorm:"foreign_key:CreateUser"`
	Email         string          `gorm:"type:varchar(64)" json:"email"`
	DisplayName   string          `gorm:"type:varchar(12)" json:"displayName"`
	RegisterTime  int64           `gorm:"type:bigint"json:"registerTime"`
	IsValidEmail  bool            `gorm:"type:boolean"json:"-"`
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

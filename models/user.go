package models

import (
	"archie/connection/postgres_conn"
	"archie/robust"
	"archie/utils"
	"fmt"
	"gorm.io/gorm"
)

type RegisterInfo struct {
}

type User struct {
	ID                 string              `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"id"`
	Username           string              `gorm:"type:varchar(20);unique;" json:"username"`
	Password           string              `gorm:"type:char(62)" json:"-"`
	Email              string              `gorm:"type:varchar(64)" json:"email"`
	DisplayName        string              `gorm:"type:varchar(12)" json:"displayName"`
	RegisterTime       int32               `gorm:"type:bigint"json:"registerTime"`
	IsValidEmail       bool                `gorm:"type:boolean"json:"isValidEmail"`
	Avatar             string              `gorm:"type:varchar(200)"json:"avatar"`
	RealName           string              `gorm:"type:varchar(10)" json:"realName"`
	Intro              string              `gorm:"type:varchar(50)" json:"intro"`        // 个人简介
	City               string              `gorm:"type:varchar(50)"json:"city"`          // 所在城市
	CompanyName        string              `gorm:"type:varchar(80)" json:"companyName"`  // 公司名称
	CompanyTitle       string              `gorm:"type:varchar(80)" json:"companyTitle"` // 职位头衔
	Github             string              `gorm:"type:varchar(100)" json:"github"`      // GitHub 地址
	Blog               string              `gorm:"varchar(100)" json:"blog"`             // 博客地址
	PayQRCode          string              `gorm:"varchar(200)" json:"payQRCode"`        // 打赏支付二维码
	Organizations      *[]Organization     `gorm:"many2many:user_organizations"json:"-"`
	LoginTime          int32               `gorm:"type:bigint"json:"loginTime"`
	Messages           []Message           `gorm:"many2many:user_messages"json:"-"`
	UserOrganizations  []*UserOrganization `json:"-"`
	FocusOrganizations []Organization      `gorm:"many2many:focus_organizations" json:"followOrganizations"`
	FocusUsers         []User              `gorm:"many2many:focus_users;association_jointable_foreignkey:focus_user_id" json:"followUsers"`
	RegisterInfo
}

func (user *User) SearchName(name string, page int, pageSize int, users *[]User) error {
	return postgres_conn.DB.Instance().Model(user).Where("username LIKE ?", fmt.Sprintf("%s%%", name)).Offset((page - 1) * pageSize).Limit(pageSize).Find(users).Error
}

func (user *User) FindAllMessages(page int, pageSize int) error {
	return postgres_conn.DB.Instance().Model(user).Where("id = ?", user.ID).Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pageSize).Limit(pageSize).Order("send_time desc")
	}).Find(user).Error
}

func (user *User) Follow() error {
	return postgres_conn.DB.Instance().Create(user).Error
}

func (user *User) Register() error {
	user.RegisterTime = utils.Now()
	// make password more security
	user.Password = utils.Hash(user.Password)
	user.IsValidEmail = false

	return postgres_conn.DB.Instance().Create(user).Find(user).Error
}

func (user *User) Find(queryKey string, queryBody string) error {
	return postgres_conn.DB.Instance().Find(&user, fmt.Sprintf("%s = ?", queryKey), queryBody).Error
}

func (user *User) UpdateLoginTime() error {
	return postgres_conn.DB.Instance().Model(user).Where("id = ?", user.ID).Update("login_time", utils.Now()).Error
}

func (user *User) GetUserInfoByID() error {
	return postgres_conn.DB.Instance().Model(&User{}).Preload("FocusUsers").Preload("FocusOrganizations").Find(&user, "id = ?", user.ID).Error
}

// 更新 user model 里有值的字段
func (user *User) UpdateAvatar() error {
	return postgres_conn.DB.Instance().Model(user).Where("id = ?", user.ID).Update("avatar", user.Avatar).Error
}

func (user *User) FindByUsername(username string) error {
	findUser, err := findUser("username", username)
	*user = findUser
	return err
}

func (user *User) UpdateUserInfo() error {
	return postgres_conn.DB.Instance().Model(user).Where("id = ?", user.ID).Updates(*user).Find(user).Error
}

func findUser(queryKey string, queryBody string) (user User, err error) {
	user = User{}
	err = postgres_conn.DB.Instance().Find(&user, fmt.Sprintf("%s = ?", queryKey), queryBody).Error

	if user.ID == "" {
		err = robust.USER_DOSE_NOT_EXIST
	}

	return
}

func FindOneByUsername(username string) (User, error) {
	return findUser("username", username)
}

func FindOneByEmail(email string) (User, error) {
	return findUser("email", email)
}

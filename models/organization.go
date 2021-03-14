package models

import (
	"archie/connection/postgres_conn"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Organization struct {
	ID          string  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"id"`
	Name        string  `gorm:"type:varchar(20);unique;"json:"name"`
	Description string  `gorm:"type:varchar(50)"json:"description"`
	Owner       string  `gorm:"type:uuid;default:uuid_generate_v4()"json:"owner"`
	Cover       string  `gorm:"type:varchar(200)"json:"cover"`
	CreateTime  string  `gorm:"type:varchar(200)"json:"createTime"`
	IsPublic    bool    `gorm:"type:bool;default:FALSE"json:"isPublic"`
	FollowUsers []*User `gorm:"many2many:organization_follow_users" json:"followUsers"`
	Members     []*User `gorm:"many2many:members" json:"members"`
}

type OrganizationName struct {
	OrganizeName string
}

func (organization *Organization) Find(key string, val interface{}) error {
	return postgres_conn.DB.Instance().Preload("FollowUsers").Find(organization, fmt.Sprintf("%s = ?", key), val).Error
}

func (organization *Organization) FindOneByID() error {
	return organization.Find("id", organization.ID)
}

func (organization *Organization) FindOneByOrganizeName() error {
	return organization.Find("name", organization.Name)
}

func (organization *Organization) Update(id string) func(key string, val interface{}) error {
	return func(key string, val interface{}) error {
		return postgres_conn.DB.Instance().Model(&organization).Where("id = ?", id).Update(key, val).Error
	}
}

func (organization *Organization) BatchUpdates(source map[string]interface{}) error {
	return postgres_conn.DB.Instance().Model(Organization{}).Omit("id").Where("id = ?", organization.ID).Updates(source).Error
}

func (organization *Organization) Create() error {
	return postgres_conn.DB.Instance().Create(organization).Find(organization).Error
}

func (organization *Organization) New(username string) error {
	return postgres_conn.DB.Transaction(func(db *gorm.DB) error {
		organization.CreateTime = time.Now().String()
		user, err := FindOneByUsername(username)

		if err != nil {
			return err
		}

		organization.Owner = user.ID

		if err := postgres_conn.DB.Instance().Create(organization).Error; err != nil {
			return err
		}

		userOrganization := UserOrganization{
			UserID:         user.ID,
			OrganizationID: organization.ID,
		}

		dp := OrganizationPermission{UserID: user.ID, OrganizationID: organization.ID}
		if err := dp.NewMulti(AllPermissions()); err != nil {
			return err
		}

		return userOrganization.New(true)
	})
}

func (organization *Organization) GetAllNames() (names []OrganizationName, err error) {
	names = []OrganizationName{}
	err = postgres_conn.DB.Instance().Select("organize_name").Find(organization).Scan(&names).Error

	return
}

func (organization *Organization) DeleteAssociation(associationName string, identity interface{}) error {
	return postgres_conn.DB.Instance().Model(organization).Association(associationName).Delete(identity)
}

func (organization *Organization) AppendAssociation(associationName string, identity interface{}) error {
	return postgres_conn.DB.Instance().Model(organization).Association(associationName).Append(identity)
}

func (organization *Organization) AllByUserId(id string) (organizations []Organization, err error) {
	organizations = []Organization{}
	err = postgres_conn.DB.Instance().Find(&organizations).Where("id=?", id).Error

	return
}

func (organization *Organization) RemoveOrganization() error {
	return postgres_conn.DB.Instance().Delete(organization).Error
}

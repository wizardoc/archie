package models

import (
	"archie/connection/postgres_conn"
	"archie/utils"
	"gorm.io/gorm"
)

type Organization struct {
	ID           string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"id"`
	OrganizeName string `gorm:"type:varchar(20);unique;"json:"organizeName"`
	Description  string `gorm:"type:varchar(50)"json:"description"`
	HasValid     bool   `gorm:"type:bool;default:TRUE"json:"hasValid"`
	Owner        string `gorm:"type:uuid;default:uuid_generate_v4()"json:"-"` // related userID
	CreateTime   int32  `gorm:"type:bigint"json:"createTime"`
	IsPublic     bool   `gorm:"type:bool;default:TRUE"json:"isPublic"`
	FocusUsers   []User `gorm:"many2many:focus_organizations" json:"followUsers"`
}

type OrganizationName struct {
	OrganizeName string
}

func (organization *Organization) FindOneByID(id string) error {
	return postgres_conn.DB.Instance().Find(organization, "id = ?", id).Error
}

func (organization *Organization) FindOneByOrganizeName() error {
	return postgres_conn.DB.Instance().Find(organization, "organize_name=?", organization.OrganizeName).Error
}

func (organization *Organization) Update(id string) func(key string, val interface{}) error {
	return func(key string, val interface{}) error {
		return postgres_conn.DB.Instance().Model(&organization).Where("id = ?", id).Update(key, val).Error
	}
}

func (organization *Organization) BatchUpdates(source map[string]interface{}) error {
	return postgres_conn.DB.Instance().Model(Organization{}).Omit("id").Where("id = ?", organization.ID).Updates(source).Error
}

func (organization *Organization) New(username string) error {
	return postgres_conn.DB.Transaction(func(db *gorm.DB) error {
		organization.HasValid = true
		organization.CreateTime = utils.Now()
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

		// 赋予 owner 权限
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

func (organization *Organization) AllByUserId(id string) (organizations []Organization, err error) {
	organizations = []Organization{}
	err = postgres_conn.DB.Instance().Find(&organizations).Where("id=?", id).Error

	return
}

func (organization *Organization) RemoveOrganization() error {
	return postgres_conn.DB.Instance().Delete(organization).Error
}

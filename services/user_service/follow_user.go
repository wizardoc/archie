package user_service

import (
	"archie/connection/postgres_conn"
	"archie/models"
	"gorm.io/gorm"
)

func FollowUser(userID string, followUserID string) error {
	user := models.User{
		ID: userID,
	}
	followUser := models.User{
		ID: followUserID,
	}

	return postgres_conn.DB.Transaction(func(db *gorm.DB) error {
		err := user.AppendAssociation("Followings", &followUser)
		if err != nil {
			return err
		}

		return followUser.AppendAssociation("Followers", &user)
	})
}

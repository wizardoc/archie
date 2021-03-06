package user_service

import (
	"archie/connection/postgres_conn"
	"archie/models"
	"gorm.io/gorm"
)

func UnfollowUser(userID string, unfollowUserID string) error {
	user := models.User{ID: userID}
	unfollowUser := models.User{ID: unfollowUserID}

	return postgres_conn.DB.Transaction(func(db *gorm.DB) error {
		err := user.DeleteAssociation("Followings", unfollowUser)
		if err != nil {
			return err
		}

		return unfollowUser.DeleteAssociation("Followers", user)
	})
}

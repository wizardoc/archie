package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, err robust.ArchieError) {
	helper.Send(context, gin.H{"isValid": false}, err)
}

/** 验证用户 baseInfo */
func ValidBaseInfo(context *gin.Context) {
	email := context.PostForm("email")
	username := context.PostForm("username")

	user := models.FindOneByUsername(username)

	if user.ID != "" {
		sendError(context, robust.REGISTER_EXIST_USER)

		return
	}

	if user.Email == email {
		sendError(context, robust.EMAIL_DOSE_EXIST)

		return
	} else if models.FindOneByEmail(email).Email == email {
		sendError(context, robust.EMAIL_DOSE_EXIST)

		return
	}

	helper.Send(context, gin.H{"isValid": true}, nil)
}

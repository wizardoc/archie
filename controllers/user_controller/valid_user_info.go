package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

/** 验证用户 baseInfo */
func ValidBaseInfo(context *gin.Context) {
	email := context.PostForm("email")
	username := context.PostForm("username")

	user := models.FindOneByUsername(username)
	errRes := helper.Res{Status: http.StatusBadRequest}

	if user.ID != "" {
		errRes.Err = robust.REGISTER_EXIST_USER
		errRes.Send(context)
		return
	}

	if user.Email == email {
		errRes.Err = robust.EMAIL_DOSE_EXIST
		errRes.Send(context)
		return
	}

	if models.FindOneByEmail(email).Email == email {
		errRes.Err = robust.EMAIL_DOSE_EXIST
		errRes.Send(context)
		return
	}

	helper.Res{Data: gin.H{"isValid": true}}.Send(context)
}

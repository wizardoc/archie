package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseInfo struct {
	Email    string `form:"email"`
	Username string `form:"username"`
}

/** 验证用户 baseInfo */
func ValidBaseInfo(context *gin.Context) {
	errRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	var baseInfo BaseInfo
	if err := context.Bind(&baseInfo); err != nil {
		errRes.Err = robust.INVALID_PARAMS
		errRes.Send(context)
		return
	}

	email := context.PostForm("email")
	username := context.PostForm("username")

	user := models.FindOneByUsername(username)

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

	res.Data = gin.H{"isValid": true}
	res.Send(context)
}

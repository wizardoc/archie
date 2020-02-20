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

// validate base info of user when register
func ValidBaseInfo(context *gin.Context) {
	errRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	var baseInfo BaseInfo
	if err := context.Bind(&baseInfo); err != nil {
		errRes.Err = robust.INVALID_PARAMS
		errRes.Send(context)
		return
	}

	// user does exist
	if _, err := models.FindOneByUsername(baseInfo.Username); err == nil {
		errRes.Err = robust.REGISTER_EXIST_USER
		errRes.Send(context)
		return
	}

	// the email of user does exist
	if _, err := models.FindOneByEmail(baseInfo.Email); err == nil {
		errRes.Err = robust.EMAIL_DOSE_EXIST
		errRes.Send(context)
		return
	}

	res.Data = gin.H{
		"isValid": true,
	}
	res.Send(context)
}

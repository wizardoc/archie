package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterInfo struct {
	Username    string `validate:"gt=4,lt=20"`
	Password    string `validate:"gt=4,lt=20"`
	Email       string `validate:"email"`
	DisplayName string `validate:"gt=2,lt=10"`
	//Avatar string
}

/** 用户注册 */
func Register(context *gin.Context) {
	//organizationName := context.PostForm("organizationName")
	//organizationDescription := context.PostForm("organizationDescription")
	//
	//utils.Green(organizationName)
	//utils.Green(organizationDescription)
	var user = RegisterInfo{}
	if err := context.Bind(&user); err != nil {
		helper.Res{Status: http.StatusBadRequest, Err: robust.INVALID_PARAMS}.Send(context)
		return
	}

	validator := robust.Validation{Target: user}
	if err := validator.Valid(); err != nil {
		helper.Res{Status: http.StatusBadRequest, Err: err}.Send(context)
		return
	}

	findUser := models.FindOneByUsername(user.Username)

	if findUser.ID != "" {
		helper.Res{Err: robust.REGISTER_EXIST_USER}.Send(context)
		return
	}

	var userValue interface{} = user
	userModel := userValue.(models.User)

	ok := userModel.Register()

	if !ok {
		helper.Res{Err: robust.CREATE_DATA_FAILURE}.Send(context)
		return
	}

	helper.Res{Data: user}.Send(context)
}

package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

/** 用户注册 */
func Register(context *gin.Context) {
	//organizationName := context.PostForm("organizationName")
	//organizationDescription := context.PostForm("organizationDescription")
	//
	//utils.Green(organizationName)
	//utils.Green(organizationDescription)
	res := helper.Res{}
	errRes := helper.Res{Status: http.StatusBadRequest}

	var user = models.User{}
	if err := helper.BindWithValid(context, &user); err != nil {
		errRes.Err = err
		errRes.Send(context)
		return
	}

	findUser := models.FindOneByUsername(user.Username)

	if findUser.ID != "" {
		errRes.Err = robust.REGISTER_EXIST_USER
		errRes.Send(context)
		return
	}

	ok := user.Register()

	if !ok {
		errRes.Err = robust.CREATE_DATA_FAILURE
		errRes.Send(context)
		return
	}

	res.Data = user
	res.Send(context)
}

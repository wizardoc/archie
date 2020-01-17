package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserAvatar struct {
	Avatar string `form:"avatar" validate:"required"`
}

// 更新用户头像
func UpdateAvatar(context *gin.Context) {
	userAvatar := UserAvatar{}
	err := helper.BindWithValid(context, &userAvatar)

	errRes := helper.Res{Status: http.StatusBadRequest}
	serverErrRes := helper.Res{Status: http.StatusInternalServerError}
	res := helper.Res{}

	if err != nil {
		errRes.Err = err
		errRes.Send(context)
		return
	}

	user := models.User{
		Avatar: userAvatar.Avatar,
	}

	if err := user.UpdateAvatar(); err != nil {
		serverErrRes.Err = robust.DB_UPDATE_FAILURE
		serverErrRes.Send(context)
		return
	}

	res.Send(context)
}

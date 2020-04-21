package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

type UserAvatar struct {
	Avatar string `form:"avatar" validate:"required"`
}

// 更新用户头像
func UpdateAvatar(ctx *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(ctx)
	authRes := helper.GenAuthRes()
	badReqRes := helper.GenBadReqRes()
	successRes := helper.GenSuccessRes()
	serverErrRes := helper.GenServerErrRes()

	if err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	utils.Green("aaaa")

	userAvatar := UserAvatar{}
	err = helper.BindWithValid(ctx, &userAvatar)

	if err != nil {
		badReqRes.Err = err
		badReqRes.Send(ctx)
		return
	}

	user := models.User{
		Avatar: userAvatar.Avatar,
		ID:     parsedClaims.UserId,
	}

	if err := user.UpdateAvatar(); err != nil {
		serverErrRes.Err = robust.DB_UPDATE_FAILURE
		serverErrRes.Send(ctx)
		return
	}

	successRes.Send(ctx)
}

package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

/** 获取用户信息 */
func GetUserInfo(ctx *gin.Context) {
	unAuthRes := helper.Res{Status: http.StatusUnauthorized}
	serverErrRes := helper.Res{Status: http.StatusInternalServerError}
	res := helper.Res{}

	claims, err := middlewares.GetClaims(ctx)

	// claims 不存在
	if err != nil {
		unAuthRes.Err = err
		unAuthRes.Send(ctx)
		return
	}

	user := models.User{
		ID: claims.UserId,
	}
	userInfo, err := user.GetUserInfoByID()

	// 找不到用户
	if err != nil {
		serverErrRes.Err = robust.CANNOT_FIND_USER
		serverErrRes.Send(ctx)
		return
	}

	res.Data = gin.H{
		"userInfo": userInfo,
	}
	res.Send(ctx)
}

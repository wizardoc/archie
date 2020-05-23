package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

/** 获取用户信息 */
func GetUserInfo(ctx *gin.Context) {
	res := helper.Res{}

	claims, err := middlewares.GetClaims(ctx)

	// claims 不存在
	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	user := models.User{
		ID: claims.UserId,
	}

	// 找不到用户
	if err := user.GetUserInfoByID(); err != nil {
		res.Status(http.StatusNotFound).Error(ctx, err)
		return
	}

	res.Data = gin.H{
		"userInfo": user,
	}
	res.Send(ctx, nil)
}

package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

/** 用户登录 */
func Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	// check user is exist
	user := models.FindOneByUsername(username)

	if helper.IsEmpty(user) || user.ID == "" {
		helper.Send(context, nil, robust.LOGIN_USER_DOES_NOT_EXIST)

		return
	}

	if utils.Hash(password) != user.Password {
		helper.Send(context, nil, robust.LOGIN_PASSWORD_NOT_VALID)

		return
	}

	// 验证是否在黑名单
	if middlewares.IsExistInBlackSet(user.ID) {
		helper.Send(context, nil, robust.JWT_NOT_ALLOWED)

		return
	}

	user.UpdateLoginTime()

	claims := utils.Claims{
		UserId: user.ID,
	}

	helper.Send(context, gin.H{
		"jwt":      claims.SignJWT(),
		"userInfo": user,
	}, nil)
}

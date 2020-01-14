package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

/** 用户登录 */
func Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	// 检查用户是否存在
	user := models.FindOneByUsername(username)
	errRes := helper.Res{Status: http.StatusBadRequest}

	if helper.IsEmpty(user) || user.ID == "" {
		errRes.Err = robust.LOGIN_USER_DOES_NOT_EXIST
		errRes.Send(context)
		return
	}

	if utils.Hash(password) != user.Password {
		errRes.Err = robust.LOGIN_PASSWORD_NOT_VALID
		errRes.Send(context)
		return
	}

	// 验证是否在黑名单
	if middlewares.IsExistInBlackSet(user.ID) {
		errRes.Err = robust.JWT_NOT_ALLOWED
		errRes.Send(context)
		return
	}

	user.UpdateLoginTime()

	claims := utils.Claims{
		UserId: user.ID,
	}

	helper.Res{Data: gin.H{
		"jwt":      claims.SignJWT(),
		"userInfo": user,
	}}.Send(context)
}

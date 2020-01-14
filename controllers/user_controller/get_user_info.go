package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

/** 获取用户信息 */
func GetUserInfo(context *gin.Context) {
	claims, isExist := context.Get("claims")
	authRes := helper.Res{Status: http.StatusBadRequest}
	unAuthRes := helper.Res{Status: http.StatusUnauthorized}
	res := helper.Res{}

	// claims 不存在
	if !isExist {
		unAuthRes.Err = robust.JWT_PARSE_ERROR
		unAuthRes.Send(context)
		return
	}

	parsedClaims, ok := claims.(utils.Claims)

	// 解析 claims 错误
	if !ok {
		authRes.Err = robust.JWT_PARSE_ERROR
		authRes.Send(context)
		return
	}

	user := models.User{
		ID: parsedClaims.UserId,
	}
	userInfo := user.GetUserInfoByID()

	res.Data = gin.H{
		"userInfo": userInfo,
	}
	res.Send(context)
}

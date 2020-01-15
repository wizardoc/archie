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
	serverErrRes := helper.Res{Status: http.StatusInternalServerError}
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
	userInfo, err := user.GetUserInfoByID()

	// 找不到用户
	if err != nil {
		serverErrRes.Err = robust.CANNOT_FIND_USER
		serverErrRes.Send(context)
		return
	}

	res.Data = gin.H{
		"userInfo": userInfo,
	}
	res.Send(context)
}

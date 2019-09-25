package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)

/** 获取用户信息 */
func GetUserInfo(context *gin.Context) {
	claims, isExist := context.Get("claims")

	fmt.Println(claims)

	if !isExist {
		helper.Send(context, nil, robust.JWT_PARSE_ERROR)

		return
	}

	parsedClaims, ok := claims.(utils.Claims)

	if !ok {
		helper.Send(context, nil, robust.JWT_PARSE_ERROR)

		return
	}

	user := models.User{
		ID: parsedClaims.UserId,
	}

	helper.Send(context, gin.H{
		"userInfo": user.GetUserInfoByID(),
	}, nil)
}

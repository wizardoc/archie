package controllers

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context) {
	organizationName := context.PostForm("organizationName")
	organizationDescription := context.PostForm("organizationDescription")

	utils.Green(organizationName)
	utils.Green(organizationDescription)

	user := models.User{
		Password:    context.PostForm("password"),
		Avatar:      "",
		Username:    context.PostForm("username"),
		DisplayName: context.PostForm("displayName"),
		Email:       context.PostForm("email"),
	}

	findUser := models.FindOneByUsername(user.Username)

	if findUser.ID != "" {
		context.JSON(http.StatusOK, gin.H{
			"data": nil,
			"err":  robust.REGISTER_EXIST_USER,
		})

		return
	}

	ok := user.Register()

	if !ok {
		context.JSON(http.StatusOK, gin.H{
			"data": nil,
			"err":  robust.CREATE_DATA_FAILURE,
		})

		return
	}

	context.JSON(200, gin.H{
		"data": user,
		"err":  nil,
	})
}

func Login(context *gin.Context) {
	username := context.PostForm("username")

	// check user is exist
	user := models.FindOneByUsername(username)

	if utils.IsEmpty(user) || user.ID == "" {
		utils.Send(context, nil, robust.LOGIN_USER_DOES_NOT_EXIST)

		return
	}

	user.UpdateLoginTime()

	// 验证是否在黑名单
	if middlewares.IsExistInBlackSet(user.ID) {
		utils.Send(context, nil, robust.JWT_NOT_ALLOWED)

		return
	}

	claims := utils.Claims{
		UserId: user.ID,
	}

	utils.Send(context, gin.H{
		"jwt":      claims.SignJWT(),
		"userInfo": user,
	}, nil)
}

func GetUserInfo(context *gin.Context) {
	claims, isExist := context.Get("claims")

	fmt.Println(claims)

	if !isExist {
		utils.Send(context, nil, robust.JWT_PARSE_ERROR)

		return
	}

	parsedClaims, ok := claims.(utils.Claims)

	if !ok {
		utils.Send(context, nil, robust.JWT_PARSE_ERROR)

		return
	}

	user := models.User{
		ID: parsedClaims.UserId,
	}

	utils.Send(context, gin.H{
		"userInfo": user.GetUserInfoByID(),
	}, nil)
}

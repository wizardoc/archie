package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/** 用户注册 */
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

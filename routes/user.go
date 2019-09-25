package routes

import (
	"archie/controllers/user_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func userRouter(router *gin.Engine) {
	user := router.Group("/user")

	user.POST("/valid/info/base", user_controller.ValidBaseInfo)
	user.POST("/register", user_controller.Register)
	user.POST("/login", user_controller.Login)
	user.GET("/info", middlewares.ValidateToken, user_controller.GetUserInfo)
}

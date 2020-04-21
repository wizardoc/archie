package routes

import (
	"archie/controllers/user_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func userRouter(router *gin.Engine) {
	user := router.Group("/user")

	user.POST("/valid/info/base", user_controller.ValidBaseInfo)
	user.POST("/register", user_controller.Register, user_controller.Login)
	user.POST("/login", user_controller.Login)
	user.GET("/info", middlewares.ValidateToken, user_controller.GetUserInfo)
	user.PUT("/avatar", middlewares.ValidateToken, user_controller.UpdateAvatar)
	user.GET("/name/search", user_controller.SearchName)
}

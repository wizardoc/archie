package routes

import (
	"archie/controllers"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func userRouter(router *gin.Engine) {
	user := router.Group("/user")

	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	user.GET("/info", middlewares.ValidateToken, controllers.GetUserInfo)
}

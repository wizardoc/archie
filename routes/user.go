package routes

import (
	"archie/controllers"
	"github.com/gin-gonic/gin"
)

func userRouter(router *gin.Engine) {
	user := router.Group("/user")

	user.GET("/register", controllers.Register)
	user.POST("/login")
}

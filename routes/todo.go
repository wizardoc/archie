package routes

import (
	"archie/controllers/user_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func TodoRouter(router *gin.Engine) {
	todo := router.Group("/todo")

	todo.POST("/add", middlewares.ValidateToken, user_controller.AddTodo)
	todo.DELETE("/remove", middlewares.ValidateToken, user_controller.RemoveTodo)
}

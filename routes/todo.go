package routes

import (
	"archie/controllers/todo_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func TodoRouter(router *gin.Engine) {
	todo := router.Group("/todo")

	todo.POST("/new", middlewares.ValidateToken, todo_controller.AddTodo)
	todo.DELETE("/remove", middlewares.ValidateToken, todo_controller.RemoveTodo)
}

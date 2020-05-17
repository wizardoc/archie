package routes

import (
	"archie/controllers/todo_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func todoRouter(router *gin.Engine) {
	todo := router.Group("/todo", middlewares.ValidateToken)

	todo.POST("/new", todo_controller.AddTodo)
	todo.DELETE("/remove", todo_controller.RemoveTodo)
	todo.GET("/all", todo_controller.GetAllTodo)
}

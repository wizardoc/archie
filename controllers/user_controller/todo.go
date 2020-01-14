package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

/** 添加待办事项 */
func AddTodo(context *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(context)
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(context)
		return
	}

	name := context.PostForm("name")
	description := context.PostForm("description")
	route := context.PostForm("route")

	todoItem := models.UserTodo{
		UserID:      parsedClaims.UserId,
		Name:        name,
		Description: description,
		Route:       route,
	}

	todoItem.AddUserTodoItem()

	res.Send(context)
}

/** 删除待办事项 */
func RemoveTodo(context *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(context)
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(context)
		return
	}

	name := context.PostForm("name")
	todoItem := models.UserTodo{
		Name:   name,
		UserID: parsedClaims.UserId,
	}

	todoItem.RemoveUserTodoItem()

	res.Send(context)
}

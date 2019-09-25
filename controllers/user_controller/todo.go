package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

/** 添加待办事项 */
func AddTodo(context *gin.Context) {
	parsedClaims, archieErr := middlewares.GetClaims(context)

	if archieErr.Msg != "" {
		helper.Send(context, nil, archieErr)

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

	helper.Send(context, nil, nil)
}

/** 删除待办事项 */
func RemoveTodo(context *gin.Context) {
	parsedClaims, archieErr := middlewares.GetClaims(context)

	if archieErr.Msg != "" {
		helper.Send(context, nil, archieErr)

		return
	}

	name := context.PostForm("name")
	todoItem := models.UserTodo{
		Name:   name,
		UserID: parsedClaims.UserId,
	}

	todoItem.RemoveUserTodoItem()

	helper.Send(context, nil, nil)
}

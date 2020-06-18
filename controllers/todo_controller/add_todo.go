package todo_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddTodoPayload struct {
	Name        string `form:"name" validate:"required"`
	Description string `form:"description" validate:"required"`
	Route       string `form:"route" validate:"required"`
}

/** 添加待办事项 */
func AddTodo(ctx *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(ctx)
	res := helper.Res{}

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	var payload AddTodoPayload
	if err := helper.BindWithValid(ctx, &payload); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)

		return
	}

	todoItem := models.UserTodo{
		UserID:      parsedClaims.User.ID,
		Name:        payload.Name,
		Description: payload.Description,
		Route:       payload.Route,
	}

	if err := todoItem.AddUserTodoItem(); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	res.Send(ctx, nil)
}

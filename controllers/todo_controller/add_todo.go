package todo_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
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
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	var payload AddTodoPayload
	if err := helper.BindWithValid(ctx, &payload); err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	todoItem := models.UserTodo{
		UserID:      parsedClaims.UserId,
		Name:        payload.Name,
		Description: payload.Description,
		Route:       payload.Route,
	}

	if err := todoItem.AddUserTodoItem(); err != nil {
		authRes.Err = robust.DOUBLE_KEY
		authRes.Send(ctx)
		return
	}

	res.Send(ctx)
}

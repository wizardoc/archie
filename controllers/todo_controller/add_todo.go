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
func AddTodo(context *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(context)
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(context)
		return
	}

	var payload AddTodoPayload
	if err := helper.BindWithValid(context, &payload); err != nil {
		authRes.Err = err
		authRes.Send(context)
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
		authRes.Send(context)
		return
	}

	res.Send(context)
}

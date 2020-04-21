package todo_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RemoveTodoPayload struct {
	Name string `form:"name" validate:"required"`
}

/** 删除待办事项 */
func RemoveTodo(ctx *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(ctx)
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	var payload RemoveTodoPayload
	if err := helper.BindWithValid(ctx, &payload); err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	todoItem := models.UserTodo{
		Name:   payload.Name,
		UserID: parsedClaims.UserId,
	}

	if err := todoItem.RemoveUserTodoItem(); err != nil {
		authRes.Err = robust.DOUBLE_KEY
		authRes.Send(ctx)
		return
	}

	res.Send(ctx)
}

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
	res := helper.Res{}

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, robust.DB_READ_FAILURE)
		return
	}

	var payload RemoveTodoPayload
	if err := helper.BindWithValid(ctx, &payload); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, robust.DB_READ_FAILURE)
		return
	}

	todoItem := models.UserTodo{
		Name:   payload.Name,
		UserID: parsedClaims.UserId,
	}

	if err := todoItem.RemoveUserTodoItem(); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, robust.DB_READ_FAILURE)
		return
	}

	res.Send(ctx, nil)
}

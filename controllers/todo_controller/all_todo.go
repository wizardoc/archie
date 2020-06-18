package todo_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllTodo(ctx *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(ctx)
	res := helper.Res{}

	if err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	todo := models.UserTodo{
		UserID: parsedClaims.User.ID,
	}
	todos, err := todo.GetAllTodoItemsByID()

	if err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, robust.DB_READ_FAILURE)
		return
	}

	res.Send(ctx, todos)
}

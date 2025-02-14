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
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	todo := models.UserTodo{
		UserID: parsedClaims.User.ID,
	}
	todos, err := todo.GetAllTodoItemsByID()

	if err != nil {
		res.Status(http.StatusBadRequest).Error(robust.DB_READ_FAILURE).Send(ctx)
		return
	}

	res.Success(todos).Send(ctx)
}

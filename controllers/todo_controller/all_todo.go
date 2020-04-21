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
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	todo := models.UserTodo{
		UserID: parsedClaims.UserId,
	}
	todos, err := todo.GetAllTodoItemsByID()

	if err != nil {
		authRes.Err = robust.DB_READ_FAILURE
		authRes.Send(ctx)
		return
	}

	res.Data = todos
	res.Send(ctx)
}

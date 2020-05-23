package user_controller

import (
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchName(ctx *gin.Context) {
	var results []models.User
	searchName := ctx.Query("username")
	res := helper.Res{}
	user := models.User{}

	if searchName == "" {
		res.Send(ctx, []models.User{})
		return
	}

	if err := user.SearchName(searchName, &results); err != nil {
		res.Status(http.StatusInternalServerError).Error(ctx, err)
		return
	}

	res.Send(ctx, results)
}

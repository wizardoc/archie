package user_controller

import (
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

func SearchName(ctx *gin.Context) {
	var results []models.User

	serverErrRes := helper.GenServerErrRes()
	successRes := helper.GenSuccessRes()
	searchName := ctx.Query("username")
	user := models.User{}

	if searchName == "" {
		successRes.Data = []models.User{}
		successRes.Send(ctx)
		return
	}

	if err := user.SearchName(searchName, &results); err != nil {
		serverErrRes.Err = err
		serverErrRes.Send(ctx)
		return
	}

	successRes.Data = results
	successRes.Send(ctx)
}

package user_controller

import (
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserDetail(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	user := models.User{ID: id}
	res := helper.Res{}

	if err := user.Find("id", id); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	res.Send(ctx, user)
}

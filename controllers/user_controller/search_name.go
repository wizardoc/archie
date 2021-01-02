package user_controller

import (
	"archie/models"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchNameParams struct {
	SearchName string `validate:"required" json:"searchName" form:"searchName"`
	utils.PageInfo
}

func SearchName(ctx *gin.Context) {
	var results []models.User
	res := helper.Res{}
	user := models.User{}

	var params SearchNameParams

	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	if params.SearchName == "" {
		res.Success([]models.User{}).Send(ctx)
		return
	}

	params.ParsePageInfo()

	if err := user.SearchName(params.SearchName, params.Page, params.PageSize, &results); err != nil {
		res.Status(http.StatusInternalServerError).Error(err).Send(ctx)
		return
	}

	res.Success(results).Send(ctx)
}

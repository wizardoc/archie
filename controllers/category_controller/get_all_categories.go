package category_controller

import (
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllCategoriesParams struct {
	OrganizationID string `form:"organizationID" validate:"required"`
}

func GetAllCategories(ctx *gin.Context) {
	payload := GetAllCategoriesParams{}
	res := helper.Res{}

	if err := helper.BindWithValid(ctx, &payload); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	category := models.Category{
		OrganizationID: payload.OrganizationID,
	}
	var results []models.ResCategory

	if err := category.All(&results); err != nil {
		res.Status(http.StatusInternalServerError).Error(err).Send(ctx)
		return
	}

	res.Success(results).Send(ctx)
}

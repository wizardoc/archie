package category_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCategoryParams struct {
	Name           string `form:"name" validate:"lt=15,gt=0,required" json:"name"`
	Description    string `form:"description" validate:"lt=50,gt=0,required" json:"description"`
	Cover          string `form:"cover" json:"cover" validate:"required"`
	OrganizationID string `form:"organizationID" json:"organizationID" validate:"required"`
}

func CreateCategory(ctx *gin.Context) {
	payload := CreateCategoryParams{}
	res := helper.Res{}

	claims, err := middlewares.GetClaims(ctx)

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	if err := helper.BindWithValid(ctx, &payload); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	category := models.Category{
		Name:           payload.Name,
		Description:    payload.Description,
		Cover:          payload.Cover,
		UserID:         claims.User.ID,
		OrganizationID: payload.OrganizationID,
	}

	if err := category.New(); err != nil {
		res.Status(http.StatusInternalServerError).Error(ctx, err)
		return
	}

	res.Data = category
	res.Send(ctx, nil)
}

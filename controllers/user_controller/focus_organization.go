package user_controller

import (
	"archie/middlewares"
	"archie/models/focus_models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FocusOrganizationParams struct {
	OrganizationID string `form:"organizationID" json:"organizationId" validate:"required"`
}

func FocusOrganization(ctx *gin.Context) {
	var res helper.Res
	var params FocusOrganizationParams

	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	claims, err := middlewares.GetClaims(ctx)

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	fo := focus_models.FocusOrganization{
		UserID:         claims.ID,
		OrganizationID: params.OrganizationID,
	}
	if err := fo.New(); err != nil {
		res.Status(http.StatusForbidden).Error(ctx, err)
		return
	}

	res.Send(ctx, nil)
}

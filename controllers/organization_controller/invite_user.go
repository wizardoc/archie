package organization_controller

import (
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InviteUserParams struct {
	Username         string `json:"username" form:"username" validate:"required"`
	OrganizationName string `json:"organizationName" form:"organizationName" validate:"required"`
}

func InviteUser(ctx *gin.Context) {
	var inviteUserParams InviteUserParams
	res := helper.Res{}

	if err := helper.BindWithValid(ctx, &inviteUserParams); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	res.Send(ctx, nil)
}

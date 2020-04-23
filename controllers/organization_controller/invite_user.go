package organization_controller

import (
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

type InviteUserParams struct {
	Username         string `json:"username" form:"username" validate:"required"`
	OrganizationName string `json:"organizationName" form:"organizationName" validate:"required"`
}

func InviteUser(ctx *gin.Context) {
	var inviteUserParams InviteUserParams

	badReqRes := helper.GenBadReqRes()
	successRes := helper.GenSuccessRes()

	if err := helper.BindWithValid(ctx, &inviteUserParams); err != nil {
		badReqRes.Err = err
		badReqRes.Send(ctx)
		return
	}

	successRes.Send(ctx)
}

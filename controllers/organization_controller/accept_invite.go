package organization_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"archie/utils/jwt_utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AcceptInviteParams struct {
	OrganizeName string `json:"organizeName" validate:"required" form:"organizeName"`
}

func AcceptInvite(ctx *gin.Context) {
	inviteToken := ctx.Params.ByName("inviteToken")
	var params AcceptInviteParams
	res := helper.Res{}

	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	var inviteClaims jwt_utils.InviteClaims
	if err := middlewares.ParseToken2Claims(inviteToken, &inviteClaims); err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	if inviteClaims.OrganizeName != params.OrganizeName {
		res.Status(http.StatusBadRequest).Error(ctx, robust.ORGANIZATION_INVITE_ERROR)
		return
	}

	targetOrg := models.Organization{OrganizeName: inviteClaims.OrganizeName}
	if err := targetOrg.FindOneByOrganizeName(); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	if err := InsertUserToOrganization(inviteClaims.OrganizeName, inviteClaims.InviteUser, false); err != nil {
		res.Status(http.StatusInternalServerError).Error(ctx, err)
		return
	}

	res.Send(ctx, targetOrg)
}
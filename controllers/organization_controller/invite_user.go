package organization_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/services"
	"archie/utils/helper"
	"archie/utils/jwt_utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InviteUserParams struct {
	Username     string `json:"username" form:"username" validate:"required"`
	OrganizeName string `json:"organizeName" form:"organizeName" validate:"required"`
}

func InviteUser(ctx *gin.Context) {
	var inviteUserParams InviteUserParams
	res := helper.Res{}

	if err := helper.BindWithValid(ctx, &inviteUserParams); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	claims, err := middlewares.GetClaims(ctx)
	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	// 邀请自己
	if claims.Username == inviteUserParams.Username {
		res.Status(http.StatusBadRequest).Error(ctx, robust.ORGANIZATION_INVITE_YOURSELF)
		return
	}

	targetOrg := models.Organization{OrganizeName: inviteUserParams.OrganizeName}
	if err := targetOrg.FindOneByOrganizeName(); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, robust.ORGANIZATION_FIND_EMPTY)
		return
	}

	inviteUser := models.User{}
	if err := inviteUser.FindByUsername(inviteUserParams.Username); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	userOrganization := models.UserOrganization{
		UserID:         inviteUser.ID,
		OrganizationID: targetOrg.ID,
	}
	// 邀请组内已经存在的人员
	isExist, err := userOrganization.IsExist()

	if err != nil {
		res.Status(http.StatusInternalServerError).Send(ctx, err)
		return
	}

	if isExist {
		res.Status(http.StatusBadRequest).Error(ctx, robust.ORGANIZATION_INVITE_EXIST)
		return
	}

	inviteClaims := jwt_utils.InviteClaims{
		InviteUser:   inviteUserParams.Username,
		OrganizeName: inviteUserParams.OrganizeName,
	}
	inviteToken := inviteClaims.SignJWT(1)

	msg := services.Message{
		From: claims.ID,
		To:   inviteUser.ID,
	}
	if err := msg.SendInviteMessage(inviteToken, claims.Username, inviteUserParams.OrganizeName); err != nil {
		res.Status(http.StatusInternalServerError).Error(ctx, err)
		return
	}

	res.Send(ctx, nil)
}

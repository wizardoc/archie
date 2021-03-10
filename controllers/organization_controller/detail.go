package organization_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrganizationDetailData struct {
	models.Organization
	Members []models.User `json:"members"`
}

func OrganizationDetail(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	res := helper.Res{}
	claims, err := middlewares.GetClaims(ctx)

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	userOrg := models.UserOrganization{
		OrganizationID: id,
		UserID:         claims.ID,
	}

	// 成员不在依然可以查看组织详情
	//isExist, err := userOrg.IsExist()
	//if err != nil {
	//	res.Status(http.StatusForbidden).Error(err).Send(ctx)
	//	return
	//}
	//
	//if !isExist {
	//	res.Status(http.StatusForbidden).Error(robust.INVALID_PERMISSION).Send(ctx)
	//	return
	//}

	org := models.Organization{ID: id}
	if err := org.FindOneByID(); err != nil {
		res.Status(http.StatusNotFound).Error(err).Send(ctx)
		return
	}

	var members []models.User
	if err := userOrg.FindMembers(id, &members); err != nil {
		res.Status(http.StatusForbidden).Error(err).Send(ctx)
		return
	}

	resData := OrganizationDetailData{
		Organization: org,
		Members:      members,
	}

	res.Success(resData).Send(ctx)
}

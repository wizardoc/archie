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
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	userOrg := models.UserOrganization{
		OrganizationID: id,
		UserID:         claims.ID,
	}

	// 成员不在依然可以查看组织详情
	//isExist, err := userOrg.IsExist()
	//if err != nil {
	//	res.Status(http.StatusForbidden).Error(ctx, err)
	//	return
	//}
	//
	//if !isExist {
	//	res.Status(http.StatusForbidden).Error(ctx, robust.INVALID_PERMISSION)
	//	return
	//}

	org := models.Organization{}
	if err := org.FindOneByID(id); err != nil {
		res.Status(http.StatusNotFound).Error(ctx, err)
		return
	}

	var members []models.User
	if err := userOrg.FindMembers(id, &members); err != nil {
		res.Status(http.StatusForbidden).Error(ctx, err)
		return
	}

	resData := OrganizationDetailData{
		Organization: org,
		Members:      members,
	}

	res.Send(ctx, resData)
}

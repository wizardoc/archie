package organization_controller

import (
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type EditOrganizationParams struct {
	OrganizeName string `json:"organizeName" form:"organizeName" validate:"required"`
	Description  string `json:"description" form:"description" validate:"required"`
}

// 编辑组织信息，传递一个组织信息，按需更新
func EditOrganization(ctx *gin.Context) {
	var organizationInfo EditOrganizationParams
	res := helper.Res{}
	authRes := helper.Res{}
	err := helper.BindWithValid(ctx, &organizationInfo)
	id := ctx.Params.ByName("id")

	if err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	organization := models.Organization{
		ID: id,
	}
	updates := make(map[string]interface{})

	if err := mapstructure.Decode(organizationInfo, &updates); err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	if err := organization.BatchUpdates(updates); err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	res.Send(ctx)
}

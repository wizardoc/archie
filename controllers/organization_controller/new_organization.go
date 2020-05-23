package organization_controller

import (
	"archie/models"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrganizationUser struct {
	OrganizeName string `form:"organizeName" validate:"required"`
	Description  string `form:"organizeDescription" validate:"required"`
	Username     string `form:"username" validate:"required"`
}

// 创建一个新的组织，创建组织成功后将用户插入至该组织
func NewOrganization(ctx *gin.Context) {
	var organizationUser OrganizationUser
	res := helper.Res{}

	if err := helper.BindWithValid(ctx, &organizationUser); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	organization := models.Organization{}
	utils.CpStruct(&organizationUser, &organization)

	if err := organization.New(organizationUser.Username); err != nil {
		res.Status(http.StatusConflict).Error(ctx, err)
		return
	}

	res.Send(ctx, nil)
}

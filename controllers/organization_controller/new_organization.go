package organization_controller

import (
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrganizationUser struct {
	OrganizeName        string `form:"organizeName" validate:"required"`
	OrganizeDescription string `form:"organizeDescription" validate:"required"`
	Username            string `form:"username" validate:"required"`
}

// 创建一个新的组织，创建组织成功后将用户插入至该组织
func NewOrganization(context *gin.Context) {
	var organizationUser OrganizationUser
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err := helper.BindWithValid(context, &organizationUser); err != nil {
		authRes.Err = err
		authRes.Send(context)
		return
	}

	if err := CreateNewOrganization(
		organizationUser.OrganizeName,
		organizationUser.OrganizeDescription,
		organizationUser.Username,
	); err != nil {
		res.Err = robust.DOUBLE_KEY
		res.Send(context)
		return
	}

	InsertUserToOrganization(organizationUser.OrganizeName, organizationUser.Username, true)
	res.Send(context)
}

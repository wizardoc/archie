package organization_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

func GetAllOrganizationNames(ctx *gin.Context) {
	var organization models.Organization
	names, err := organization.GetAllNames()
	res := helper.Res{}

	if err != nil {
		res.Err = robust.CANNOT_FIND_ORGANIZATION
		res.Send(ctx)
		return
	}

	res.Data = gin.H{
		"organizeNames": names,
	}
	res.Send(ctx)
}

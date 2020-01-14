package organization_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

func GetAllOrganizationNames(context *gin.Context) {
	var organization models.Organization
	names, ok := organization.GetAllNames()
	res := helper.Res{}

	if !ok {
		res.Err = robust.CANNOT_FIND_ORGANIZATION
		res.Send(context)
		return
	}

	res.Data = gin.H{
		"organizeNames": names,
	}
	res.Send(context)
}

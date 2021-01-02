package organization_controller

import (
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllOrganizationNames(ctx *gin.Context) {
	var organization models.Organization
	names, err := organization.GetAllNames()
	res := helper.Res{}

	if err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	res.Success(gin.H{
		"organizeNames": names,
	}).Send(ctx)
}

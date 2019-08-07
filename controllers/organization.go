package controllers

import (
	"archie/models"
	"archie/robust"
	"archie/utils"
	"github.com/gin-gonic/gin"
)

func GetAllOrganizationNames(context *gin.Context) {
	var organization models.Organization
	names, ok := organization.GetAllNames()

	if !ok {
		utils.Send(context, nil, robust.CANNOT_FIND_ORGANIZATION)
	}

	context.JSON(200, gin.H{
		"organizeNames": names,
	})
}

func NewOrganization(context *gin.Context) {
	organizeName := context.PostForm("name")
	organizeDescription := context.PostForm("description")

	ok := CreateNewOrganization(organizeName, organizeDescription)

	if !ok {
		utils.Send(context, nil, robust.CONNOT_CREATE_ORGANIZATION)
	}

	context.JSON(200, gin.H{
		"ok": true,
	})
}

func CreateNewOrganization(name string, description string) (ok bool) {
	organization := models.Organization{
		OrganizeName: name,
		Description:  description,
	}

	return organization.New()
}

package controllers

import (
	"archie/models"
	"archie/robust"
	"github.com/gin-gonic/gin"
)

func GetAllOrganizationNames(context *gin.Context) {
	var organization models.Organization
	names, ok := organization.GetAllNames()

	if !ok {
		context.Error(robust.CANNOT_FIND_ORGANIZATION)
	}

	context.JSON(200, gin.H{
		"organizeNames": names,
	})
}

func NewOrganization(context *gin.Context) {
	organizeName := context.PostForm("name")
	organizeDescription := context.PostForm("description")

	organization := models.Organization{
		OrganizeName: organizeName,
		Description:  organizeDescription,
	}

	ok := organization.New()

	if !ok {
		context.Error(robust.CONNOT_CREATE_ORGANIZATION)
	}

	context.JSON(200, gin.H{
		"ok": true,
	})
}

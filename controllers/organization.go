package controllers

import (
	"archie/models"
	"archie/robust"
	"archie/utils"
	"github.com/gin-gonic/gin"
)

func insertUserToOrganization(organizeName string, username string, isOwner bool) {
	organization := models.Organization{OrganizeName: organizeName}
	organization.FindOneByOrganizeName()

	user := models.FindOneByUsername(username)
	userOrganization := models.UserOrganization{
		UserID:         user.ID,
		OrganizationID: organization.ID,
	}

	userOrganization.New(isOwner)
}

func GetAllOrganizationNames(context *gin.Context) {
	var organization models.Organization
	names, ok := organization.GetAllNames()

	if !ok {
		utils.Send(context, nil, robust.CANNOT_FIND_ORGANIZATION)

		return
	}

	utils.Send(context, gin.H{
		"organizeNames": names,
	}, nil)
}

func NewOrganization(context *gin.Context) {
	organizeName := context.PostForm("organizeName")
	organizeDescription := context.PostForm("organizeDescription")
	username := context.PostForm("username")

	ok := CreateNewOrganization(organizeName, organizeDescription)

	if !ok {
		utils.Send(context, nil, robust.CONNOT_CREATE_ORGANIZATION)

		return
	}

	insertUserToOrganization(organizeName, username, true)

	utils.Send(context, "success", nil)
}

func JoinOrganization(context *gin.Context) {
	organizeName := context.PostForm("organizeName")
	username := context.PostForm("username")

	insertUserToOrganization(organizeName, username, false)

	utils.Send(context, "success", nil)
}

func CreateNewOrganization(name string, description string) (ok bool) {
	organization := models.Organization{
		OrganizeName: name,
		Description:  description,
	}

	return organization.New()
}

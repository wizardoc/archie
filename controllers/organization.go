package controllers

import (
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
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

func NewOrganization(context *gin.Context) {
	organizeName := context.PostForm("organizeName")
	organizeDescription := context.PostForm("organizeDescription")
	username := context.PostForm("username")
	res := helper.Res{}

	ok := CreateNewOrganization(organizeName, organizeDescription, username)

	if !ok {
		res.Err = robust.CONNOT_CREATE_ORGANIZATION
		res.Send(context)
		return
	}

	insertUserToOrganization(organizeName, username, true)
	res.Send(context)
}

func JoinOrganization(context *gin.Context) {
	organizeName := context.PostForm("organizeName")
	username := context.PostForm("username")
	res := helper.Res{}

	insertUserToOrganization(organizeName, username, false)

	res.Send(context)
}

func CreateNewOrganization(name string, description string, username string) (ok bool) {
	organization := models.Organization{
		OrganizeName: name,
		Description:  description,
	}

	return organization.New(username)
}

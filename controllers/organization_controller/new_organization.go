package organization_controller

import (
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

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

	InsertUserToOrganization(organizeName, username, true)
	res.Send(context)
}

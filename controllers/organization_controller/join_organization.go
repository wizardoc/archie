package organization_controller

import (
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

func JoinOrganization(context *gin.Context) {
	organizeName := context.PostForm("organizeName")
	username := context.PostForm("username")
	res := helper.Res{}

	InsertUserToOrganization(organizeName, username, false)

	res.Send(context)
}

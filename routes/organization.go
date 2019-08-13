package routes

import (
	"archie/controllers"
	"github.com/gin-gonic/gin"
)

func organizationRouter(router *gin.Engine) {
	organization := router.Group("/organization")

	organization.GET("/name/all", controllers.GetAllOrganizationNames)
	organization.POST("/new", controllers.NewOrganization)
	organization.POST("/join", controllers.JoinOrganization)
}

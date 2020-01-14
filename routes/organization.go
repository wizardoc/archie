package routes

import (
	"archie/controllers/organization_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func organizationRouter(router *gin.Engine) {
	organization := router.Group("/organization")

	organization.GET("/name/all", organization_controller.GetAllOrganizationNames)
	organization.GET("/joins/all", middlewares.ValidateToken, organization_controller.GetAllJoinOrganization)
	organization.DELETE("/remove/:name", middlewares.ValidateToken, organization_controller.RemoveOwnOrganization)
	organization.POST("/new", organization_controller.NewOrganization)
	organization.POST("/join", organization_controller.JoinOrganization)
}

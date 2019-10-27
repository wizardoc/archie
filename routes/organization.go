package routes

import (
	"archie/controllers"
	"archie/controllers/organization_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func organizationRouter(router *gin.Engine) {
	organization := router.Group("/organization")

	organization.GET("/name/all", controllers.GetAllOrganizationNames)
	organization.GET("/joins/all", middlewares.ValidateToken, organization_controller.GetAllJoinOrganization)
	organization.DELETE("/remove/:name", middlewares.ValidateToken, organization_controller.RemoveOwnOrganization)
	organization.POST("/new", controllers.NewOrganization)
	organization.POST("/join", controllers.JoinOrganization)
}

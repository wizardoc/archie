package routes

import (
	"archie/controllers"
	"github.com/gin-gonic/gin"
)

func organizationRouter(router *gin.Engine) {
	organization := router.Group("/organization")

	organization.GET("/all", controllers.GetAllOrganizationNames)
}

package routes

import (
	"archie/controllers/permission_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func permissionRouter(router *gin.Engine) {
	permission := router.Group("/permission")

	permission.GET("/organization/:id", middlewares.ValidateToken, permission_controller.GetOrganizationPermission)
}

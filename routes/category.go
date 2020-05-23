package routes

import (
	"archie/controllers/category_controller"
	"archie/middlewares"
	"archie/middlewares/permission_validators"
	"archie/models"
	"github.com/gin-gonic/gin"
)

func categoryRouter(router *gin.Engine) {
	category := router.Group("/category")

	category.POST(
		"/",
		middlewares.ValidateToken,
		permission_validators.OrganizationPermission([]int{models.CATEGORY_CREATE}),
		category_controller.CreateCategory,
	)

	category.GET("/", category_controller.GetAllCategories)
}

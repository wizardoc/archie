package routes

import (
	"archie/controllers"
	"github.com/gin-gonic/gin"
)

func DocRouter(router *gin.Engine) {
	doc := router.Group("/doc")

	doc.GET("/wizard", controllers.WizardIntroduction)
}

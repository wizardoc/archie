package routes

import (
	"archie/controllers/doc_controller"
	"github.com/gin-gonic/gin"
)

func DocRouter(router *gin.Engine) {
	doc := router.Group("/doc")

	doc.GET("/wizard", doc_controller.WizardIntroduction)
}

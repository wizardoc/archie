package routes

import (
	"archie/controllers/doc_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func docRouter(router *gin.Engine) {
	doc := router.Group("/doc")

	doc.GET("/wizard", doc_controller.WizardIntroduction)
	doc.POST("/", middlewares.ValidateToken, doc_controller.NewDocument)
	doc.GET("/", doc_controller.GetAllDocuments)
	doc.POST("/comment", middlewares.ValidateToken, doc_controller.NewComment)
	doc.GET("/detail/:document_id", middlewares.ValidateToken, doc_controller.Detail)
	doc.GET("/comments/:document_id", doc_controller.GetAllComments)
	doc.PUT("/comment", middlewares.ValidateToken, doc_controller.UpdateCommentStatus)
}

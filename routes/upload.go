package routes

import (
	"archie/controllers/upload_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func uploadRouter(router *gin.Engine) {
	user := router.Group("/upload", middlewares.ValidateToken)

	user.GET("/qiniu/token", upload_controller.GetQiNiuToken)
}

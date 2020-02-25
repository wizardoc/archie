package routes

import (
	"archie/controllers/upload_controller"
	"github.com/gin-gonic/gin"
)

func uploadRouter(router *gin.Engine) {
	user := router.Group("/upload")

	user.GET("/qiniu/token", upload_controller.GetQiNiuToken)
}

package upload_controller

import (
	"archie/services"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

func GetQiNiuToken(ctx *gin.Context) {
	qiniu := services.QiNiu{}
	res := helper.Res{}

	qiniu.New()
	token := qiniu.GenToken()

	res.Success(token).Send(ctx)
}

package upload_controller

import (
	"archie/services"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

func GetQiNiuToken(context *gin.Context) {
	qiniu := services.QiNiu{}

	qiniu.New()

	token := qiniu.GenToken()
	res := helper.Res{
		Data: token,
	}

	res.Send(context)
}

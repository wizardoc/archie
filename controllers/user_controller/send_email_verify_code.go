package user_controller

import (
	"archie/robust"
	"archie/services/email_service"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendEmailVerifyCodeParams struct {
	Email string `json:"email" form:"email" validate:"required"`
}

func SendEmailVerifyCode(ctx *gin.Context) {
	var res helper.Res
	var params SendEmailVerifyCodeParams

	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	emailService := email_service.EmailService{
		Email: params.Email,
	}

	if err := emailService.SendVerifyCode(); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, robust.SEND_VERIFY_CODE_FAILURE)
		return
	}

	emailService.SaveCode()

	res.Send(ctx, nil)
}

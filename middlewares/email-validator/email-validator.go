package email_validator

import (
	"archie/robust"
	"archie/services/email_service"
	"archie/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmailValidParams struct {
	Code  string `json:"code" form:"code" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
}

func EmailValidator(ctx *gin.Context) {
	var res helper.Res
	var params EmailValidParams

	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		ctx.Abort()
		return
	}

	emailService := email_service.EmailService{
		Email: params.Email,
		Code:  params.Code,
	}

	// (effect method) get the code from Redis and assign to Code of emailService
	emailService.GetCode()

	fmt.Println(emailService.Code)

	if emailService.Code == "" || params.Code != emailService.Code {
		res.Status(http.StatusBadRequest).Error(ctx, robust.INVALID_EMAIL_CODE)
		ctx.Abort()
		return
	}

	// (effect method) delete the code from Redis
	emailService.DelCode()

	ctx.Set("email", params.Email)
	ctx.Next()
}

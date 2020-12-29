package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateEmail(ctx *gin.Context) {
	var res helper.Res
	email, exist := ctx.Get("email")

	if !exist {
		res.Status(http.StatusBadRequest).Error(ctx, robust.MISSING_PARAMS)
		return
	}

	claims, err := middlewares.GetClaims(ctx)
	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	user := models.User{
		ID: claims.ID,
	}
	if err := user.GetUserInfoByID(); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	// refuse the request if the previous email has not been successful verify
	if !user.IsValidEmail {
		res.Status(http.StatusForbidden).Error(ctx, robust.NO_VALID_EMAIL)
		return
	}

	// refuse the request if the previous email is equal to current email
	if email == user.Email {
		res.Status(http.StatusForbidden).Error(ctx, robust.REPEAT_EMAIL)
		return
	}

	user.Email = email.(string)
	user.IsValidEmail = true
	if err := user.UpdateUserInfo(); err != nil {
		res.Status(http.StatusForbidden).Error(ctx, err)
		return
	}

	res.Send(ctx, user)
}

package user_controller

import (
	"archie/middlewares"
	"archie/models/focus_models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FocusUserParams struct {
	UserID string `json:"userID" form:"userID" validate:"required"`
}

func FocusUser(ctx *gin.Context) {
	var res helper.Res
	var params FocusUserParams

	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	claims, err := middlewares.GetClaims(ctx)
	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	if claims.ID == params.UserID {
		res.Status(http.StatusConflict).Error(ctx, robust.CONNOT_FOLLOW_YOURSELF)
		return
	}

	fu := focus_models.FocusUser{
		UserID:      claims.ID,
		FocusUserID: params.UserID,
	}

	if err := fu.New(); err != nil {
		res.Status(http.StatusForbidden).Error(ctx, robust.REPEAT_FOLLOW_USER)
		return
	}

	res.Send(ctx, err)
}

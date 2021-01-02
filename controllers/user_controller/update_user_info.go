package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateUserInfoParams struct {
	ID           string
	DisplayName  string `json:"displayName" form:"displayName"`
	RealName     string `json:"realName" form:"realName"`
	Intro        string `json:"intro" form:"intro"`
	City         string `json:"city" form:"city"`
	CompanyName  string `json:"companyName" form:"companyName"`
	CompanyTitle string `json:"companyTitle" form:"companyTitle"`
	Github       string `json:"github" form:"github"`
	Blog         string `json:"blog" form:"blog"`
}

func UpdateUserInfo(ctx *gin.Context) {
	var params UpdateUserInfoParams
	var res helper.Res

	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	claims, err := middlewares.GetClaims(ctx)
	if err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	params.ID = claims.ID
	var user models.User

	utils.CpStruct(&params, &user)

	if err := user.UpdateUserInfo(); err != nil {
		res.Status(http.StatusForbidden).Error(robust.CANNOT_UPDATE_USERINFO).Send(ctx)
		return
	}

	res.Send(ctx)
}

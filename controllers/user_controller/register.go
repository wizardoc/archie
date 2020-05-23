package user_controller

import (
	"archie/models"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterInfo struct {
	Username                string `json:"username" form:"username" validate:"gt=4,lt=20,required"`
	Password                string `form:"password" validate:"required,gt=4,lt=20"`
	DisplayName             string `json:"displayName" form:"displayName" validate:"required,gt=2,lt=10"`
	Email                   string `json:"email" form:"email" validate:"email,required"`
	OrganizationName        string `json:"organizationName" form:"organizationName" validate:"required"`
	OrganizationDescription string `json:"organizationDescription" form:"organizationDescription" validate:"required"`
}

/** 用户注册 */
func Register(ctx *gin.Context) {
	res := helper.Res{}

	var info = RegisterInfo{}
	if err := helper.BindWithValid(ctx, &info); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		ctx.Abort()
		return
	}

	_, err := models.FindOneByUsername(info.Username)

	if err == nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)

		ctx.Abort()
		return
	}

	user := models.User{}
	utils.CpStruct(&info, &user)

	if err := user.Register(); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		ctx.Abort()
		return
	}

	organization := models.Organization{
		OrganizeName: info.OrganizationName,
		Description:  info.OrganizationDescription,
	}
	if err := organization.New(user.Username); err != nil {
		res.Status(http.StatusInternalServerError).Error(ctx, err)
		ctx.Abort()
		return
	}

	ctx.Next()
}

package user_controller

import (
	"archie/models"
	"archie/robust"
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
	errRes := helper.Res{Status: http.StatusBadRequest}
	serverErrRes := helper.Res{Status: http.StatusInternalServerError}

	var info = RegisterInfo{}
	if err := helper.BindWithValid(ctx, &info); err != nil {
		errRes.Err = err
		errRes.Send(ctx)
		ctx.Abort()
		return
	}

	_, err := models.FindOneByUsername(info.Username)

	if err == nil {
		errRes.Err = robust.REGISTER_EXIST_USER
		errRes.Send(ctx)
		ctx.Abort()
		return
	}

	user := models.User{}
	utils.CpStruct(&info, &user)

	if err := user.Register(); err != nil {
		errRes.Err = robust.CREATE_DATA_FAILURE
		errRes.Send(ctx)
		ctx.Abort()
		return
	}

	organization := models.Organization{
		OrganizeName: info.OrganizationName,
		Description:  info.OrganizationDescription,
	}
	if err := organization.New(user.Username); err != nil {
		serverErrRes.Err = robust.ORGANIZATION_CREATE_FAILURE
		serverErrRes.Send(ctx)
		ctx.Abort()
		return
	}

	ctx.Next()
}

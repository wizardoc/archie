package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"fmt"
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
func Register(context *gin.Context) {
	errRes := helper.Res{Status: http.StatusBadRequest}
	serverErrRes := helper.Res{Status: http.StatusInternalServerError}

	var info = RegisterInfo{}
	if err := helper.BindWithValid(context, &info); err != nil {
		errRes.Err = err
		errRes.Send(context)
		context.Abort()
		return
	}

	_, err := models.FindOneByUsername(info.Username)

	if err == nil {
		errRes.Err = robust.REGISTER_EXIST_USER
		errRes.Send(context)
		context.Abort()
		return
	}

	user := models.User{}
	utils.CpStruct(&info, &user)

	if err := user.Register(); err != nil {
		errRes.Err = robust.CREATE_DATA_FAILURE
		errRes.Send(context)
		context.Abort()
		return
	}

	organization := models.Organization{
		OrganizeName: info.OrganizationName,
		Description:  info.OrganizationDescription,
	}
	if err := organization.New(user.Username); err != nil {
		fmt.Println(err)

		serverErrRes.Err = robust.ORGANIZATION_CREATE_FAILURE
		serverErrRes.Send(context)
		context.Abort()
		return
	}

	context.Next()
}

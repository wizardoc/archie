package organization_controller

import (
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrganizationJoinInfo struct {
	OrganizeName string `form:"organizeName"`
	Username     string `form:"username"`
}

func JoinOrganization(ctx *gin.Context) {
	res := helper.Res{}
	authRes := helper.Res{Status: http.StatusBadRequest}

	var joinInfo OrganizationJoinInfo
	if err := ctx.Bind(&joinInfo); err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	if err := InsertUserToOrganization(joinInfo.OrganizeName, joinInfo.Username, false); err != nil {
		authRes.Err = robust.INVALID_PARAMS
		authRes.Send(ctx)
		return
	}

	res.Send(ctx)
}

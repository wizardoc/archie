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

func JoinOrganization(context *gin.Context) {
	res := helper.Res{}
	authRes := helper.Res{Status: http.StatusBadRequest}

	var joinInfo OrganizationJoinInfo
	if err := context.Bind(&joinInfo); err != nil {
		authRes.Err = err
		authRes.Send(context)
		return
	}

	if err := InsertUserToOrganization(joinInfo.OrganizeName, joinInfo.Username, false); err != nil {
		authRes.Err = robust.INVALID_PARAMS
		authRes.Send(context)
		return
	}

	res.Send(context)
}

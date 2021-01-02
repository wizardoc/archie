package organization_controller

import (
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

	var joinInfo OrganizationJoinInfo
	if err := ctx.Bind(&joinInfo); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)

		return
	}

	if err := InsertUserToOrganization(joinInfo.OrganizeName, joinInfo.Username, false); err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	res.Send(ctx)
}

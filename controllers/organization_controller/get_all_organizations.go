package organization_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取用户加入的所有组织
func GetAllJoinOrganization(ctx *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(ctx)
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	userOrganization := models.UserOrganization{}
	userOrganization.UserID = parsedClaims.UserId
	organizations, err := userOrganization.FindUserJoinOrganizations()

	if err != nil {
		authRes.Err = robust.ORGANIZATION_FIND_EMPTY
		authRes.Send(ctx)
		return
	}

	res.Data = gin.H{"organizations": organizations}
	res.Send(ctx)
}

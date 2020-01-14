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
func GetAllJoinOrganization(context *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(context)
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(context)
		return
	}

	userOrganization := models.UserOrganization{}
	userOrganization.UserID = parsedClaims.UserId
	organizations, err := userOrganization.FindUserJoinOrganizations()

	if err != nil {
		authRes.Err = robust.ORGANIZATION_FIND_EMPTY
		authRes.Send(context)
		return
	}

	res.Data = gin.H{"organizations": organizations}
	res.Send(context)
}

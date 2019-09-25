package organization_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

// 获取用户加入的所有组织
func GetAllJoinOrganization(context *gin.Context) {
	parsedClaims, archieErr := middlewares.GetClaims(context)

	if archieErr.Msg != "" {
		helper.Send(context, nil, archieErr)

		return
	}

	userOrganization := models.UserOrganization{}
	userOrganization.UserID = parsedClaims.UserId
	organizations, err := userOrganization.FindUserJoinOrganizations()

	if err != nil {
		helper.Send(context, nil, robust.ORGANIZATION_FIND_EMPTY)

		return
	}

	helper.Send(context, gin.H{"organizations": organizations}, nil)
}

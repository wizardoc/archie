package organization_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取用户加入的所有组织
func GetAllJoinOrganization(ctx *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(ctx)
	res := helper.Res{}

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	userOrganization := models.UserOrganization{}
	userOrganization.UserID = parsedClaims.User.ID
	organizations, err := userOrganization.FindUserJoinOrganizations()

	if err != nil {
		res.Status(http.StatusNotFound).Error(err).Send(ctx)
		return
	}

	res.Success(gin.H{"organizations": organizations}).Send(ctx)
}

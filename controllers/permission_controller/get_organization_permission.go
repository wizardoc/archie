package permission_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取对应组织的现有用户所拥有的所有权限
func GetOrganizationPermission(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	claims, err := middlewares.GetClaims(ctx)
	res := helper.Res{}

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	// 权限值的集合
	var permissionValues []int
	op := models.OrganizationPermission{UserID: claims.User.ID, OrganizationID: id}

	if err := op.AllAsValue(&permissionValues); err != nil {
		res.Status(http.StatusInternalServerError).Error(err).Send(ctx)
		return
	}

	res.Data = permissionValues
	res.Send(ctx)
}

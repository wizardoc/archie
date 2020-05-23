package permission_validators

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrganizationPermissionParams struct {
	OrganizationID string `form:"organizationID" validate:"required" json:"organizationID"`
}

// 验证该用户在指定组织的权限
func OrganizationPermission(limitPermissionValues []int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := OrganizationPermissionParams{}
		res := helper.Res{}

		if err := helper.BindWithValid(ctx, &payload); err != nil {
			res.Status(http.StatusBadRequest).Error(ctx, err)
			return
		}

		claims, err := middlewares.GetClaims(ctx)

		if err != nil {
			res.Status(http.StatusBadRequest).Error(ctx, err)
			return
		}

		// 验证权限
		op := models.OrganizationPermission{
			OrganizationID: payload.OrganizationID,
			UserID:         claims.UserId,
		}
		hasPermission, err := op.Has(limitPermissionValues)

		if err != nil {
			res.Status(http.StatusInternalServerError).Error(ctx, err)
			return
		}

		if !hasPermission {
			res.Status(http.StatusForbidden).Error(ctx, err)
			return
		}

		ctx.Next()
	}
}

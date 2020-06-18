package organization_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RemoveOwnOrganization(ctx *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(ctx)
	res := helper.Res{}

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	organizeName := ctx.Params.ByName("name")
	orgModel := models.Organization{
		OrganizeName: organizeName,
	}

	err = orgModel.FindOneByOrganizeName()

	if err != nil {
		res.Status(http.StatusNotFound).Error(ctx, err)

		return
	}

	if parsedClaims.User.ID != orgModel.Owner {
		res.Data = gin.H{
			"organizeName": "",
		}
		res.Status(http.StatusUnauthorized).Error(ctx, robust.REMOVE_PERMISSION)

		return
	}

	err = orgModel.RemoveOrganization()

	if err != nil {
		res.Status(http.StatusInternalServerError).Error(ctx, err)
		return
	}

	res.Send(ctx, nil)
}

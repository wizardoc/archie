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
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	organizeName := ctx.Params.ByName("name")
	orgModel := models.Organization{
		OrganizeName: organizeName,
	}

	err = orgModel.FindOneByOrganizeName()

	if err != nil {
		res.Status(http.StatusNotFound).Error(err).Send(ctx)

		return
	}

	if parsedClaims.User.ID != orgModel.Owner {
		res.Data = gin.H{
			"organizeName": "",
		}
		res.Status(http.StatusUnauthorized).Error(robust.REMOVE_PERMISSION).Send(ctx)

		return
	}

	err = orgModel.RemoveOrganization()

	if err != nil {
		res.Status(http.StatusInternalServerError).Error(err).Send(ctx)
		return
	}

	res.Send(ctx)
}

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
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(ctx)
		return
	}

	organizeName := ctx.Params.ByName("name")
	orgModel := models.Organization{
		OrganizeName: organizeName,
	}

	err = orgModel.FindOneByOrganizeName()

	if err != nil {
		authRes.Err = robust.ORGANIZATION_FIND_EMPTY
		authRes.Send(ctx)
		return
	}

	if parsedClaims.UserId != orgModel.Owner {
		authRes.Err = robust.REMOVE_PERMISSION
		authRes.Data = gin.H{
			"organizeName": "",
		}
		authRes.Send(ctx)
		return
	}

	err = orgModel.RemoveOrganization()

	if err != nil {
		authRes.Err = robust.REMOVE_ORG_FAILURE
		authRes.Send(ctx)
		return
	}

	res.Send(ctx)
}

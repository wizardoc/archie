package organization_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RemoveOwnOrganization(context *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(context)
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(context)
		return
	}

	organizeName := context.Params.ByName("name")
	orgModel := models.Organization{
		OrganizeName: organizeName,
	}

	orgModel.FindOneByOrganizeName()

	// 检验是否有权限删除组织
	if parsedClaims.UserId != orgModel.Owner {
		authRes.Err = robust.REMOVE_PERMISSION
		authRes.Data = gin.H{
			"organizeName": "",
		}
		authRes.Send(context)
		return
	}

	ok := orgModel.RemoveOrganization()

	if !ok {
		authRes.Err = robust.REMOVE_ORG_FAILURE
		authRes.Send(context)
		return
	}

	res.Send(context)
}

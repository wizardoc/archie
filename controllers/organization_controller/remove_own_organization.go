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

	err = orgModel.FindOneByOrganizeName()

	if err != nil {
		authRes.Err = robust.ORGANIZATION_FIND_EMPTY
		authRes.Send(context)
		return
	}

	// 检验是否有权限删除组织
	if parsedClaims.UserId != orgModel.Owner {
		authRes.Err = robust.REMOVE_PERMISSION
		authRes.Data = gin.H{
			"organizeName": "",
		}
		authRes.Send(context)
		return
	}

	err = orgModel.RemoveOrganization()

	if err != nil {
		authRes.Err = robust.REMOVE_ORG_FAILURE
		authRes.Send(context)
		return
	}

	res.Send(context)
}

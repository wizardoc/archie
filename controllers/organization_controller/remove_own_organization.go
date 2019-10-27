package organization_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
)

func RemoveOwnOrganization(context *gin.Context) {
	parsedClaims, archieErr := middlewares.GetClaims(context)

	if archieErr.Msg != "" {
		helper.Send(context, nil, archieErr)

		return
	}

	organizeName := context.Params.ByName("name")
	orgModel := models.Organization{
		OrganizeName: organizeName,
	}

	orgModel.FindOneByOrganizeName()

	// 检验是否有权限删除组织
	if parsedClaims.UserId != orgModel.Owner {
		helper.Send(context, gin.H{
			"organizeName": "",
		}, robust.REMOVE_PERMISSION)

		return
	}

	ok := orgModel.RemoveOrganization()

	if !ok {
		helper.Send(context, nil, robust.REMOVE_ORG_FAILURE)

		return
	}

	helper.Send(context, gin.H{
		"organizeName": organizeName,
	}, nil)
}

package doc_controller

import (
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllDocumentsParams struct {
	OrganizationID string ` json:"organizationID" form:"organizationID"`
}

func GetAllDocuments(ctx *gin.Context) {
	res := helper.Res{}

	var params GetAllDocumentsParams
	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Send(ctx, err)
		return
	}

	var docs []models.Document

	doc := models.Document{}
	if err := doc.FindAll(&docs); err != nil {
		res.Status(http.StatusForbidden).Send(ctx, err)
		return
	}

	res.Send(ctx, docs)
}

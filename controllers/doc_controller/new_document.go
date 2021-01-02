package doc_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewDocumentParams struct {
	Content        string `validate:"required" json:"content" form:"content"`
	Title          string `validate:"required" json:"title" form:"title"`
	Headings       string `validate:"required" json:"headings" form:"headings"`
	Cover          string `validate:"required" json:"cover" form:"cover"`
	Excerpt        string `validate:"required" json:"excerpt" form:"excerpt"`
	CategoryID     string `json:"categoryID" form:"categoryID"`
	OrganizationID string `validate:"required" form:"organizationID"`
	IsPublic       bool   `json:"isPublic" form:"isPublic"`
}

func NewDocument(ctx *gin.Context) {
	res := helper.Res{}
	claims, err := middlewares.GetClaims(ctx)

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	var params NewDocumentParams
	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	doc := models.Document{}
	utils.CpStruct(&params, &doc)

	doc.UserID = claims.ID

	if err := doc.New(); err != nil {
		res.Status(http.StatusForbidden).Error(err).Send(ctx)
		return
	}

	res.Success(doc).Send(ctx)
}

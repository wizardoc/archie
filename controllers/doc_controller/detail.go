package doc_controller

import (
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Detail(ctx *gin.Context) {
	res := helper.Res{}
	docId := ctx.Params.ByName("document_id")
	doc := models.Document{
		ID: docId,
	}

	if err := doc.Detail(); err != nil {
		res.Status(http.StatusBadRequest).Send(ctx, err)
		return
	}

	res.Send(ctx, doc)
}

package doc_controller

import (
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

const DOCUMENT_PATH = "public/doc/about-wizard.md"

func WizardIntroduction(ctx *gin.Context) {
	io := helper.ArchieIO{Path: DOCUMENT_PATH}
	data, err := io.ReadStringStream()
	res := helper.Res{}

	if err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, robust.CANNOT_FIND_FILE)
		return
	}

	res.Send(ctx, data)
}

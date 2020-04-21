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
	errRes := helper.Res{Status: http.StatusBadRequest}

	if err != nil {
		errRes.Err = robust.CANNOT_FIND_FILE
		errRes.Send(ctx)
		return
	}

	res := helper.Res{Data: data}
	res.Send(ctx)
}

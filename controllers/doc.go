package controllers

import (
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

const DOCUMENT_POSITION = "public/doc/about-wizard.md"

func WizardIntroduction(context *gin.Context) {
	file, err := os.Open(DOCUMENT_POSITION)
	utils.Check(err)

	content, err := ioutil.ReadAll(file)
	utils.Check(err)

	aboutContent := string(content)

	res := helper.Res{Data: aboutContent}
	res.Send(context)
}

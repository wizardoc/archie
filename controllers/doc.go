package controllers

import (
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func WizardIntroduction(context *gin.Context) {
	file, err := os.Open("public/doc/about-wizard.md")
	utils.Check(err)

	content, err := ioutil.ReadAll(file)
	utils.Check(err)

	aboutContent := string(content)

	helper.Send(context, aboutContent, nil)
}

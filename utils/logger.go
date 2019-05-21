package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LogInfo(infos ...string) {
	fmt.Fprintln(gin.DefaultWriter, infos)
}

func LogError(err interface{}) {
	fmt.Fprintln(gin.DefaultErrorWriter, err)
}

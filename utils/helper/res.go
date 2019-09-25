package helper

import (
	"archie/robust"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// response
func Send(context *gin.Context, data interface{}, err interface{}) {
	// valid
	if err != nil {
		_, ok := err.(robust.ArchieError)

		if !ok {
			panic("err must be a ArchieError or nil!")
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"data": data,
		"err":  err,
	})
}

func IsEmpty(target interface{}) bool {
	return !reflect.ValueOf(target).IsValid()
}

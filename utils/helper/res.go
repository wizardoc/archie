package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Res struct {
	Data   interface{}
	Err    error
	Status int
}

func (res Res) Send(ctx *gin.Context) {
	ctx.JSON(getStatus(res.Status), gin.H{
		"data": res.Data,
		"err":  res.Err,
	})
}

// parse status from arg
func getStatus(status int) int {
	if status == 0 {
		return http.StatusOK
	}

	return status
}

func IsEmpty(target interface{}) bool {
	return !reflect.ValueOf(target).IsValid()
}

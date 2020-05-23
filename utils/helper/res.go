package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Res struct {
	Data       interface{}
	Err        error
	StatusCode int
}

func (res *Res) spurt(ctx *gin.Context) {
	ctx.JSON(getStatus(res.StatusCode), gin.H{
		"data": res.Data,
		"err":  res.Err,
	})
}

func (res *Res) Error(ctx *gin.Context, err error) *Res {
	res.Err = err
	res.spurt(ctx)

	return res
}

func (res *Res) Send(ctx *gin.Context, data interface{}) *Res {
	res.Data = data
	res.spurt(ctx)

	return res
}

func (res *Res) Status(code int) *Res {
	res.StatusCode = code

	return res
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

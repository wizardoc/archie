package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Res struct {
	Data       interface{}
	Err        interface{}
	StatusCode int
}

func (res *Res) reset() {
	res.Data = nil
	res.Err = nil
	res.StatusCode = 0
}

func (res *Res) Error(err interface{}) *Res {
	res.Err = err

	return res
}

func (res *Res) Success(data interface{}) *Res {
	res.Data = data

	return res
}

func (res *Res) Status(code int) *Res {
	res.StatusCode = code

	return res
}

func (res *Res) Send(ctx *gin.Context) {
	ctx.JSON(getStatus(res.StatusCode), gin.H{
		"data": res.Data,
		"err":  res.Err,
	})
	res.reset()
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

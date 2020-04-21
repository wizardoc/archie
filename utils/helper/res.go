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

func GenAuthRes() Res {
	return GenRes(http.StatusUnauthorized)
}

func GenBadReqRes() Res {
	return GenRes(http.StatusBadRequest)
}

func GenServerErrRes() Res {
	return GenRes(http.StatusInternalServerError)
}

func GenSuccessRes() Res {
	return GenRes(http.StatusOK)
}

func GenRes(status int) Res {
	return Res{Status: status}
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

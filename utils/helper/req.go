package helper

import (
	"archie/robust"
	"github.com/gin-gonic/gin"
)

// @deprecated used to RESTful API
func BindWithValid(ctx *gin.Context, target interface{}) error {
	if err := ctx.Bind(target); err != nil {
		return robust.INVALID_PARAMS
	}

	validation := robust.Validation{Target: target}
	if err := validation.Valid(); err != nil {
		return err
	}

	return nil
}

// Used to GraphQL
func ValidParams(target interface{}) error {
	var err error
	validation := robust.Validation{Target: target}
	err = validation.Valid()

	if err != nil {
		err = robust.ArchieError{Code: 10001, Msg: err.Error()}
	}

	return err
}

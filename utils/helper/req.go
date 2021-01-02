package helper

import (
	"archie/robust"
	"github.com/gin-gonic/gin"
)

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
func BindParameter(target interface{}) error {
	return robust.Validation{Target: target}.Valid()
}

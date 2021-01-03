package resolver

import (
	"archie/constants"
	"archie/robust"
	"archie/utils"
	"archie/utils/jwt_utils"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type AuthParams struct {
}

func (r *Resolver) auth(ctx context.Context) (string, error) {
	ginCtx := ctx.Value(constants.GIN_CONTEXT).(*gin.Context)

	token, ok := getJWTFromHeader(ginCtx.Request)
	if !ok {
		return "", robust.JWT_DOES_NOT_EXIST
	}

	claims := jwt_utils.LoginClaims{}
	if err := parseToken2Claims(token, &claims); err != nil {
		return "", err
	}

	return token, nil
}

func getJWTFromHeader(req *http.Request) (jwtString string, ok bool) {
	headers := req.Header
	auth := headers["Authentication"]

	return auth[0], len(auth) == 0
}

func parseToken2Claims(token string, targetClaims interface{}) error {
	JWTClaims := parseToken(token)

	if err := JWTClaims.Valid(); err != nil {
		return err
	}

	if err := mapstructure.Decode(JWTClaims, &targetClaims); err != nil {
		return err
	}

	return nil
}

func parseToken(tokenStr string) jwt.Claims {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwt_utils.GetSecretKey(), nil
	})

	utils.Check(err)

	return token.Claims
}

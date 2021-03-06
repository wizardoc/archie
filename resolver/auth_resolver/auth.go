package auth_resolver

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

type AuthResolver struct {
}

func (r *AuthResolver) Auth(ctx context.Context) (*jwt_utils.LoginClaims, error) {
	ginCtx := ctx.Value(constants.GIN_CONTEXT).(*gin.Context)
	token, err := getJWTFromHeader(ginCtx.Request)

	if err != nil {
		return nil, err
	}

	claims := jwt_utils.LoginClaims{}
	if err := parseToken2Claims(token, &claims); err != nil {
		return nil, err
	}

	return &claims, nil
}

func getJWTFromHeader(req *http.Request) (string, error) {
	auth := req.Header["Authentication"]

	if len(auth) == 0 {
		return "", robust.JWT_DOES_NOT_EXIST
	}

	return auth[0], nil
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

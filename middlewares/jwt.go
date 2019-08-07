package middlewares

import (
	"archie/robust"
	"archie/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func getJWTFromHeader(req *http.Request) (jwtString string, ok bool) {
	headers := req.Header
	auth := headers["Authentication"]

	if len(auth) == 0 {
		return "", false
	}

	return auth[0], true
}

func ParseToken(tokenStr string) jwt.Claims {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return utils.GetSecretKey(), nil
	})

	utils.Check(err)

	return token.Claims
}

func ValidateToken(context *gin.Context) {
	jwtString, ok := getJWTFromHeader(context.Request)

	if !ok {
		context.JSON(http.StatusOK, gin.H{
			"data": nil,
			"err":  robust.JWT_DOES_NOT_EXIST,
		})

		context.Abort()

		return
	}

	JWTClaims := ParseToken(jwtString)
	err := JWTClaims.Valid()

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"data": nil,
			"err":  err,
		})
	}

	claims := utils.Claims{}

	mapstructure.Decode(JWTClaims, &claims)

	context.Set("claims", claims)
	context.Next()
}

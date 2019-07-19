package middlewares

import (
	"archie/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ParseToken(tokenStr string) jwt.Claims {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return utils.GetSecretKey(), nil
	})

	utils.Check(err)

	return token.Claims
}

func validateToken(context *gin.Context) {

}

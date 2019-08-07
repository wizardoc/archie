package middlewares

import (
	"archie/connection"
	"archie/robust"
	"archie/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
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

func AddInBlackSet(userId string) (err error) {
	connection.GetRedisConnMust(func(conn redis.Conn) {
		_, err = conn.Do("SADD", "black_set", userId)
	})

	return
}

func IsExistInBlackSet(userId string) (isExist bool) {
	connection.GetRedisConnMust(func(conn redis.Conn) {
		var err error

		isExist, err = redis.Bool(conn.Do("SISMEMBER", "black_set", userId))
		utils.Check(err)
	})

	return
}

func ValidateToken(context *gin.Context) {
	jwtString, ok := getJWTFromHeader(context.Request)

	if !ok {
		utils.Send(context, nil, robust.JWT_DOES_NOT_EXIST)
		context.Abort()

		return
	}

	JWTClaims := ParseToken(jwtString)
	err := JWTClaims.Valid()

	if err != nil {
		utils.Send(context, nil, err)
	}

	claims := utils.Claims{}
	mapstructure.Decode(JWTClaims, &claims)

	if IsExistInBlackSet(claims.UserId) {
		utils.Send(context, nil, robust.JWT_NOT_ALLOWED)

		return
	}

	context.Set("claims", claims)
	context.Next()
}

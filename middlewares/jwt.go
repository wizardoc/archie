package middlewares

import (
	"archie/connection"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func ValidateToken(context *gin.Context) {
	jwtString, ok := getJWTFromHeader(context.Request)
	authErrRes := helper.Res{Status: http.StatusBadRequest}
	unAuthErrRes := helper.Res{Status: http.StatusUnauthorized}
	serverErrRes := helper.Res{Status: http.StatusInternalServerError}

	/** JWT 不存在 */
	if !ok {
		authErrRes.Err = robust.JWT_DOES_NOT_EXIST
		authErrRes.Send(context)
		context.Abort()

		return
	}

	JWTClaims := ParseToken(jwtString)
	err := JWTClaims.Valid()

	/** 不合法的 Token */
	if err != nil {
		authErrRes.Err = err
		authErrRes.Send(context)
		return
	}

	claims := utils.Claims{}
	err = mapstructure.Decode(JWTClaims, &claims)

	/** 解析 Claims 失败 */
	if err != nil {
		serverErrRes.Err = robust.JWT_CANNOT_PARSE_CLAIMS
		serverErrRes.Send(context)
		return
	}

	/** 在小黑屋，JWT 不被允许 */
	if IsExistInBlackSet(claims.UserId) {
		unAuthErrRes.Err = robust.JWT_NOT_ALLOWED
		unAuthErrRes.Send(context)
		return
	}

	context.Set("claims", claims)
	context.Next()
}

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

/** 验证获取 Token */
func GetClaims(context *gin.Context) (utils.Claims, error) {
	claims, isExist := context.Get("claims")

	if !isExist {
		return utils.Claims{}, robust.JWT_DOES_NOT_EXIST
	}

	parsedClaims, ok := claims.(utils.Claims)

	if !ok {
		return utils.Claims{}, robust.JWT_PARSE_ERROR
	}

	return parsedClaims, nil
}

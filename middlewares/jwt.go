package middlewares

import (
	"archie/connection/redis_conn"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

// JWT 验证中间件，用于校验 token，将 claims 转发到下一个中间件
func ValidateToken(context *gin.Context) {
	jwtString, ok := getJWTFromHeader(context.Request)
	authErrRes := helper.Res{Status: http.StatusBadRequest}

	/** JWT 不存在 */
	if !ok {
		authErrRes.Err = robust.JWT_DOES_NOT_EXIST
		authErrRes.Send(context)
		context.Abort()

		return
	}

	claims, err := ParseToken2Claims(jwtString)

	// parse jwt fail
	if err != nil {
		fmt.Println(err)
		authErrRes.Err = robust.JWT_NOT_ALLOWED
		authErrRes.Send(context)
		return
	}

	/** 在小黑屋，JWT 不被允许 */
	//if IsExistInBlackSet(claims.UserId) {
	//	unAuthErrRes.Err = robust.JWT_NOT_ALLOWED
	//	unAuthErrRes.Send(context)
	//	return
	//}

	context.Set("claims", *claims)
	context.Next()
}

func ParseToken2Claims(token string) (*utils.Claims, error) {
	JWTClaims := ParseToken(token)

	/** 不合法的 Token */
	if err := JWTClaims.Valid(); err != nil {
		return nil, err
	}

	claims := utils.Claims{}
	if err := mapstructure.Decode(JWTClaims, &claims); err != nil {
		return nil, err
	}

	return &claims, nil
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
	redis_conn.GetRedisConnMust(func(conn redis.Conn) {
		_, err = conn.Do("SADD", "black_set", userId)
	})

	return
}

func IsExistInBlackSet(userId string) (isExist bool) {
	redis_conn.GetRedisConnMust(func(conn redis.Conn) {
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

	return ParseClaims(claims)
}

func ParseClaims(claims interface{}) (utils.Claims, error) {
	parsedClaims, ok := claims.(utils.Claims)

	if !ok {
		return utils.Claims{}, robust.JWT_PARSE_ERROR
	}

	return parsedClaims, nil
}

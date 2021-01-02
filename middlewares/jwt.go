package middlewares

import (
	"archie/connection/redis_conn"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"archie/utils/jwt_utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

// JWT 验证中间件，用于校验 token，将 claims 转发到下一个中间件
func ValidateToken(ctx *gin.Context) {
	jwtString, ok := GetJWTFromHeader(ctx.Request)
	res := helper.Res{}

	/** JWT 不存在 */
	if !ok {
		res.Status(http.StatusUnauthorized).Error(robust.JWT_DOES_NOT_EXIST).Send(ctx)
		ctx.Abort()

		return
	}

	claims := jwt_utils.LoginClaims{}

	// parse jwt fail
	if err := ParseToken2Claims(jwtString, &claims); err != nil {
		res.Status(http.StatusUnauthorized).Error(robust.JWT_NOT_ALLOWED).Send(ctx)
		ctx.Abort()

		return
	}

	/** 在小黑屋，JWT 不被允许 */
	//if IsExistInBlackSet(claims.UserID) {
	//	unAuthErrRes.Err = robust.JWT_NOT_ALLOWED
	//	unAuthErrRes.Send(ctx)
	//	return
	//}
	ctx.Set("claims", claims)
	ctx.Next()
}

func ParseToken2Claims(token string, targetClaims interface{}) error {
	JWTClaims := ParseToken(token)

	/** 不合法的 Token */
	if err := JWTClaims.Valid(); err != nil {
		return err
	}

	if err := mapstructure.Decode(JWTClaims, &targetClaims); err != nil {
		return err
	}

	return nil
}

func GetJWTFromHeader(req *http.Request) (jwtString string, ok bool) {
	headers := req.Header
	auth := headers["Authentication"]

	fmt.Println(headers["authentication"])

	if len(auth) == 0 {
		return "", false
	}

	return auth[0], true
}

func ParseToken(tokenStr string) jwt.Claims {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwt_utils.GetSecretKey(), nil
	})

	utils.Check(err)

	return token.Claims
}

func AddInBlackSet(userId string) (err error) {
	redis_conn.GetRedisConnMust(func(conn redis.Conn) error {
		_, err = conn.Do("SADD", "black_set", userId)

		return err
	})

	return
}

func IsExistInBlackSet(userId string) (isExist bool) {
	redis_conn.GetRedisConnMust(func(conn redis.Conn) error {
		var err error

		isExist, err = redis.Bool(conn.Do("SISMEMBER", "black_set", userId))

		return err
	})

	return
}

/** 验证获取 Token */
func GetClaims(ctx *gin.Context) (jwt_utils.LoginClaims, error) {
	claims, isExist := ctx.Get("claims")

	if !isExist {
		return jwt_utils.LoginClaims{}, robust.JWT_DOES_NOT_EXIST
	}

	return ParseClaims(claims)
}

func ParseClaims(claims interface{}) (jwt_utils.LoginClaims, error) {
	parsedClaims, ok := claims.(jwt_utils.LoginClaims)

	if !ok {
		return jwt_utils.LoginClaims{}, robust.JWT_PARSE_ERROR
	}

	return parsedClaims, nil
}

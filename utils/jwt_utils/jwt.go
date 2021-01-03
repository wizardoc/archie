package jwt_utils

import (
	"archie/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

const (
	SecretKey = "secret key"
)

type IClaims interface {
	SignJWT(duration int) string
}

type Claims struct {
	ISS string // Issue
	JTI string // JWT ID
	IAT int32  // issued at
	EXP int32  // expiration
}

func (claims Claims) SignJWT(duration int, c IClaims) string {
	jwtMap := make(map[string]interface{})
	token := jwt.New(jwt.SigningMethodHS256)

	claims.ISS = "younccat"
	claims.JTI = uuid.NewV4().String()
	claims.EXP = utils.ParseToMillisecond(time.Now().Add(time.Hour * time.Duration(duration)).UnixNano())
	claims.IAT = utils.Now()

	if err := mapstructure.Decode(c, &jwtMap); err != nil {
		log.Println(err)
	}

	fmt.Println(jwtMap)

	token.Claims = jwt.MapClaims(jwtMap)

	tokenStr, err := token.SignedString(GetSecretKey())

	utils.Check(err)

	return tokenStr
}

func GetSecretKey() []byte {
	return []byte(SecretKey)
}

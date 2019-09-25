package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"github.com/satori/go.uuid"
	"time"
)

const (
	SecretKey = "secret key"
)

type Claims struct {
	ISS    string // Issue
	JTI    string // JWT ID
	IAT    int64  // issued at
	EXP    int64  // expiration
	UserId string
}

func (claims Claims) SignJWT() string {
	jwtMap := make(map[string]interface{})
	token := jwt.New(jwt.SigningMethodHS256)

	claims.ISS = "younccat"
	claims.JTI = uuid.NewV4().String()
	claims.EXP = ParseToMillisecond(time.Now().Add(time.Hour * time.Duration(24)).UnixNano())
	claims.IAT = Now()

	mapstructure.Decode(claims, &jwtMap)

	token.Claims = jwt.MapClaims(jwtMap)

	tokenStr, err := token.SignedString(GetSecretKey())

	Check(err)

	return tokenStr
}

func GetSecretKey() []byte {
	return []byte(SecretKey)
}

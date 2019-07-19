package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
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

	mapstructure.Decode(claims, &jwtMap)

	token.Claims = jwt.MapClaims(jwtMap)

	tokenStr, err := token.SignedString(GetSecretKey())

	Check(err)

	return tokenStr
}

func GetSecretKey() []byte {
	return []byte(SecretKey)
}

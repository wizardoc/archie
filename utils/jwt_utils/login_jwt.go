package jwt_utils

import "archie/models"

type LoginClaims struct {
	Claims      `mapstructure:"squash"`
	models.User `mapstructure:"squash"`
}

func (claims *LoginClaims) SignJWT(duration int) string {
	return claims.Claims.SignJWT(duration, claims)
}

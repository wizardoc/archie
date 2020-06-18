package jwt_utils

type InviteClaims struct {
	Claims
	InviteUser   string
	OrganizeName string
}

func (claims *InviteClaims) SignJWT(duration int) string {
	return claims.Claims.SignJWT(duration, claims)
}

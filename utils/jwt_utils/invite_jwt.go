package jwt_utils

type InviteClaims struct {
	Claims
	UserID string
	OrgID  string
	Role   int
}

func (claims *InviteClaims) SignJWT(duration int) string {
	return claims.Claims.SignJWT(duration, claims)
}

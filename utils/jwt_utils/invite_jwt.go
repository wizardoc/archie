package jwt_utils

type InviteClaims struct {
	Claims
	InviteUserID string
	OrgID        string
	Role         int
	UserID       string
}

func (claims *InviteClaims) SignJWT(duration int) string {
	return claims.Claims.SignJWT(duration, claims)
}

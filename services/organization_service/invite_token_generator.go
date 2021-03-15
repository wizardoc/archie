package organization_service

import (
	"archie/utils/jwt_utils"
)

const EXPIRE_TOKEN_TIME_HOUR = 2

func InviteTokenGenerator(
	userID string,
	orgID string,
	inviteUserID string,
	role int,
) string {
	claims := jwt_utils.InviteClaims{
		UserID:       userID,
		OrgID:        orgID,
		Role:         role,
		InviteUserID: inviteUserID,
	}
	token := claims.SignJWT(EXPIRE_TOKEN_TIME_HOUR)

	return token
}

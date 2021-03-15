package organization_service

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/jwt_utils"
)

// Parse token when user visit the URL that generate by JWT
// check the token whether is valid or not, and append the user into
// the organization
func InviteUser(token string, loginClaims jwt_utils.LoginClaims) (string, error) {
	inviteClaims := jwt_utils.InviteClaims{}
	if err := middlewares.ParseToken2Claims(token, &inviteClaims); err != nil {
		return "", err
	}

	// Verify that the user who logged is same person as the invited user
	if inviteClaims.InviteUserID != loginClaims.ID {
		return "", robust.ORGANIZATION_INVITE_AUTH
	}

	user := models.User{
		ID: inviteClaims.InviteUserID,
	}
	org := models.Organization{
		ID: inviteClaims.OrgID,
	}

	member := models.Member{
		UserID:         user.ID,
		OrganizationID: org.ID,
		Role:           inviteClaims.Role,
	}

	if err := member.Create(); err != nil {
		return "", err
	}

	return inviteClaims.OrgID, nil
}

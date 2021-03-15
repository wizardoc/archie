package organization_resolver

import (
	"archie/constants/organization_rbac"
	"archie/robust"
	"archie/services/organization_service"
	"archie/utils/jwt_utils"
	"context"
)

type InviteTokenParams struct {
	UserID  string
	OrgID   string
	Role    float64
	OrgName string
}

func (r *OrganizationResolver) InviteToken(ctx context.Context, params InviteTokenParams) (string, error) {
	var claims jwt_utils.LoginClaims
	if err := r.Auth(ctx, &claims); err != nil {
		return "", err
	}

	// Verify the user permission
	isVerify, err := organization_service.VerifyUserPermission(claims.ID, params.OrgID, organization_rbac.ORG_INVITE)
	if err != nil {
		return "", err
	}

	// The user does have sufficient permission
	if !isVerify {
		return "", robust.INVALID_PERMISSION
	}

	return organization_service.InviteTokenGenerator(claims.ID, params.OrgID, params.UserID, int(params.Role)), nil
}

package organization_resolver

import (
	"archie/services/organization_service"
	"context"
)

type InviteTokenParams struct {
	UserID string
	OrgID  string
	Role   float64
}

func (r *OrganizationResolver) InviteToken(ctx context.Context, params InviteTokenParams) string {
	return organization_service.InviteTokenGenerator(params.UserID, params.OrgID, int(params.Role))
}

package organization_service

import (
	"archie/constants/organization_rbac"
	"archie/models"
	"archie/robust"
	"archie/utils"
)

func VerifyUserPermission(userID string, orgID string, specifyPermission int) (bool, error) {
	m := models.Member{
		OrganizationID: orgID,
		UserID:         userID,
	}
	if err := m.Query(); err != nil {
		return false, err
	}

	// The user have not yet join the rog
	if m.JoinTime == "" {
		return false, robust.ORGANIZATION_INVITE_EXIST
	}

	role := organization_rbac.Role{Name: m.Role}
	permissions := role.GetPermission()

	return utils.ArrayIncludes(permissions, specifyPermission), nil
}

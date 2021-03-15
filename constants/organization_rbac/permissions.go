package organization_rbac

const (
	// document
	DOCUMENT_WRITE = iota
	DOCUMENT_READ
	DOCUMENT_CREATE

	// organization
	ORG_INVITE
	ORG_EDIT
	ORG_DELETE
)

var roleMap = map[int][]int{
	OWNER:  {DOCUMENT_WRITE, DOCUMENT_READ, DOCUMENT_CREATE, ORG_EDIT, ORG_INVITE, ORG_DELETE},
	ADMIN:  {DOCUMENT_READ, DOCUMENT_CREATE, DOCUMENT_WRITE, ORG_EDIT, ORG_INVITE},
	EDITOR: {DOCUMENT_READ, DOCUMENT_WRITE},
	MEMBER: {DOCUMENT_READ},
}

type Role struct {
	Name int
}

func (r *Role) GetPermission() []int {
	return roleMap[r.Name]
}

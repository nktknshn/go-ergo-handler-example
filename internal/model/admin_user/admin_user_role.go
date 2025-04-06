package admin_user

type AdminUserRole string

const (
	AdminUserRoleAdmin     AdminUserRole = "admin"
	AdminUserRoleModerator AdminUserRole = "moderator"
	AdminUserRolePublisher AdminUserRole = "publisher"
)

func (r AdminUserRole) IsValid() bool {
	switch r {
	case AdminUserRoleAdmin, AdminUserRoleModerator, AdminUserRolePublisher:
		return true
	default:
		return false
	}
}

func (r AdminUserRole) String() string {
	return string(r)
}

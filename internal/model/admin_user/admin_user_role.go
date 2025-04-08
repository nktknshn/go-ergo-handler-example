package admin_user

import "fmt"

type AdminUserRole struct {
	role string
}

var (
	AdminUserRoleUnknown   = AdminUserRole{}
	AdminUserRoleAdmin     = AdminUserRole{role: "admin"}
	AdminUserRoleModerator = AdminUserRole{role: "moderator"}
	AdminUserRolePublisher = AdminUserRole{role: "publisher"}
)

func AdminRoleFromString(role string) (AdminUserRole, error) {
	switch role {
	case "admin":
		return AdminUserRoleAdmin, nil
	case "moderator":
		return AdminUserRoleModerator, nil
	case "publisher":
		return AdminUserRolePublisher, nil
	default:
		return AdminUserRoleUnknown, fmt.Errorf("invalid admin role: %s", role)
	}
}

func (r AdminUserRole) IsValid() bool {
	switch r {
	case AdminUserRoleAdmin, AdminUserRoleModerator, AdminUserRolePublisher:
		return true
	default:
		return false
	}
}

func (r AdminUserRole) String() string {
	return r.role
}

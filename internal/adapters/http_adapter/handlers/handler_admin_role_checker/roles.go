package handler_admin_role_checker

import "github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"

var (
	AdminAndPublisher = NewAdminRoleChecker(
		CheckSlice([]admin_user.AdminUserRole{admin_user.AdminUserRoleAdmin, admin_user.AdminUserRolePublisher}),
	)
	Admin = NewAdminRoleChecker(
		CheckSlice([]admin_user.AdminUserRole{admin_user.AdminUserRoleAdmin}),
	)
)

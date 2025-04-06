package admin_user

type AdminUserID int

func (id AdminUserID) Int() int {
	return int(id)
}

type AdminUserUsername string

func (u AdminUserUsername) String() string {
	return string(u)
}

type AdminUser struct {
	ID       AdminUserID
	Username AdminUserUsername
	Role     AdminUserRole
}

func (a *AdminUser) HasID() bool {
	return a.ID != 0
}

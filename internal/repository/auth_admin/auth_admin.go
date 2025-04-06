package auth_admin

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
	"github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/auth_admin"
)

type AuthAdminRepository struct {
	adminByToken map[string]admin_user.AdminUserID
}

func NewAuthAdminRepository() *AuthAdminRepository {
	return &AuthAdminRepository{
		adminByToken: make(map[string]admin_user.AdminUserID),
	}
}

func (r *AuthAdminRepository) GetAdminID(ctx context.Context, token string) (admin_user.AdminUserID, error) {
	adminID, ok := r.adminByToken[token]
	if !ok {
		return 0, auth_admin.ErrTokenInvalid
	}
	return adminID, nil
}

func (r *AuthAdminRepository) SetAdminID(ctx context.Context, token string, adminID admin_user.AdminUserID) error {
	r.adminByToken[token] = adminID
	return nil
}

func (r *AuthAdminRepository) DeleteAdminID(ctx context.Context, token string) error {
	delete(r.adminByToken, token)
	return nil
}

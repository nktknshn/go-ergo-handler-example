package auth_admin

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
)

type AuthAdminUseCase struct {
	authAdminRepository authAdminRepository
	adminUserRepository adminUserRepository
}

type authAdminRepository interface {
	GetAdminID(ctx context.Context, token string) (admin_user.AdminUserID, error)
}

type adminUserRepository interface {
	GetAdminByID(ctx context.Context, adminID admin_user.AdminUserID) (admin_user.AdminUser, error)
}

func NewAuthAdminUseCase(authAdminRepository authAdminRepository, adminUserRepository adminUserRepository) *AuthAdminUseCase {
	return &AuthAdminUseCase{authAdminRepository, adminUserRepository}
}

func (u *AuthAdminUseCase) GetUser(ctx context.Context, token string) (*admin_user.AdminUser, bool, error) {
	adminID, err := u.authAdminRepository.GetAdminID(ctx, token)
	if err != nil {
		return nil, false, err
	}
	admin, err := u.adminUserRepository.GetAdminByID(ctx, adminID)
	if err != nil {
		return nil, false, err
	}
	return &admin, true, nil
}

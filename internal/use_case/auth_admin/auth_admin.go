package auth_admin

import (
	"context"
	"errors"
	"log/slog"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
	authAdminRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/auth_admin"
	authAdminUseCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/auth_admin"
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

func (u *AuthAdminUseCase) ValidateToken(ctx context.Context, token string) (*admin_user.AdminUser, bool, error) {
	adminID, err := u.authAdminRepository.GetAdminID(ctx, token)
	if errors.Is(err, authAdminRepoValObj.ErrTokenInvalid) {
		return nil, false, authAdminUseCaseValObj.ErrTokenInvalid
	}
	if errors.Is(err, authAdminRepoValObj.ErrTokenNotFound) {
		return nil, false, authAdminUseCaseValObj.ErrTokenInvalid
	}
	if err != nil {
		slog.Error("u.authAdminRepository.GetAdminID", "error", err)
		return nil, false, authAdminUseCaseValObj.ErrGetAdminIDFailed
	}
	admin, err := u.adminUserRepository.GetAdminByID(ctx, adminID)
	if err != nil {
		slog.Error("u.adminUserRepository.GetAdminByID", "error", err)
		return nil, false, authAdminUseCaseValObj.ErrGetAdminFailed
	}
	return &admin, true, nil
}

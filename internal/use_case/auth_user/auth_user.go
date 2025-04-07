package auth_user

import (
	"context"
	"log/slog"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	authUserUseCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/auth_user"
)

type AuthUserUseCase struct {
	authUserRepository authUserRepository
	userRepository     userRepository
}

type authUserRepository interface {
	GetUserID(ctx context.Context, token string) (user.UserID, error)
}

type userRepository interface {
	GetUserByID(ctx context.Context, userID user.UserID) (user.User, error)
}

func NewAuthUserUseCase(authUserRepository authUserRepository, userRepository userRepository) *AuthUserUseCase {
	return &AuthUserUseCase{authUserRepository, userRepository}
}

func (u *AuthUserUseCase) ValidateToken(ctx context.Context, token string) (*user.User, bool, error) {
	userID, err := u.authUserRepository.GetUserID(ctx, token)
	if err != nil {
		slog.Error("u.authUserRepository.GetUserID", "error", err)
		return nil, false, authUserUseCaseValObj.ErrGetUserIDFailed
	}
	user, err := u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		slog.Error("u.userRepository.GetUserByID", "error", err)
		return nil, false, authUserUseCaseValObj.ErrGetUserFailed
	}
	return &user, true, nil
}

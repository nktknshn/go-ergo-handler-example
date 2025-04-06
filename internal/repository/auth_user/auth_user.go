package auth_user

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	"github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/auth_user"
)

type AuthUserRepository struct {
	userByToken map[string]user.UserID
}

func NewAuthUserRepository() *AuthUserRepository {
	return &AuthUserRepository{
		userByToken: make(map[string]user.UserID),
	}
}

func (r *AuthUserRepository) GetUserID(ctx context.Context, token string) (user.UserID, error) {
	userID, ok := r.userByToken[token]
	if !ok {
		return 0, auth_user.ErrTokenInvalid
	}
	return userID, nil
}

func (r *AuthUserRepository) SetUserID(ctx context.Context, token string, userID user.UserID) error {
	r.userByToken[token] = userID
	return nil
}

func (r *AuthUserRepository) DeleteUserID(ctx context.Context, token string) error {
	delete(r.userByToken, token)
	return nil
}

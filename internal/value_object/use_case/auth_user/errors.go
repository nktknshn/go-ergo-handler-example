package auth_user

import "errors"

var (
	ErrTokenInvalid    = errors.New("token invalid")
	ErrGetUserIDFailed = errors.New("failed to get user id")
	ErrGetUserFailed   = errors.New("failed to get user")
)

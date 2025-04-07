package auth_admin

import "errors"

var (
	ErrTokenInvalid     = errors.New("token invalid")
	ErrGetAdminIDFailed = errors.New("failed to get admin id")
	ErrGetAdminFailed   = errors.New("failed to get admin")
)

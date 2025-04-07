package books

import "errors"

var (
	ErrCreateBookFailed  = errors.New("failed to create book")
	ErrGetBookListFailed = errors.New("failed to get book list")
	ErrGetBookFailed     = errors.New("failed to get book")
	ErrBookNotFound      = errors.New("book not found")
	ErrInvalidCursor     = errors.New("invalid cursor")
)

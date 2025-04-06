package books

import "errors"

var (
	ErrBookNotFound  = errors.New("book not found")
	ErrInvalidCursor = errors.New("invalid cursor")
)

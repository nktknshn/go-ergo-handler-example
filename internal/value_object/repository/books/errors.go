package books

import "errors"

var (
	ErrBookNotFound          = errors.New("book not found")
	ErrBookListCursorInvalid = errors.New("book list cursor in invalid")
)

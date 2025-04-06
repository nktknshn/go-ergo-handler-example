package book_favorites

import "errors"

var (
	ErrBookNotFound          = errors.New("book not found")
	ErrBookAlreadyInFavorite = errors.New("book already in favorite")
)

package book_favorites

import "errors"

var (
	ErrBookAlreadyInFavorites  = errors.New("book already in favorites")
	ErrBookNotFoundInFavorites = errors.New("book not found in favorites")
)

package book_favorites

import "errors"

var (
	ErrAddBookToFavoritesFailed = errors.New("failed to add book to favorites")
	ErrGetBookFavoritesFailed   = errors.New("failed to get book favorites")
	ErrBookNotFound             = errors.New("book not found")
	ErrBookAlreadyInFavorite    = errors.New("book already in favorite")
)

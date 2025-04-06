package books_favorite

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book_favorite"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
)

type BooksFavoriteUseCase struct {
	bookFavoriteRepository bookFavoriteRepository
}

type bookFavoriteRepository interface {
	AddFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) (book_favorite.BookFavorite, error)
	RemoveFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) error
	GetFavoriteBooks(ctx context.Context, userID user.UserID) ([]book.BookID, error)
}

func NewBooksFavoriteUseCase(bookFavoriteRepository bookFavoriteRepository) *BooksFavoriteUseCase {
	return &BooksFavoriteUseCase{bookFavoriteRepository: bookFavoriteRepository}
}

func (u *BooksFavoriteUseCase) AddFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) (book_favorite.BookFavorite, error) {
	return u.bookFavoriteRepository.AddFavoriteBook(ctx, userID, bookID)
}

func (u *BooksFavoriteUseCase) RemoveFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) error {
	return u.bookFavoriteRepository.RemoveFavoriteBook(ctx, userID, bookID)
}

func (u *BooksFavoriteUseCase) GetFavoriteBooks(ctx context.Context, userID user.UserID) ([]book.BookID, error) {
	return u.bookFavoriteRepository.GetFavoriteBooks(ctx, userID)
}

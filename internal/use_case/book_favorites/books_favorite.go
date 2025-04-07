package book_favorites

import (
	"context"
	"errors"
	"log/slog"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book_favorite"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	bookFavoriteRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/book_favorites"
	bookRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books"
	booksFavoriteUseCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/book_favorites"
)

type BookFavoritesUseCase struct {
	bookFavoriteRepository bookFavoriteRepository
	bookRepository         bookRepository
}

type bookFavoriteRepository interface {
	AddFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) (book_favorite.BookFavorite, error)
	RemoveFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) error
	GetFavoriteBooks(ctx context.Context, userID user.UserID) ([]book.BookID, error)
}

type bookRepository interface {
	GetBookByID(ctx context.Context, bookID book.BookID) (book.Book, error)
}

func NewBooksFavoriteUseCase(bookFavoriteRepository bookFavoriteRepository, bookRepository bookRepository) *BookFavoritesUseCase {
	return &BookFavoritesUseCase{
		bookFavoriteRepository,
		bookRepository,
	}
}

func (u *BookFavoritesUseCase) AddFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) (book_favorite.BookFavorite, error) {

	_, err := u.bookRepository.GetBookByID(ctx, bookID)

	if errors.Is(err, bookRepoValObj.ErrBookNotFound) {
		return book_favorite.BookFavorite{}, booksFavoriteUseCaseValObj.ErrBookNotFound
	}

	if err != nil {
		slog.Error("u.bookRepository.GetBookByID", "error", err)
		return book_favorite.BookFavorite{}, booksFavoriteUseCaseValObj.ErrGetBookFavoritesFailed
	}

	favorite, err := u.bookFavoriteRepository.AddFavoriteBook(ctx, userID, bookID)

	if errors.Is(err, bookFavoriteRepoValObj.ErrBookAlreadyInFavorites) {
		return book_favorite.BookFavorite{}, booksFavoriteUseCaseValObj.ErrBookAlreadyInFavorite
	}

	if err != nil {
		slog.Error("u.bookFavoriteRepository.AddFavoriteBook", "error", err)
		return book_favorite.BookFavorite{}, booksFavoriteUseCaseValObj.ErrAddBookToFavoritesFailed
	}

	return favorite, nil
}

func (u *BookFavoritesUseCase) RemoveFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) error {
	return u.bookFavoriteRepository.RemoveFavoriteBook(ctx, userID, bookID)
}

func (u *BookFavoritesUseCase) GetFavoriteBooks(ctx context.Context, userID user.UserID) ([]book.BookID, error) {
	return u.bookFavoriteRepository.GetFavoriteBooks(ctx, userID)
}

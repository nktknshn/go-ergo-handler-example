package books_favorite

import (
	"context"
	"errors"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book_favorite"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	bookRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books"
	bookFavoriteRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books_favorite"
	booksFavoriteUseCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books_favorite"
)

type BooksFavoriteUseCase struct {
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

func NewBooksFavoriteUseCase(bookFavoriteRepository bookFavoriteRepository, bookRepository bookRepository) *BooksFavoriteUseCase {
	return &BooksFavoriteUseCase{
		bookFavoriteRepository,
		bookRepository,
	}
}

func (u *BooksFavoriteUseCase) AddFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) (book_favorite.BookFavorite, error) {

	_, err := u.bookRepository.GetBookByID(ctx, bookID)

	if errors.Is(err, bookRepoValObj.ErrBookNotFound) {
		return book_favorite.BookFavorite{}, booksFavoriteUseCaseValObj.ErrBookNotFound
	}

	if err != nil {
		return book_favorite.BookFavorite{}, err
	}

	_, err = u.bookFavoriteRepository.AddFavoriteBook(ctx, userID, bookID)

	if errors.Is(err, bookFavoriteRepoValObj.ErrBookAlreadyInFavorites) {
		return book_favorite.BookFavorite{}, booksFavoriteUseCaseValObj.ErrBookAlreadyInFavorite
	}

	return u.bookFavoriteRepository.AddFavoriteBook(ctx, userID, bookID)
}

func (u *BooksFavoriteUseCase) RemoveFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) error {
	return u.bookFavoriteRepository.RemoveFavoriteBook(ctx, userID, bookID)
}

func (u *BooksFavoriteUseCase) GetFavoriteBooks(ctx context.Context, userID user.UserID) ([]book.BookID, error) {
	return u.bookFavoriteRepository.GetFavoriteBooks(ctx, userID)
}

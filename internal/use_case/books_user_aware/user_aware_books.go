package books_user_aware

import (
	"context"
	"log/slog"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"

	useCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books"
	awaredUseCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books_user_aware"
)

type BooksUserAwareUseCase struct {
	booksUseCase            booksUseCase
	favoriteBooksRepository favoriteBooksRepository
}

type booksUseCase interface {
	GetBooksList(ctx context.Context, query useCaseValObj.GetBooksListQuery) (useCaseValObj.GetBooksListResponse, error)
}

type favoriteBooksRepository interface {
	GetFavoriteBooks(ctx context.Context, userID user.UserID) ([]book.BookID, error)
}

func NewBooksUserAwareUseCase(booksUseCase booksUseCase, favoriteBooksRepository favoriteBooksRepository) *BooksUserAwareUseCase {
	return &BooksUserAwareUseCase{
		booksUseCase:            booksUseCase,
		favoriteBooksRepository: favoriteBooksRepository,
	}
}

func (u *BooksUserAwareUseCase) GetBooksList(ctx context.Context, userID user.UserID, query useCaseValObj.GetBooksListQuery) (awaredUseCaseValObj.GetBooksListResponse, error) {
	response := awaredUseCaseValObj.GetBooksListResponse{}
	resp, err := u.booksUseCase.GetBooksList(ctx, query)
	if err != nil {
		slog.Error("u.booksUseCase.GetBooksList", "error", err)
		return response, useCaseValObj.ErrGetBookListFailed
	}
	response.Books = resp.Books
	favoriteBooks, err := u.favoriteBooksRepository.GetFavoriteBooks(ctx, userID)
	if err != nil {
		slog.Error("u.favoriteBooksRepository.GetFavoriteBooks", "error", err)
		return response, useCaseValObj.ErrGetBookListFailed
	}
	response.FavoriteBooks = make([]book.BookID, 0)
	if len(favoriteBooks) > 0 {
		response.FavoriteBooks = favoriteBooks
	}
	response.Cursor = resp.Cursor
	response.HasMore = resp.HasMore
	return response, nil
}

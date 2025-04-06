package use_cases

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/app/repositories"
	"github.com/nktknshn/go-ergo-handler-example/internal/use_case/auth_admin"
	"github.com/nktknshn/go-ergo-handler-example/internal/use_case/auth_user"
	"github.com/nktknshn/go-ergo-handler-example/internal/use_case/book_favorites"
	"github.com/nktknshn/go-ergo-handler-example/internal/use_case/books"
	"github.com/nktknshn/go-ergo-handler-example/internal/use_case/books_user_aware"
)

type UseCases struct {
	booksUseCase          *books.BooksUseCase
	booksUserAwareUseCase *books_user_aware.BooksUserAwareUseCase
	bookFavoriteUseCase   *book_favorites.BookFavoritesUseCase
	authUserUseCase       *auth_user.AuthUserUseCase
	authAdminUseCase      *auth_admin.AuthAdminUseCase
}

func New() *UseCases {
	return &UseCases{}
}

func (u *UseCases) Init(ctx context.Context, repositories *repositories.Repositories) error {
	u.booksUseCase = books.NewBooksUseCase(repositories.GetBookRepository())

	u.booksUserAwareUseCase = books_user_aware.NewBooksUserAwareUseCase(
		u.booksUseCase,
		repositories.GetBookFavoriteRepository(),
	)

	u.authUserUseCase = auth_user.NewAuthUserUseCase(
		repositories.GetAuthUserRepository(),
		repositories.GetUserRepository(),
	)

	u.authAdminUseCase = auth_admin.NewAuthAdminUseCase(
		repositories.GetAuthAdminRepository(),
		repositories.GetAdminUserRepository(),
	)

	u.bookFavoriteUseCase = book_favorites.NewBooksFavoriteUseCase(
		repositories.GetBookFavoriteRepository(),
		repositories.GetBookRepository(),
	)

	return nil
}

func (u *UseCases) GetBooksUseCase() *books.BooksUseCase {
	return u.booksUseCase
}

func (u *UseCases) GetBooksUserAwareUseCase() *books_user_aware.BooksUserAwareUseCase {
	return u.booksUserAwareUseCase
}

func (u *UseCases) GetAuthUserUseCase() *auth_user.AuthUserUseCase {
	return u.authUserUseCase
}

func (u *UseCases) GetAuthAdminUseCase() *auth_admin.AuthAdminUseCase {
	return u.authAdminUseCase
}

func (u *UseCases) GetBookFavoriteUseCase() *book_favorites.BookFavoritesUseCase {
	return u.bookFavoriteUseCase
}

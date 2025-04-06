package http_adapter

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/app/use_cases"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/create_book"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/create_favorite_book"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/get_book"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/get_books"
)

type HttpAdapter struct {
	handlerGetBooks           *get_books.GetBooksHandler
	handlerGetBook            *get_book.GetBookHandler
	handlerCreateBook         *create_book.CreateBookHandler
	handlerCreateFavoriteBook *create_favorite_book.CreateFavoriteBookHandler
}

func New() *HttpAdapter {
	return &HttpAdapter{}
}

func (h *HttpAdapter) Init(ctx context.Context, uc *use_cases.UseCases) error {
	h.handlerGetBooks = get_books.NewGetBooksHandler(
		uc.GetAuthUserUseCase(),
		uc.GetBooksUseCase(),
		uc.GetBooksUserAwareUseCase(),
	)
	h.handlerGetBook = get_book.NewGetBookHandler(
		uc.GetBooksUseCase(),
	)
	h.handlerCreateBook = create_book.NewCreateBookHandler(
		uc.GetAuthAdminUseCase(),
		uc.GetBooksUseCase(),
	)
	h.handlerCreateFavoriteBook = create_favorite_book.NewCreateFavoriteBookHandler(
		uc.GetAuthUserUseCase(),
		uc.GetBookFavoriteUseCase(),
	)
	return nil
}

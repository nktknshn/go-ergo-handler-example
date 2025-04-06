package create_favorite_book

import (
	"context"
	"errors"
	"net/http"

	goergohandler "github.com/nktknshn/go-ergo-handler"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handler_builder"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handlers_params"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handlers_user_auth"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book_favorite"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	useCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/book_favorites"
)

type bookFavoriteUseCase interface {
	AddFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) (book_favorite.BookFavorite, error)
}

type userUseCase interface {
	ValidateToken(ctx context.Context, token string) (*user.User, bool, error)
}

type CreateFavoriteBookHandler struct {
	userUseCase         userUseCase
	bookFavoriteUseCase bookFavoriteUseCase
}

func NewCreateFavoriteBookHandler(userUseCase userUseCase, bookFavoriteUseCase bookFavoriteUseCase) *CreateFavoriteBookHandler {
	return &CreateFavoriteBookHandler{userUseCase, bookFavoriteUseCase}
}

func (h *CreateFavoriteBookHandler) GetMethods() []string {
	return []string{http.MethodPost, http.MethodOptions}
}

func (h *CreateFavoriteBookHandler) GetPath() string {
	return "/api/v1/books/{book_id:[0-9]+}/favorite"
}

func (h *CreateFavoriteBookHandler) GetHandler() http.Handler {
	return makeHttpRequest(h.userUseCase, h.bookFavoriteUseCase)
}

func makeHttpRequest(userUseCase userUseCase, createFavoriteBookUseCase bookFavoriteUseCase) http.Handler {
	var (
		b           = handler_builder.New()
		paramBookID = handlers_params.RouterParamBookID.Attach(b)
		auth        = handlers_user_auth.UserParser.Attach(userUseCase, b)
		handlerFunc = func(h http.ResponseWriter, r *http.Request) (any, error) {
			bookID := paramBookID.GetRequest(r)
			user := auth.GetUserRequest(r)

			favorite, err := createFavoriteBookUseCase.AddFavoriteBook(r.Context(), user.ID, bookID.ToBookID())

			if errors.Is(err, useCaseValObj.ErrBookNotFound) {
				return nil, goergohandler.WrapError(err, http.StatusNotFound)
			}

			if errors.Is(err, useCaseValObj.ErrBookAlreadyInFavorite) {
				return nil, goergohandler.WrapError(err, http.StatusConflict)
			}

			if err != nil {
				return nil, err
			}
			return favorite, nil
		}
	)

	return b.BuildHandlerWrapped(handlerFunc)
}

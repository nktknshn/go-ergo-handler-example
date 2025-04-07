package get_books

import (
	"context"
	"errors"
	"net/http"

	goergohandler "github.com/nktknshn/go-ergo-handler"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handler_builder"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handlers_user_auth"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	useCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books"
	awaredUseCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books_user_aware"
)

type authUserUseCase interface {
	ValidateToken(ctx context.Context, token string) (*user.User, bool, error)
}

type getBooksUseCase interface {
	GetBooksList(ctx context.Context, query useCaseValObj.GetBooksListQuery) (useCaseValObj.GetBooksListResponse, error)
}

type userAwareGetBooksUseCase interface {
	GetBooksList(ctx context.Context, userID user.UserID, query useCaseValObj.GetBooksListQuery) (awaredUseCaseValObj.GetBooksListResponse, error)
}

type GetBooksHandler struct {
	authUseCase              authUserUseCase
	getBooksUseCase          getBooksUseCase
	userAwareGetBooksUseCase userAwareGetBooksUseCase
}

func NewGetBooksHandler(
	authUseCase authUserUseCase,
	getBooksUseCase getBooksUseCase,
	userAwareGetBooksUseCase userAwareGetBooksUseCase,
) *GetBooksHandler {
	return &GetBooksHandler{
		authUseCase,
		getBooksUseCase,
		userAwareGetBooksUseCase,
	}
}

func (h *GetBooksHandler) GetMethods() []string {
	return []string{http.MethodGet, http.MethodOptions}
}

func (h *GetBooksHandler) GetPath() string {
	return "/api/v1/books"
}

func (h *GetBooksHandler) GetHandler() http.Handler {
	return makeHttpHandler(
		h.authUseCase,
		h.getBooksUseCase,
		h.userAwareGetBooksUseCase,
	)
}

func makeHttpHandler(
	authUseCase authUserUseCase,
	getBooksUseCase getBooksUseCase,
	userAwareGetBooksUseCase userAwareGetBooksUseCase,
) http.Handler {
	var (
		builder     = handler_builder.New()
		paramCursor = queryParamCursor.Attach(builder)
		auth        = handlers_user_auth.UserMaybe.Attach(authUseCase, builder)
	)

	return builder.BuildHandlerWrapped(func(h http.ResponseWriter, r *http.Request) (any, error) {
		user, _ := auth.GetMaybe(r)
		cursor, _ := paramCursor.GetMaybe(r)
		query := useCaseValObj.GetBooksListQuery{Cursor: cursor}

		var err error
		var resp any

		if user == nil {
			resp, err = getBooksUseCase.GetBooksList(r.Context(), query)
		} else {
			resp, err = userAwareGetBooksUseCase.GetBooksList(r.Context(), user.ID, query)
		}

		if errors.Is(err, useCaseValObj.ErrInvalidCursor) {
			return nil, goergohandler.WrapWithStatusCode(err, http.StatusBadRequest)
		}

		if err != nil {
			return nil, err
		}

		return resp, nil
	})
}

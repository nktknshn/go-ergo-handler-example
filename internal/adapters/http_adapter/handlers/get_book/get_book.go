package get_book

import (
	"context"
	"errors"
	"net/http"

	geh "github.com/nktknshn/go-ergo-handler"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handler_builder"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handlers_params"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	useCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books"
)

type getBookUseCase interface {
	GetBookByID(ctx context.Context, bookID book.BookID) (book.Book, error)
}

type GetBookHandler struct {
	getBookUseCase getBookUseCase
}

func NewGetBookHandler(getBookUseCase getBookUseCase) *GetBookHandler {
	return &GetBookHandler{getBookUseCase}
}

func (h *GetBookHandler) GetMethods() []string {
	return []string{http.MethodGet, http.MethodOptions}
}

func (h *GetBookHandler) GetPath() string {
	return "/api/v1/books/{book_id:[0-9]+}"
}

func (h *GetBookHandler) GetHandler() http.Handler {
	return makeHttpHandler(h.getBookUseCase)
}

func makeHttpHandler(getBookUseCase getBookUseCase) http.Handler {
	var (
		b           = handler_builder.New()
		paramBookID = handlers_params.RouterParamBookID.Attach(b)
		handlerFunc = func(h http.ResponseWriter, r *http.Request) (any, error) {
			bookID := paramBookID.GetRequest(r)
			book, err := getBookUseCase.GetBookByID(r.Context(), bookID)
			if errors.Is(err, useCaseValObj.ErrBookNotFound) {
				return nil, geh.NewError(http.StatusNotFound, err)
			}
			if err != nil {
				return nil, err
			}
			return book, nil
		}
	)

	return b.BuildHandlerWrapped(handlerFunc)
}

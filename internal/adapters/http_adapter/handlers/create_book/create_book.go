package create_book

import (
	"context"
	"net/http"

	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handler_admin_auth"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handler_admin_role_checker"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handler_builder"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
)

type createBookUseCase interface {
	CreateBook(ctx context.Context, publisherID admin_user.AdminUserID, book book.Book) (book.Book, error)
}

type adminUserUseCase interface {
	ValidateToken(ctx context.Context, token string) (*admin_user.AdminUser, bool, error)
}

type CreateBookHandler struct {
	adminUserUseCase  adminUserUseCase
	createBookUseCase createBookUseCase
}

func NewCreateBookHandler(adminUserUseCase adminUserUseCase, createBookUseCase createBookUseCase) *CreateBookHandler {
	return &CreateBookHandler{adminUserUseCase, createBookUseCase}
}

func (h *CreateBookHandler) GetMethods() []string {
	return []string{http.MethodPost, http.MethodOptions}
}

func (h *CreateBookHandler) GetPath() string {
	return "/api/v1/books"
}

func (h *CreateBookHandler) GetHandler() http.Handler {
	return makeHttpHandler(h.adminUserUseCase, h.createBookUseCase)
}

func makeHttpHandler(adminUserUseCase adminUserUseCase, createBookUseCase createBookUseCase) http.Handler {
	var (
		b               = handler_builder.New()
		adminUserParser = handler_admin_auth.AdminUserParser.Attach(adminUserUseCase, b)
		_               = handler_admin_role_checker.AdminAndPublisher.Attach(adminUserParser, b)
		payload         = payloadCreateBook.Attach(b)
	)

	var handlerFunc = func(h http.ResponseWriter, r *http.Request) (any, error) {
		payload := payload.GetRequest(r)
		adminUser := adminUserParser.GetUserRequest(r)
		resp, err := createBookUseCase.CreateBook(r.Context(), adminUser.ID, payload.ToBook())
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return b.BuildHandlerWrapped(handlerFunc)
}

package get_books

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/get_books"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	useCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type deps struct {
	authUseCaseMock              *AuthUseCaseMock
	getBooksUseCaseMock          *GetBooksUseCaseMock
	userAwareGetBooksUseCaseMock *UserAwareGetBooksUseCaseMock
}

func (d *deps) getHandler() *get_books.GetBooksHandler {
	return get_books.NewGetBooksHandler(
		d.authUseCaseMock,
		d.getBooksUseCaseMock,
		d.userAwareGetBooksUseCaseMock,
	)
}

func setup() *deps {
	d := &deps{}
	d.authUseCaseMock = &AuthUseCaseMock{}
	d.getBooksUseCaseMock = &GetBooksUseCaseMock{}
	d.userAwareGetBooksUseCaseMock = &UserAwareGetBooksUseCaseMock{}
	return d
}

type resp = useCaseValObj.GetBooksListResponse

func TestGetBooksHandler_GetBooks(t *testing.T) {
	deps := setup()

	deps.getBooksUseCaseMock.On("GetBooksList", mock.Anything, mock.Anything).Return(resp{}, nil)

	handler := deps.getHandler()

	h := handler.GetHandler()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
}

func TestGetBooksHandler_GetBooksTable(t *testing.T) {
	type getUserReturns struct {
		user *user.User
		ok   bool
		err  error
	}

	type testCase struct {
		uri     string
		headers map[string]string
		getUser getUserReturns
	}
}

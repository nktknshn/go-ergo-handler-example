package get_books

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	useCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books"
	awaredUseCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books_user_aware"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) ValidateToken(ctx context.Context, token string) (*user.User, bool, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(*user.User), args.Bool(1), args.Error(2)
}

type GetBooksUseCaseMock struct {
	mock.Mock
}

func (m *GetBooksUseCaseMock) GetBooksList(ctx context.Context, query useCaseValObj.GetBooksListQuery) (useCaseValObj.GetBooksListResponse, error) {
	args := m.Called(ctx, query)
	return args.Get(0).(useCaseValObj.GetBooksListResponse), args.Error(1)
}

type UserAwareGetBooksUseCaseMock struct {
	mock.Mock
}

func (m *UserAwareGetBooksUseCaseMock) GetBooksList(ctx context.Context, userID user.UserID, query useCaseValObj.GetBooksListQuery) (awaredUseCaseValObj.GetBooksListResponse, error) {
	args := m.Called(ctx, userID, query)
	return args.Get(0).(awaredUseCaseValObj.GetBooksListResponse), args.Error(1)
}

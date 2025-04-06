package get_books

import (
	"context"
	"errors"

	geh "github.com/nktknshn/go-ergo-handler"
	bookRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books"
)

var (
	paramNameCursor     = "cursor"
	errParamCursorEmpty = errors.New("param cursor is empty")
)

var queryParamCursor = geh.QueryParamMaybe(paramNameCursor, func(ctx context.Context, v string) (bookRepoValObj.BookListCursor, error) {
	if v == "" {
		return "", errParamCursorEmpty
	}
	return bookRepoValObj.BookListCursor(v), nil
})

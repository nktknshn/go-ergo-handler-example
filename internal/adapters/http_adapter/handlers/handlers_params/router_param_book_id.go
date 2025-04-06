package handlers_params

import (
	"context"

	geh "github.com/nktknshn/go-ergo-handler"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
)

const routerParamName = "book_id"

var RouterParamBookID = geh.NewRouterParam(routerParamName, func(ctx context.Context, v string) (book.BookID, error) {
	return parseBookID(v)
})

package handlers_params

import (
	"context"
	"net/http"
	"strconv"

	geh "github.com/nktknshn/go-ergo-handler"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
)

var (
	queryParamName             = "book_id"
	errParamBookIDEmpty        = geh.NewErrorStr(400, "param "+queryParamName+" is empty")
	errParamBookIDInvalid      = geh.NewErrorStr(400, "param "+queryParamName+" is invalid")
	errParamBookIDInvalidValue = geh.NewErrorStr(400, "param "+queryParamName+" has invalid value")
	errParamBookIDMissing      = geh.NewErrorStr(400, "param "+queryParamName+" is missing")
)

var errorHandler geh.HandleErrorFunc = func(_ context.Context, w http.ResponseWriter, _ *http.Request, err error) {
	w.WriteHeader(400)
	_, _ = w.Write([]byte(`{"error": "invalid request"}`))
}

func parseBookID(v string) (book.BookID, error) {
	if v == "" {
		return 0, errParamBookIDEmpty
	}
	vint, err := strconv.Atoi(v)
	if err != nil {
		return 0, errParamBookIDInvalid
	}
	if vint < 1 {
		return 0, errParamBookIDInvalidValue
	}
	return book.BookID(vint), nil
}

var QueryParamBookIDMaybe = geh.NewQueryParamMaybe(queryParamName, func(ctx context.Context, v string) (book.BookID, error) {
	return parseBookID(v)
})

var QueryParamBookID = geh.NewQueryParam(queryParamName, func(ctx context.Context, v string) (book.BookID, error) {
	return parseBookID(v)
}, errParamBookIDMissing)

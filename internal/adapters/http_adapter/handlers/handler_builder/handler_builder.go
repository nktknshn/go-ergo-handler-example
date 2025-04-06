package handler_builder

import (
	"context"
	"net/http"

	goergohandler "github.com/nktknshn/go-ergo-handler"
)

// custom error handler
var handlerErrorFunc goergohandler.HandleErrorFunc = func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
}

// custom result handler
var handlerResultFunc goergohandler.HandleResultFunc = func(ctx context.Context, w http.ResponseWriter, r *http.Request, result any) {
	w.WriteHeader(http.StatusOK)
}

func New() *goergohandler.Builder {
	return goergohandler.New()
	// .
	// 	WithHandlerErrorFunc(handlerErrorFunc).
	// 	WithHandlerResultFunc(handlerResultFunc)
}

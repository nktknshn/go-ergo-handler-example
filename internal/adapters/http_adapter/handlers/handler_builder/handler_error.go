package handler_builder

import (
	"context"
	"errors"
	"net/http"

	goergohandler "github.com/nktknshn/go-ergo-handler"
	authAdminUseCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/auth_admin"
	authUserUseCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/auth_user"
)

// custom error handler
var handlerErrorFunc goergohandler.HandleErrorFunc = func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, authAdminUseCaseValObj.ErrTokenInvalid) {
		goergohandler.DefaultHandlerErrorFunc(ctx, w, r, goergohandler.NewError(http.StatusUnauthorized, err))
		return
	}
	if errors.Is(err, authUserUseCaseValObj.ErrTokenInvalid) {
		goergohandler.DefaultHandlerErrorFunc(ctx, w, r, goergohandler.NewError(http.StatusUnauthorized, err))
		return
	}
	goergohandler.DefaultHandlerErrorFunc(ctx, w, r, err)
}

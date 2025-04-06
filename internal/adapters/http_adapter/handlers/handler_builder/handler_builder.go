package handler_builder

import (
	goergohandler "github.com/nktknshn/go-ergo-handler"
)

func New() *goergohandler.Builder {
	return goergohandler.New().
		WithHandlerErrorFunc(handlerErrorFunc)
}

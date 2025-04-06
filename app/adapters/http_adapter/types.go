package http_adapter

import "net/http"

type HttpHandlerAdder interface {
	AddHandler(handler interface {
		GetHandler() http.Handler
		GetMethods() []string
		GetPath() string
	})
}

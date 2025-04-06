package http_server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	router *mux.Router
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		router: mux.NewRouter(),
	}
}

func (s *HttpServer) AddHandler(handler interface {
	GetHandler() http.Handler
	GetMethods() []string
	GetPath() string
}) {
	s.router.Handle(
		handler.GetPath(),
		handler.GetHandler(),
	).Methods(handler.GetMethods()...)
}

func (s *HttpServer) Start() {
	http.ListenAndServe(":8080", s.router)
}

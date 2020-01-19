package http

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"banhmi/root"
)

type Service struct {
	Root root.AppService
	router *mux.Router
}

func (s *Service) NewRouter() {
	s.router = mux.NewRouter()
}

// Router ..
func (s *Service) MuxRouter() *mux.Router {
	return s.router
}

func (s *Service) SetRouters() {
	fmt.Println("set routers from http package")
	s.Root.SetRouters()
}

func (s *Service) ListenAndServe() {
	http.ListenAndServe(":8888", s.router)
}
package server

import (
	"database/sql"
	"net/http"

	"github.com/kapralovs/wb0/internal/logger"
)

const (
	subject = ""
)

func New() *server {
	return &server{}
}

func (s *server) Run() error {
	
	s.initRouter()
	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *server) initRouter() {
	s.router.HandleFunc("/order/{id}", getByID()).Methods("GET")
}

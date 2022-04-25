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
	db, err := newDB(s.config.DatabaseURL)
	if err != nil {
		logger.Error(subject).Err(err).Msg("cannot connect to postgres")
	}
	s.storage.db:=
	s.initRouter()
	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *server) initRouter() {
	s.router.HandleFunc("/order/{id}", getByID()).Methods("GET")
}

func newDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

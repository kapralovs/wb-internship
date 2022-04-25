package server

import (
	"github.com/gorilla/mux"
	"github.com/kapralovs/wb0/internal/storage"
)

type server struct {
	config  *Config
	router  *mux.Router
	storage *storage.Storage
}

type Config struct {
	Port        string `toml:"port"`
	DatabaseURL string `toml:"database_url"`
}

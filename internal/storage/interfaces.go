package storage

import (
	"database/sql"

	"github.com/kapralovs/wb0/internal/storage/inmemory"
)

type Storage struct {
	cache inmemory.Cache
	db    *sql.DB
}

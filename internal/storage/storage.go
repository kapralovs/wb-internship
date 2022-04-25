package storage

import "github.com/kapralovs/wb0/internal/logger"

const (
	subject = "storage"
)

func New() *Storage {
	return &Storage{}
}

func (s *Storage) RestoreCache() {
	query := `SELECT * FROM`
	rows, err := s.db.Query(query)
	if err != nil {
		logger.Error(subject).Err(err).Msg("cannot get rows from db")
	}
}

func Init() {
	db, err := newDB(s.config.DatabaseURL)
	if err != nil {
		logger.Error(subject).Err(err).Msg("cannot connect to postgres")
	}

}

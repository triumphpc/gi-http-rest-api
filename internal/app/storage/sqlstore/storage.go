package sqlstore

import (
	"database/sql"
)
import _ "github.com/lib/pq"

type Store struct {
	db   *sql.DB
	user *UserRepository
}

// New storage entity
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() *UserRepository {
	if s.user == nil {
		s.user = &UserRepository{
			store: s,
		}
	}

	return s.user
}

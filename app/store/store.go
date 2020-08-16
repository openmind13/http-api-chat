package store

import (
	"database/sql"

	// postgresql driver
	_ "github.com/lib/pq"
)

// SQLStore struct
type SQLStore struct {
	db *sql.DB
}

// NewSQLStore ...
func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

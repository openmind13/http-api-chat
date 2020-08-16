package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/openmind13/http-api-chat/app/store"
)

// Start ...
func Start(config *Config) error {
	// oped database with params from config
	db, err := initDB(config.DatabaseDriver, config.DatabaseURL)
	if err != nil {
		return err
	}

	// close db on exit
	defer db.Close()

	// Create new sql store
	store := store.NewSQLStore(db)

	server := newServer(store)

	if err := http.ListenAndServe(config.BindAddr, server); err != nil {
		return err
	}

	return nil
}

func initDB(databaseDriver, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open(databaseDriver, databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, errConnDB
	}

	return db, nil
}

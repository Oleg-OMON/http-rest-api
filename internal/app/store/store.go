package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	DB     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=260616 dbname=football_database sslmode=disable")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.DB = db
	return nil
}

func (s *Store) Close() {
	s.DB.Close()
}

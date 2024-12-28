package storage

import (
	"database/sql"
	"my_lib_log/internal/storage/db/postgres"
)

type StorageInterface interface {
	AddLog() (int, error)
}

type Storage struct {
	StorageInterface
}

func New(db *sql.DB) *Storage {
	return &Storage{
		StorageInterface: postgres.New(db),
	}
}

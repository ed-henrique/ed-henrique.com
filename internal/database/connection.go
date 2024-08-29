package database

import (
	"database/sql"
	"log/slog"
)

func NewDBConnection(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		slog.Error("Could not open database", "err", err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		slog.Error("Could not ping database", "err", err.Error())
		panic(err)
	}

	return db
}

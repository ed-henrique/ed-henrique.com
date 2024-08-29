package main

//go:generate rm -f db/db.sqlite
//go:generate sqlite3 db/db.sqlite ".read ./internal/database/schema.sql"

//go:generate sqlc -f ./internal/database/sqlc.yaml generate

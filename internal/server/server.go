package server

import (
	"database/sql"
	"net/http"

	"github.com/ed-henrique/ed-henrique.com/internal/database"
)

type Server struct {
	Port          string
	Mux           *http.ServeMux

	q             *database.Queries
	version       int
	adminUsername string
	adminPassword string
	disableDev    bool
}

func New(db *sql.DB, port string, version int, adminUsername, adminPassword string, disableDev bool) *Server {
	return &Server{
		Port:          port,
		Mux:           http.NewServeMux(),
		q:             database.New(db),
		version:       version,
		adminUsername: adminUsername,
		adminPassword: adminPassword,
		disableDev:    disableDev,
	}
}


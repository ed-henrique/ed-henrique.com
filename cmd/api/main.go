package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/ed-henrique/ed-henrique.com/internal/database"
	"github.com/ed-henrique/ed-henrique.com/internal/server"
)

func main() {
	disableDev := flag.Bool("disable-dev", false, "disable dev mode")
	port := flag.String("port", "8080", "server port")
	dbPath := flag.String("db", "db/db.sqlite", "DB path")

	flag.Parse()

	versionFile := os.Getenv("STATIC_FILES_VERSION_FILE")
	if versionFile == "" {
		slog.Error("The passed static files version file is empty")
		return
	}

	versionRaw, err := os.ReadFile(versionFile)
	if err != nil {
		slog.Error("Could not read static files version file", "err", err.Error())
		panic(err)
	}

	version, err := strconv.Atoi(string(versionRaw))
	if err != nil {
		slog.Error("Could not convert static files version to int", "err", err.Error())
		panic(err)
	}

	adminUsernameFile := os.Getenv("ADMIN_USERNAME_FILE")
	if adminUsernameFile == "" {
		slog.Error("The passed admin username is empty")
		return
	}

	adminUsernameRaw, err := os.ReadFile(adminUsernameFile)
	if err != nil {
		slog.Error("Could not read admin username file", "err", err.Error())
		panic(err)
	}

	adminUsername := string(adminUsernameRaw)
	if adminUsername == "" {
		slog.Error("The passed admin username is empty")
		return
	}

	adminPasswordFile := os.Getenv("ADMIN_PASSWORD_FILE")
	if adminPasswordFile == "" {
		slog.Error("The passed admin password is empty")
		return
	}

	adminPasswordRaw, err := os.ReadFile(adminPasswordFile)
	if err != nil {
		slog.Error("Could not read admin password file", "err", err.Error())
		panic(err)
	}

	adminPassword := string(adminPasswordRaw)
	if adminPassword == "" {
		slog.Error("The passed admin password is empty")
		return
	}

	db := database.NewDBConnection(*dbPath)

	s := server.New(
		db,
		*port,
		version,
		adminUsername,
		adminPassword,
		*disableDev,
	)

	s.RegisterRoutes()
	if err := http.ListenAndServe(":"+s.Port, s.Mux); err != nil {
		slog.Error("The server could not be started.", slog.Any("err", err))
	}
}

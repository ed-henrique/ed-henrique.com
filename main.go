package main

import (
	"flag"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/ed-henrique/ed-henrique.com/pages"
)

type Server struct {
	port string
	mux  *http.ServeMux
}

func disableCacheInDevMode(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func main() {
	s := &Server{
		port: *flag.String("port", "8080", "server port"),
		mux:  http.NewServeMux(),
	}

	flag.Parse()

	s.mux.Handle("/public/", disableCacheInDevMode(http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))))

	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/" {
			http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		}
	})

	s.mux.Handle("GET /home", templ.Handler(pages.Index()))
	s.mux.Handle("GET /cv", templ.Handler(pages.CV()))
	s.mux.Handle("GET /posts", templ.Handler(pages.Posts()))
	s.mux.Handle("GET /posts/:name", templ.Handler(pages.Post()))
	s.mux.Handle("GET /projects", templ.Handler(pages.Projects()))
	s.mux.Handle("GET /projects/:name", templ.Handler(pages.Project()))

	if err := http.ListenAndServe(":"+s.port, s.mux); err != nil {
		slog.Error("The server could not be started.", slog.Any("err", err))
	}
}

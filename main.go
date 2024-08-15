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

func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self' https://fonts.googleapis.com https://fonts.gstatic.com")
		next.ServeHTTP(w, r)
	})
}

func disableCacheInDevMode(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func main() {
	disableDev := flag.Bool("disable-dev", false, "disable dev mode")
	port := flag.String("port", "8080", "server port")

	flag.Parse()

	s := &Server{
		port: *port,
		mux:  http.NewServeMux(),
	}

	if *disableDev {
		s.mux.Handle("/public/", func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Cache-Control", "max-age=31536000")
				next.ServeHTTP(w, r)
			})
		}(http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))))
	} else {
		s.mux.Handle("/public/", disableCacheInDevMode(http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))))
	}

	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/" {
			http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		}
	})

	s.mux.Handle("GET /home", commonHeaders(templ.Handler(pages.Index())))
	s.mux.Handle("GET /cv", commonHeaders(templ.Handler(pages.CV())))
	s.mux.Handle("GET /posts", commonHeaders(templ.Handler(pages.Posts())))
	s.mux.Handle("GET /posts/:name", commonHeaders(templ.Handler(pages.Post())))
	s.mux.Handle("GET /projects", commonHeaders(templ.Handler(pages.Projects())))
	s.mux.Handle("GET /projects/:name", commonHeaders(templ.Handler(pages.Project())))

	if err := http.ListenAndServe(":"+s.port, s.mux); err != nil {
		slog.Error("The server could not be started.", slog.Any("err", err))
	}
}

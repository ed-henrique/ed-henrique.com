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
  mux *http.ServeMux
}

func main() {
  s := &Server{
    port: *flag.String("port", "8080", "server port"),
    mux: http.NewServeMux(),
  }

  flag.Parse()

  s.mux.Handle("GET /", templ.Handler(pages.Index()))
  s.mux.Handle("GET /cv", templ.Handler(pages.CV()))
  s.mux.Handle("GET /about", templ.Handler(pages.About()))
  s.mux.Handle("GET /posts", templ.Handler(pages.Posts()))
  s.mux.Handle("GET /posts/:name", templ.Handler(pages.Post()))
  s.mux.Handle("GET /projects", templ.Handler(pages.Projects()))
  s.mux.Handle("GET /projects/:name", templ.Handler(pages.Project()))

  if err := http.ListenAndServe(":" + s.port, s.mux); err != nil {
    slog.Error("The server could not be started.", slog.Any("err", err))
  }
}

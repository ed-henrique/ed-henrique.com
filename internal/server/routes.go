package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/ed-henrique/ed-henrique.com/admin"
	"github.com/ed-henrique/ed-henrique.com/pages"

	_ "modernc.org/sqlite"
)

func (s *Server) RegisterRoutes() {
	// Static Files
	if s.disableDev {
		s.Mux.Handle("/public/", cache(http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))))
	} else {
		s.Mux.Handle("/public/", disableCacheInDevMode(http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))))
	}

	// Public
	s.Mux.HandleFunc("/", s.index)
	s.Mux.Handle("GET /home/", commonHeaders(http.HandlerFunc(s.home)))
	s.Mux.Handle("GET /cv/", commonHeaders(templ.Handler(pages.CV(s.version))))
	s.Mux.Handle("GET /posts/", commonHeaders(http.HandlerFunc(s.posts)))
	s.Mux.Handle("GET /posts/{id}", commonHeaders(http.HandlerFunc(s.post)))
	s.Mux.Handle("GET /projects/", commonHeaders(templ.Handler(pages.Projects(s.version))))
	s.Mux.Handle("GET /projects/{id}", commonHeaders(templ.Handler(pages.Project(s.version))))

	// Admin
	s.Mux.Handle("GET /admin/", basicAuth(commonHeaders(http.HandlerFunc(s.adminIndex))))
	s.Mux.Handle("GET /admin/posts/{id}", basicAuth(commonHeaders(http.HandlerFunc(s.adminEditPost))))
	s.Mux.Handle("GET /admin/posts/create", basicAuth(commonHeaders(templ.Handler(admin.CreatePost(s.version)))))

	// Admin API
	s.Mux.Handle("POST /admin/posts/", basicAuth(http.HandlerFunc(s.adminAPICreatePost)))
	s.Mux.Handle("PATCH /admin/posts/{id}", basicAuth(http.HandlerFunc(s.adminAPIEditPost)))
	s.Mux.Handle("DELETE /admin/posts/{id}", basicAuth(http.HandlerFunc(s.adminAPIRemovePost)))
}

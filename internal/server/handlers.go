package server

import (
	"database/sql"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ed-henrique/ed-henrique.com/admin"
	"github.com/ed-henrique/ed-henrique.com/internal/database"
	"github.com/ed-henrique/ed-henrique.com/internal/md"
	"github.com/ed-henrique/ed-henrique.com/pages"
)

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.Path == "/" {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	}
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
		ps, err := s.q.ListPosts(r.Context())
		if err != nil {
			slog.Error("Could not get post", "err", err.Error())
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		var contentHtml string
		for _, p := range ps {
			contentHtml = string(md.PostToHtml([]byte(p.Content)))
			p.Content = contentHtml
		}

		posts := make([]database.Post, len(ps))
		for i := range len(ps) {
			posts[i] = database.Post{
				ID:          ps[i].ID,
				Title:       ps[i].Title,
				Content:     ps[i].Content,
				Description: ps[i].Description,
				CreatedAt:   ps[i].CreatedAt,
				UpdatedAt:   ps[i].UpdatedAt,
			}
		}

		err = pages.Index(s.version, posts).Render(r.Context(), w)
		if err != nil {
			slog.Error("Could not render template", "err", err.Error())
			http.Error(w, "Rendering error", http.StatusInternalServerError)
			return
		}
}

func (s *Server) posts(w http.ResponseWriter, r *http.Request) {
		ps, err := s.q.ListPosts(r.Context())
		if err != nil {
			slog.Error("Could not get post", "err", err.Error())
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		var contentHtml string
		for _, p := range ps {
			contentHtml = string(md.PostToHtml([]byte(p.Content)))
			p.Content = contentHtml
		}

		posts := make([]database.Post, len(ps))
		for i := range len(ps) {
			posts[i] = database.Post{
				ID:          ps[i].ID,
				Title:       ps[i].Title,
				Content:     ps[i].Content,
				Description: ps[i].Description,
				CreatedAt:   ps[i].CreatedAt,
				UpdatedAt:   ps[i].UpdatedAt,
			}
		}

		err = pages.Posts(s.version, posts).Render(r.Context(), w)
		if err != nil {
			slog.Error("Could not render template", "err", err.Error())
			http.Error(w, "Rendering error", http.StatusInternalServerError)
			return
		}
}

func (s *Server) post(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.ParseInt(idRaw, 10, 64)
		if err != nil {
			http.Error(w, "ID is not a valid int", http.StatusBadRequest)
			return
		}

		if id <= 0 {
			http.Error(w, "ID not greater than 0", http.StatusBadRequest)
			return
		}

		p, err := s.q.GetPost(r.Context(), id)
		if err != nil {
			slog.Error("Could not get post", "err", err.Error())
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		contentHtml := md.PostToHtml([]byte(p.Content))
		err = pages.Post(s.version, database.Post{
			ID:          id,
			Title:       p.Title,
			Description: p.Description,
			Content:     string(contentHtml),
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		}).Render(r.Context(), w)
		if err != nil {
			slog.Error("Could not render template", "err", err.Error())
			http.Error(w, "Rendering error", http.StatusInternalServerError)
			return
		}
}

func (s *Server) adminIndex(w http.ResponseWriter, r *http.Request) {
		ps, err := s.q.ListPosts(r.Context())
		if err != nil {
			if err == sql.ErrNoRows {
				ps = []database.ListPostsRow{}
			} else {
				slog.Error("Could not list posts", "err", err.Error())
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
		}

		posts := make([]database.Post, len(ps))
		for i, p := range ps {
			posts[i] = database.Post{
				ID:          p.ID,
				Title:       p.Title,
				Content:     p.Content,
				Description: p.Description,
				CreatedAt:   p.CreatedAt,
				UpdatedAt:   p.UpdatedAt,
			}
		}

		err = admin.Index(s.version, posts).Render(r.Context(), w)
		if err != nil {
			slog.Error("Could not render template", "err", err.Error())
			http.Error(w, "Rendering error", http.StatusInternalServerError)
			return
		}
}

func (s *Server) adminEditPost(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.ParseInt(idRaw, 10, 64)
		if err != nil {
			http.Error(w, "ID is not a valid int", http.StatusBadRequest)
			return
		}

		if id <= 0 {
			http.Error(w, "ID not greater than 0", http.StatusBadRequest)
			return
		}

		p, err := s.q.GetPost(r.Context(), id)
		if err != nil {
			slog.Error("Could not get post", "err", err.Error())
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		err = admin.EditPost(s.version, database.Post{
			ID:          id,
			Title:       p.Title,
			Content:     p.Content,
			Description: p.Description,
		}).Render(r.Context(), w)
		if err != nil {
			slog.Error("Could not render template", "err", err.Error())
			http.Error(w, "Rendering error", http.StatusInternalServerError)
			return
		}
}

func (s *Server) adminAPICreatePost(w http.ResponseWriter, r *http.Request) {
		body := struct {
			title       string
			content     string
			description string
		}{}

		r.ParseForm()
		body.title = r.Form.Get("title")
		body.content = r.Form.Get("content")
		body.description = r.Form.Get("description")

		if body.title == "" && body.content == "" && body.description == "" {
			http.Error(w, "Nothing to add", http.StatusBadRequest)
			return
		}

		if _, err := s.q.CreatePost(r.Context(), database.CreatePostParams{
			Title:       body.title,
			Content:     body.content,
			Description: body.description,
		}); err != nil {
			slog.Error("There was an error in the database", "err", err.Error())
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
}

func (s *Server) adminAPIEditPost(w http.ResponseWriter, r *http.Request) {
		body := struct {
			title       string
			content     string
			description string
		}{}

		idRaw := r.PathValue("id")
		id, err := strconv.ParseInt(idRaw, 10, 64)
		if err != nil {
			http.Error(w, "ID is not a valid int", http.StatusBadRequest)
			return
		}

		if id <= 0 {
			http.Error(w, "ID not greater than 0", http.StatusBadRequest)
			return
		}

		r.ParseForm()
		body.title = r.Form.Get("title")
		body.content = r.Form.Get("content")
		body.description = r.Form.Get("description")

		if body.title == "" && body.content == "" && body.description == "" {
			http.Error(w, "Nothing to update", http.StatusBadRequest)
			return
		}

		if err := s.q.UpdatePost(r.Context(), database.UpdatePostParams{
			ID:      id,
			Title:   body.title,
			Content: body.content,
			Description: body.description,
		}); err != nil {
			slog.Error("There was an error in the database", "err", err.Error())
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
}

func (s *Server) adminAPIRemovePost(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.ParseInt(idRaw, 10, 64)
		if err != nil {
			http.Error(w, "ID is not a valid int", http.StatusBadRequest)
			return
		}

		if id <= 0 {
			http.Error(w, "ID not greater than 0", http.StatusBadRequest)
			return
		}

		if err := s.q.DeletePost(r.Context(), id); err != nil {
			slog.Error("There was an error in the database", "err", err.Error())
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
}

package api

import (
	"net/http"
	"strings"

	"github.com/aksel/todo-api/internal/store"
)

// NewServer returns an http.Handler with all API routes registered.
func NewServer(s *store.Memory) http.Handler {
	h := &Handler{Store: s}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /tasks", h.List)
	mux.HandleFunc("POST /tasks", h.Create)
	mux.HandleFunc("POST /tasks/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/toggle") {
			h.Toggle(w, r)
			return
		}
		http.NotFound(w, r)
	})
	mux.HandleFunc("DELETE /tasks/", h.Delete)
	return mux
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

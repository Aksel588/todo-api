package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/aksel/todo-api/internal/store"
)

type Handler struct {
	Store *store.Memory
}

type createBody struct {
	Title string `json:"title"`
}

func (h *Handler) List(w http.ResponseWriter, _ *http.Request) {
	tasks := h.Store.List()
	writeJSON(w, http.StatusOK, tasks)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body createBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	body.Title = strings.TrimSpace(body.Title)
	if body.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}
	t := h.Store.Create(body.Title)
	writeJSON(w, http.StatusCreated, t)
}

func (h *Handler) Toggle(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id = strings.TrimSuffix(id, "/toggle")
	id = strings.TrimSpace(id)
	if id == "" {
		http.NotFound(w, r)
		return
	}
	t, err := h.Store.Toggle(id)
	if errors.Is(err, store.ErrNotFound) {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id = strings.TrimSpace(id)
	if id == "" {
		http.NotFound(w, r)
		return
	}
	if err := h.Store.Delete(id); errors.Is(err, store.ErrNotFound) {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

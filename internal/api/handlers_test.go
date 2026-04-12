package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aksel/todo-api/internal/store"
)

func TestServer_health(t *testing.T) {
	srv := NewServer(store.NewMemory())
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d", rec.Code)
	}
}

func TestServer_CRUD(t *testing.T) {
	srv := NewServer(store.NewMemory())

	// list empty
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tasks", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("GET /tasks: %d", rec.Code)
	}
	var empty []store.Task
	if err := json.NewDecoder(rec.Body).Decode(&empty); err != nil {
		t.Fatal(err)
	}
	if len(empty) != 0 {
		t.Fatalf("want empty, got %d", len(empty))
	}

	// create
	body := bytes.NewBufferString(`{"title":"  learn go  "}`)
	rec = httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	req.Header.Set("Content-Type", "application/json")
	srv.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("POST /tasks: %d body=%s", rec.Code, rec.Body.String())
	}
	var task store.Task
	if err := json.NewDecoder(rec.Body).Decode(&task); err != nil {
		t.Fatal(err)
	}
	if task.Title != "learn go" || task.ID != "1" {
		t.Fatalf("task: %+v", task)
	}

	// toggle
	rec = httptest.NewRecorder()
	srv.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/tasks/1/toggle", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("toggle: %d", rec.Code)
	}
	if err := json.NewDecoder(rec.Body).Decode(&task); err != nil {
		t.Fatal(err)
	}
	if !task.Done {
		t.Fatal("expected done")
	}

	// delete
	rec = httptest.NewRecorder()
	srv.ServeHTTP(rec, httptest.NewRequest(http.MethodDelete, "/tasks/1", nil))
	if rec.Code != http.StatusNoContent {
		t.Fatalf("delete: %d", rec.Code)
	}

	// not found
	rec = httptest.NewRecorder()
	srv.ServeHTTP(rec, httptest.NewRequest(http.MethodDelete, "/tasks/1", nil))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("delete again: %d", rec.Code)
	}
}

func TestServer_Create_validation(t *testing.T) {
	srv := NewServer(store.NewMemory())
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(`{"title":""}`)))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", rec.Code)
	}
}

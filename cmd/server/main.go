package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aksel/todo-api/internal/api"
	"github.com/aksel/todo-api/internal/store"
)

func main() {
	addr := ":8080"
	if v := os.Getenv("PORT"); v != "" {
		addr = ":" + strings.TrimPrefix(strings.TrimSpace(v), ":")
	}

	s := store.NewMemory()
	srv := api.NewServer(s)

	log.Printf("listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, srv); err != nil {
		log.Fatal(err)
	}
}

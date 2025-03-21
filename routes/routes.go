package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Task Manager API is up and running!"))
	})

	return r
}

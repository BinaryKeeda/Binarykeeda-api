package routes

import (
	"binarykeeda-api/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Register(userHandler *handlers.UserHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", handlers.Health)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetAll)
		r.Post("/", userHandler.Create)
	})

	return r
}

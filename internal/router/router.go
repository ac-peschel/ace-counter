package router

import (
	"github.com/ac-peschel/ace-counter/internal/handler"
	"github.com/go-chi/chi/v5"
)

func HandleRoutes(r chi.Router) {
	r.Get("/", handler.HomeHandler)
}

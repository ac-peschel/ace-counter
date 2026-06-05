package router

import (
	"database/sql"

	"github.com/ac-peschel/ace-counter/internal/handler"
	"github.com/go-chi/chi/v5"
)

func HandleRoutes(r chi.Router, db *sql.DB) {
	r.Get("/", handler.HomeHandler(db))
	r.Get("/counter", handler.CounterHandler(db))
	r.Get("/counter/edit/{id}", handler.EditCounterHandler(db))
	r.Get("/counter/delete/{id}", handler.RemoveCounterHandler(db))
	r.Get("/counter/create", handler.CreateCounterHandler(db))
	r.Post("/counter/save/{id}", handler.SaveCounterHandler(db))
	r.Get("/increment/{id}", handler.IncrementHandler(db))
}

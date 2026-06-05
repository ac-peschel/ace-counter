package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	idb "github.com/ac-peschel/ace-counter/internal/db"
	"github.com/ac-peschel/ace-counter/internal/model"
	"github.com/ac-peschel/ace-counter/internal/templates"
	"github.com/go-chi/chi/v5"
)

func IncrementHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		cid, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid counter id", http.StatusBadRequest)
			return
		}
		_, err = idb.CreateIncrement(db, cid)
		if err != nil {
			http.Error(w, "error creating increment", http.StatusBadRequest)
			return
		}

		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
	}
}

func HomeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("t")
		existingCounter, err := idb.GetAllCounter(db)
		if err != nil {
			http.Error(w, "Database error: select counters", http.StatusInternalServerError)
			return
		}

		incs := []model.Increment{}
		for _, c := range existingCounter {
			sum, err := idb.GetIncrements(db, c.Id)
			if err != nil {
				continue
			}
			incs = append(incs, model.Increment{
				Counter: c,
				Sum:     sum,
			})
		}

		c := templates.Index(incs)
		err = templates.Layout(c).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}

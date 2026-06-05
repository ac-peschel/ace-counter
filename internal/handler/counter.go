package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	idb "github.com/ac-peschel/ace-counter/internal/db"
	"github.com/ac-peschel/ace-counter/internal/templates"
	"github.com/go-chi/chi/v5"
)

func RemoveCounterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		cid, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid counter id", http.StatusBadRequest)
			return
		}

		err = idb.DeleteCounter(db, cid)
		if err != nil {
			http.Error(w, "error deleting counter", http.StatusBadRequest)
			return
		}

		w.Header().Set("HX-Redirect", "/counter")
		w.WriteHeader(http.StatusOK)
	}
}

func SaveCounterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		cid, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid counter id", http.StatusBadRequest)
			return
		}
		err = r.ParseForm()
		if err != nil {
			http.Error(w, "error parsing form", http.StatusBadRequest)
			return
		}
		newName := r.FormValue("cname")
		newOwner := r.FormValue("cowner")

		err = idb.UpdateCounter(db, cid, newName, newOwner)
		if err != nil {
			http.Error(w, "error updating counter", http.StatusBadRequest)
			return
		}

		w.Header().Set("HX-Redirect", "/counter")
		w.WriteHeader(http.StatusOK)
	}
}

func CreateCounterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := idb.CreateCounter(db)
		if err != nil {
			http.Error(w, "error creating new counter", http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Redirect", fmt.Sprintf("/counter/edit/%d", c.Id))
		w.WriteHeader(http.StatusOK)
	}
}

func EditCounterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		cid, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid counter id", http.StatusBadRequest)
			return
		}
		counter, err := idb.GetCounterByID(db, cid)
		if err != nil {
			http.Error(w, "no counter found", http.StatusBadRequest)
			return
		}

		t := templates.EditCounter(*counter)
		err = templates.Layout(t).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}

func CounterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		existingCounter, err := idb.GetAllCounter(db)
		if err != nil {
			http.Error(w, "Database error: select counters", http.StatusInternalServerError)
		}

		t := templates.Counter(existingCounter)
		err = templates.Layout(t).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}

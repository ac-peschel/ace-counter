package handler

import (
	"net/http"

	"github.com/ac-peschel/ace-counter/internal/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	c := templates.Index()
	err := templates.Layout(c).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

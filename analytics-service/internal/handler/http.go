package handler

import (
	"analytics-service/internal/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Store *storage.Storage
}

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	visits, err := h.Store.GetVisits(code)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(visits)
}

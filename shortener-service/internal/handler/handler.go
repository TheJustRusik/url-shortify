package handler

import (
	"encoding/json"
	"net/http"
	"shortener-service/internal/grpcclient"
	"shortener-service/internal/service"
	"shortener-service/internal/storage"

	"github.com/gorilla/mux"
)

type Handler struct {
	Storage         *storage.Storage
	AnalyticsClient *grpcclient.AnalyticsClient
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Short string `json:"short"`
}

// Shorten godoc
// @Summary      Shorten a URL
// @Description  Generate a short code for a given URL
// @Tags         shortener
// @Accept       json
// @Produce      json
// @Param        input  body  shortenRequest  true  "URL to shorten"
// @Success      200    {object}  shortenResponse
// @Failure      400    {string}  string  "Invalid input"
// @Failure      500    {string}  string  "Internal server error"
// @Router       /shorten [post]
func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest
	json.NewDecoder(r.Body).Decode(&req)

	code := service.GenerateShortCode(6)
	err := h.Storage.SaveURL(code, req.URL)
	if err != nil {
		http.Error(w, "Failed to store URL", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shortenResponse{Short: code})
}

// Expand godoc
// @Summary      Expand short code
// @Description  Redirect to the original URL by short code
// @Tags         shortener
// @Produce      plain
// @Param        code  path  string  true  "Short code"
// @Success      302
// @Failure      404    {string}  string  "Not found"
// @Router       /{code} [get]
func (h *Handler) Expand(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	url, err := h.Storage.GetOriginal(code)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	// после получения URL
	go h.AnalyticsClient.SendVisit(code, r.RemoteAddr, r.UserAgent(), url)
	http.Redirect(w, r, url, http.StatusFound)
}

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/blog-log/extractor/pkg/api/v1/extract"
)

func (h *ExtractHandler) Extract(w http.ResponseWriter, r *http.Request) {
	// extract request
	var request extract.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// do actual work
	repo, err := h.extractor(r.Context(), request.Repo)
	if err != nil {
		respondWithError(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(repo); err != nil {
		respondWithError(w, err.Error(), 500)
		return
	}
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	respondWithJSON(w, code, map[string][]string{"error": {message}})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

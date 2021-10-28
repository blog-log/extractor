package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/blog-log/extractor/pkg/api/v1/extract"
)

func (h *ExtractHandler) Extract(w http.ResponseWriter, r *http.Request) {
	// extract request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var request extract.Request
	if err := json.Unmarshal(reqBody, &request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// do actual work
	repo, err := h.extractor(r.Context(), request.Repo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repo)
}

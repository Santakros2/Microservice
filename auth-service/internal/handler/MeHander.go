package handler

import (
	"encoding/json"
	"net/http"
)

func (h *AuthHandler) MeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{
		"message": "Welcome user",
	}
	json.NewEncoder(w).Encode(resp)
}

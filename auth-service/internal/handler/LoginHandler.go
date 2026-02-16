package handler

import (
	"auth-service/internal/dto"
	"encoding/json"
	"net/http"
)

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}

	var req dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	token, err := h.service.LoginService(r.Context(), req.Email, req.Password)

	if err != nil {
		http.Error(w, "invalid crendentials", http.StatusUnauthorized)
		return
	}

	resp := map[string]string{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

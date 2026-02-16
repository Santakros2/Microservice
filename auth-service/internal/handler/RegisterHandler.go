package handler

import (
	"auth-service/internal/dto"
	"encoding/json"
	"net/http"
)

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// validate the method
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	// Read the request
	var req dto.SignupRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// validate the data
	if req.Email == "" || req.Password == "" || req.Name == "" {
		http.Error(w, "Not Allowed", http.StatusBadRequest)
		return
	}
	// call the service
	user, e := h.service.SignupService(r.Context(), "user", req.Name, req.Email, req.Password)
	// return the res
	if e != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

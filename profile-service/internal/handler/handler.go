package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"profile-service/internal/domain"
	"profile-service/internal/service"
)

type ProfileHandler struct {
	Service *service.ProfileService
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {

	email := r.Header.Get("X-User-Email")

	if email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	profile, err := h.Service.GetProfile(r.Context(), email)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if profile == nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(profile)
}

func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	log.Println("EMail:", req.Email, "  name:", req.Name)

	profile := domain.NewProfile(req.Email, req.Name)

	if err := h.Service.Create(r.Context(), profile); err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

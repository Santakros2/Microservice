package handler

import (
	"auth-service/internal/service"
	"net/http"
	"strings"
)

type AuthHandler struct {
	service *service.Service
	// 	jwt     func(http.Handler) http.Handler
}

func NewAuthHandler(s *service.Service) *AuthHandler {
	return &AuthHandler{
		service: s,
		// jwt:     jwt,
	}
}

func (h *AuthHandler) AuthRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/auth")

	switch path {
	case "/register":
		h.RegisterHandler(w, r)

	case "/login":
		h.LoginHandler(w, r)

	case "/refresh":
		h.RefreshHandler(w, r)

	case "/logout":
		h.LogoutHandler(w, r)

	case "/me":
		// h.MeHandler()
		// h.jwt(http.HandlerFunc(h.MeHandler)).ServeHTTP(w, r)
		// h.protected(h.MeHandler).ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}

}

// helper func because the 40line is ugly
// func (h *AuthHandler) protected(fn func(http.ResponseWriter, *http.Request)) http.Handler {
// 	return h.jwt(http.HandlerFunc(fn))
// }

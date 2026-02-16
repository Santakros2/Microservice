package router

import (
	"gateway/internal/config"
	"gateway/internal/middleware"
	"gateway/internal/proxy"
	"net/http"
)

func New(cfg *config.Config) http.Handler {

	mux := http.NewServeMux()

	// protectedProxy := proxy.New(cfg.ProtectedURL)

	authProxy := proxy.New(cfg.AuthURL)
	profileProxy := proxy.New(cfg.ProfileURL)

	authLimiter := middleware.NewRateLimiter(2, 5)
	profileLimiter := middleware.NewRateLimiter(50, 100)

	mux.Handle("/auth/",
		authLimiter(http.StripPrefix("/auth", authProxy)),
	)

	mux.Handle("/profile/",
		profileLimiter(middleware.JWT(cfg.AuthSecret,
			http.StripPrefix("/profile", profileProxy),
		)),
	)

	return middleware.Logger(mux)
}

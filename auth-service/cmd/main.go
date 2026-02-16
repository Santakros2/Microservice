package main

import (
	"auth-service/config"
	"auth-service/internal/db"
	"auth-service/internal/handler"
	middleware "auth-service/internal/middlerware"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"log"
	"net/http"
)

// I am using the default net/http and ServeMux
// to understanding the internal working of the http package.
// Cons:
// 1 . Have to handle the REST API route manually (put, post, get, delete)
// 2 . It matches route on best matching pattern like / will match all the routes. So we have to handle that ourself.

// Solution :
// REST API:
// |Method    |Route          |Purpose                   |
// |POST      |/auth/register |create account            |
// |POST      |/auth/login    |authenticate user         |
// |POST      |/auth/refresh  |refresh JWT               |
// |POST      |/auth/logout   |revoke token              |
// |GET       |/auth/me       |current user info         |
// |POST      |/auth/validate |internal token validation |

// Routing Structure :
// /auth          → collection router
// /auth/         → action router

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()

	r := repository.NewRepo(db.NewMySQL())
	s := service.NewService(cfg.JWTSecret, r)
	// Cofig the jwt in here
	// jwtMW := middleware.JWT(cfg.JWTSecret)
	han := handler.NewAuthHandler(s)

	mux.HandleFunc("/login", han.LoginHandler)
	mux.HandleFunc("/signup", han.RegisterHandler)

	// protected now return the Handler that then
	// protected := func(h http.HandlerFunc) http.Handler {
	// 	return jwtMW(h)
	// }

	// mux.HandleFunc("/auth/refresh", han.RegisterHandler)
	// mux.HandleFunc("/auth/signup", han.RegisterHandler)
	//
	// mux.Handle("/auth/me", protected(han.MeHandler))
	// mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello world")
	// })

	log.Println("Starting auth service on", cfg.Port)

	err := http.ListenAndServe("0.0.0.0:"+cfg.Port, middleware.Logger(mux))
	if err != nil {
		log.Fatal(err)
	}
}

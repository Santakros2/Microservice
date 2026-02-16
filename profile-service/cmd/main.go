package main

import (
	"log"
	"net/http"
	"profile-service/internal/db"
	"profile-service/internal/handler"
	"profile-service/internal/repository"
	"profile-service/internal/service"
)

func main() {

	database := db.NewMySQL()

	repo := &repository.ProfileRepo{DB: database}
	svc := &service.ProfileService{Repo: repo}
	h := &handler.ProfileHandler{Service: svc}

	http.HandleFunc("/profile", h.GetProfile)
	http.HandleFunc("/create-profile", h.CreateProfile)

	log.Println("Profile service running on :8002")
	http.ListenAndServe("0.0.0.0:8002", nil)
}

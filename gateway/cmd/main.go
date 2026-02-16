package main

import (
	"gateway/internal/config"
	"gateway/internal/router"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Gateway waiting for services...")
	time.Sleep(30 * time.Second) // TEMP HACK

	cfg := config.Load()

	handler := router.New(cfg)

	log.Println("Gateway running on :8000")
	http.ListenAndServe(":8000", handler)
}

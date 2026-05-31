package main

import (
	"log"
	"net/http"

	"go-api-starter/internal/config"
	"go-api-starter/internal/database"
	"go-api-starter/internal/modules/items"
)

func main() {
	cfg := config.Load()

	db := database.Connect(cfg)
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	items.RegisterRoutes(mux)

	address := ":" + cfg.AppPort

	log.Println("Ambiente:", cfg.AppEnv)
	log.Println("Servidor rodando em http://localhost" + address)

	err := http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}`))
}

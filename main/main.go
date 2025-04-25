package main

import (
	"its_wasnt_me/routes"
	"its_wasnt_me/services"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	// Verificar la API key antes de iniciar
	if err := services.VerifyAPIKey(); err != nil {
		log.Fatalf("Error al verificar la API key: %v", err)
	}
	router := routes.SetupRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

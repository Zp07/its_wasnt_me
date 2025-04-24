package main

import (
	"log"
	"net/http"
	"its_wasnt_me/routes"
)

func main() {
	router := routes.SetupRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
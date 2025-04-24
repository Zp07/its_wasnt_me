package routes

import (
	handlers "its_wasnt_me/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(" Bienvenido ! Usa /generar-meme para enviar tu deseo."))
	}).Methods("GET")

	// Endpoint principal
	r.HandleFunc("/generar-meme", handlers.GenerateMeme).Methods("POST")

	//
	r.PathPrefix("/genio_responde").Handler(
		http.StripPrefix("/genio_responde", http.FileServer(http.Dir("genio_responde"))),
	)
	return r
}

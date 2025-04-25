package routes

import (
	"its_wasnt_me/handlers"
	"its_wasnt_me/utils/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(" ¡ Bienvenido ! Usa /api/v1/generar-meme para enviar tu deseo."))
	}).Methods("GET")

	// subroute para la API
	api := r.PathPrefix("/api/v1").Subrouter()

	// Endpoint para generar memes con middleware de autenticación
	api.Handle("/generar-meme", middleware.AuthMiddleware(http.HandlerFunc(handlers.GenerateMeme))).Methods("POST")

	// Servir archivos estáticos manteniendo compatiblidad con HTML
	r.PathPrefix("/genio_responde").Handler(
		http.StripPrefix("/genio_responde", http.FileServer(http.Dir("genio_responde"))),
	)
	return r
}

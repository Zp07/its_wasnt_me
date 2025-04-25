package handlers

import (
	"encoding/json"
	"its_wasnt_me/services"
	"net/http"
	"strings"
)

// Json esperado para la peticion
type PromptRequest struct {
	Prompt string `json:"prompt"`
}

func GenerateMeme(w http.ResponseWriter, r *http.Request) {
	var req PromptRequest

	// Intenta decodificar el cuerpo del request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error al leer cuerpo de la solicitud"+err.Error(), http.StatusBadRequest)
		return
	}

	// Valida el prompt antes de procesar la solicitud
	trimmedPrompt := strings.TrimSpace(req.Prompt)
	if trimmedPrompt == "" {
		http.Error(w, "Prompt vacío, no permitido", http.StatusBadRequest)
		return
	}

	// Llamada al servicio para generar la imagen
	path, err := services.GenerateImageFromPrompt(req.Prompt)
	if err != nil {
		http.Error(w, "error al generar imagen"+err.Error(), http.StatusInternalServerError)
		return

	}

	// Configura el encabezado de la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")
	// Retorna la imagen generada como respuesta
	json.NewEncoder(w).Encode(map[string]string{"image_path": path})

}

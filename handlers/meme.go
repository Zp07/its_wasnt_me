package hanlders

import (
	"encoding/json"
	"net/http"
	"its_wasnt_me/services"
)

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

func GenerateMeme(w http.ResponseWriter, r *http.Request) {
	var req PromptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error al leer prompt", http.StatusBadRequest)
		return
	}

	// Llamada al servicio para generar la imagen
	path, err := services.GenerateImageFromPrompt(req.Prompt)
	if err != nil {
		http.Error(w, "error al generar imagen"+err.Error(), http.StatusInternalServerError)
		return

	}

	json.NewEncoder(w).Encode(map[string]string{"image_path": path})

}


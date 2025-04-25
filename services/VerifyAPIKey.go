package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func VerifyAPIKey() error {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OPENAI_API_KEY no está definida en el entorno")
	}

	// Crear una solicitud a la API de modelos para verificar la clave
	req, err := http.NewRequest("GET", "https://api.openai.com/v1/models", nil)
	if err != nil {
		return fmt.Errorf("error creando request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error al conectar con OpenAI: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("la API key no es válida o hay un problema con OpenAI: %s", string(bodyBytes))
	}

	fmt.Println("✅ API key de OpenAI verificada correctamente")
	return nil
}

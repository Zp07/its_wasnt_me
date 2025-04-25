package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// interfaz para mockear en tests
type OpenAIClient interface {
	GenerateImage(prompt string) (string, error)
}

// implementación de la interfaz por defecto
type OpenAIClientImpl struct{}

func newOpenAIClient() OpenAIClient {
	return &OpenAIClientImpl{}
}

func (c *OpenAIClientImpl) GenerateImage(prompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY no está definida en el entorno")
	}

	// Estructura del payload para la API de OpenAI de DALL-E
	payload := map[string]interface{}{
		"model":           "dall-e-3",
		"prompt":          prompt,
		"n":               1,
		"size":            "1024x1024",
		"response_format": "url",
		"style":           "vivid",
	}

	//logs
	fmt.Println("Enviando solicitud a OpenAI con payload:", payload) // Más detalles en el log

	fmt.Println("Prompt:", prompt) // Debug
	// fin

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error al convertir payload a JSON: %v", err)
	}

	//logs
	fmt.Println("Payload final:", string(body)) // 👈 Este print va justo aquí
	// fin

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error creando request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	//logs
	fmt.Println("Authorization: Bearer " + apiKey[:5] + "...")
	// fin

	client := &http.Client{Timeout: 30 * time.Second} // Timeout de 30 segundos
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error enviando request: %v", err)
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	bodyBytes, _ := io.ReadAll(resp.Body)
	responseBody := string(bodyBytes)

	if resp.StatusCode != 200 {
		// logs
		fmt.Println("OpenAI response code:", resp.StatusCode)
		fmt.Println("OpenAI response body:", responseBody)
		//fin

		return "", fmt.Errorf("OpenAI ERROR: %s", string(bodyBytes))
	}

	// Recrear el lector del cuerpo para poder decodificar el JSON
	respBodyReader := bytes.NewReader(bodyBytes)

	var result struct {
		Data []struct {
			URL string `json:"url"`
		} `json:"data"`
	}

	if err := json.NewDecoder(respBodyReader).Decode(&result); err != nil {
		return "", fmt.Errorf("error al decodificar respuesta: %v, respuesta: %s", err, responseBody)
	}

	if len(result.Data) == 0 {
		return "", fmt.Errorf("no se generó ninguna imagen")
	}

	return result.Data[0].URL, nil
}

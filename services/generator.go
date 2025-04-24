package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"its_wasnt_me/utils"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func GenerateImageFromPrompt(prompt string) (string, error) {
	// Llamar a la API de generación de imágenes con el prompt
	imageURL, err := callImageAPI(prompt)
	if err != nil {
		return "", err
	}

	// Crear la carpeta si no existe
	if err := utils.CreateFolderIfNotExist("genio_responde"); err != nil {
		return "", err
	}

	// Generar el nombre del archivo y la ruta
	filename := fmt.Sprintf("genio_responde_%d.png", time.Now().Unix())
	path := filepath.Join("genio_responde", filename)

	// Descargar la imagen y guardarla en la carpeta
	if err := downloadImage(imageURL, path); err != nil {
		return "", err
	}

	return "/" + path, nil
}

func callImageAPI(prompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY no está definida en el entorno")
	}

	payload := map[string]interface{}{
		"model":  "dall-e-2",
		"prompt": prompt,
		"n":      1,
		"size":   "1024x1024",
	}
	fmt.Println("Enviando solicitud a OpenAI con payload:", payload) // Más detalles en el log

	fmt.Println("Prompt:", prompt) // Debug

	body, _ := json.Marshal(payload)
	fmt.Println("Payload JSON:", string(body))

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error creando request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	fmt.Println("Authorization: Bearer " + apiKey[:5] + "...")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error enviando request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("OpenAI response code:", resp.StatusCode)
		fmt.Println("OpenAI response body:", string(bodyBytes))

		return "", fmt.Errorf("OpenAI ERROR: %s", string(bodyBytes))
	}

	var result struct {
		Data []struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Data) == 0 {
		return "", fmt.Errorf("no se generó ninguna imagen")
	}

	return result.Data[0].URL, nil
}
func downloadImage(url, path string) error {
	// Hacemos la descarga de la imagen desde la URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Crear el archivo donde se guardará la imagen
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copiar el contenido de la respuesta HTTP al archivo
	_, err = io.Copy(file, resp.Body)
	return err

}

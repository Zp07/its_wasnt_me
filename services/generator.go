package services

import (
	"fmt"
	"io"
	"its_wasnt_me/utils"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var openAI OpenAIClient = newOpenAIClient()

func SetupOpenAIClient(client OpenAIClient) {
	openAI = client
}

func GenerateImageFromPrompt(prompt string) (string, error) {
	// Llamar a la API de generación de imágenes con el prompt
	imageURL, err := openAI.GenerateImage(prompt)
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

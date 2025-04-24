package utils 

import (
	"os"
)

// Crea la carpeta si no existe
func CreateFolderIfNotExist(folderName string) error {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		return os.Mkdir(folderName, os.ModePerm)
	}

	return nil
}
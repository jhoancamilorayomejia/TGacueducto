// utils/env.go
package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv carga las variables de entorno desde el archivo .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno existentes")
	}
}

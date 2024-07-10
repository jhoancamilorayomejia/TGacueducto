package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB es la instancia de la base de datos que se usará en toda la aplicación
var DB *sql.DB

// ConnectDB establece y retorna una conexión a la base de datos PostgreSQL
func ConnectDB() (*sql.DB, error) {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Obtener las variables de entorno para la conexión a la base de datos
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construir el string de conexión
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Abrir una nueva conexión a la base de datos
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Probar la conexión a la base de datos
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Asignar la instancia de DB para usarla en otros paquetes
	DB = db

	return DB, nil
}

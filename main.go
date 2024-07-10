package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/jhoancamilorayomejia/TGacueducto/controllers"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer conn.Close()

	r := gin.Default()

	db.ConnectDB()

	r.GET("/api/admins", controllers.GetAdmins)

	// Cambia el puerto del backend a 8081
	err = r.Run(":8081")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

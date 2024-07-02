package main

import (
	"fmt"
	"log"
	"nombre_del_modulo/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to the database!")
}

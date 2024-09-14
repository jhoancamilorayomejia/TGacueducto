package main

import (
	"log"

	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	r.Use(cors.New(config))

	r.GET("/api/admins", controllers.GetAdmins)
	r.GET("/api/AllCompany", controllers.GetCompanies) // Ruta mostrar tabla de empresas
	r.GET("/api/customer", controllers.GetUsuarios)    // Ruta mostrar tabla de clientes/usuarios

	// Nueva ruta para obtener idcompany y name
	r.GET("/api/get-switch", controllers.GetSwitch)

	r.POST("/api/login", controllers.Login)

	// Ruta de acciones para admin
	r.DELETE("/api/admins/:idadmin", controllers.DeleteAdmin)
	r.PUT("/api/admins/:idadmin", controllers.UpdateAdmin)

	//Ruta de acciones para customer
	r.PUT("/api/customer/:idcustomer", controllers.UpdateCustomer)

	//Ruta de acciones para company
	r.PUT("/api/companies/:id", controllers.UpdateCompany)

	//r.DELETE("/api/companies/:id", controllers.DeleteCompany)

	//r.DELETE("/api/customer/:id", controllers.DeleteCustomer)

	err = r.Run(":8081")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

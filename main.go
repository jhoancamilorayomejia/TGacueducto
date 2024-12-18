package main

import (
	"log"

	//"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"github.com/jhoancamilorayomejia/TGacueducto/controllers"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file ")
	}

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer conn.Close()

	r := gin.Default()

	r.GET("/api/admins", controllers.GetAdmins)
	r.GET("/api/AllCompany", controllers.GetCompanies) // Ruta mostrar tabla de empresas
	r.GET("/api/customer", controllers.GetUsuarios)    // Ruta mostrar tabla de clientes/usuarios

	// Nueva ruta para obtener idcompany y name
	r.GET("/api/get-switch", controllers.GetSwitch)

	r.POST("/api/login", controllers.Login)

	// Ruta de acciones para admin
	r.POST("/api/register-user", controllers.CreateAdmin)
	r.POST("/api/register-company", controllers.CreateCompany)
	r.DELETE("/api/admins/:idadmin", controllers.DeleteAdmin)
	r.PUT("/api/admins/:idadmin", controllers.UpdateAdmin)
	r.GET("/api/company/:idcompany", controllers.GetCompanyByID) // Obtiene una empresa por ID
	//r.GET("/api/customer/:idcompany", controllers.GetCustomersByCompany)

	//Ruta de acciones para company
	r.PUT("/api/companies/:idcompany", controllers.UpdateCompany)
	r.DELETE("/api/companies/:idcompany", controllers.DeleteCompany)
	r.GET("/api/allcustomer/:idcompany", controllers.GetUsuariosPorIDCompany)
	r.POST("/api/registerCustomer", controllers.RegisterCustomer)
	r.POST("/api/facturas", controllers.CreateFacture)
	r.DELETE("/api/facture/:idfacture", controllers.DeleteFacture)
	r.DELETE("/api/customer/:idcustomer", controllers.DeleteCustomer)
	r.PUT("/api/facturas/:idfacture", controllers.UpdateFacture)

	r.GET("/api/facturas/:idcustomer", controllers.GetAllFactures)

	//Ruta de acciones para customer
	r.PUT("/api/customer/:idcustomer", controllers.UpdateCustomer)
	r.GET("/api/facturas", controllers.GetAllFactures)

	r.POST("/api/payment/preference", controllers.CreatePreference)

	// Nueva ruta para enviar el correo
	r.POST("/api/send-email", controllers.SendEmail)

	r.GET("/api/customers", controllers.GetUsuarios) // Ruta para obtener todos los clientes

	//cambio de clave para company
	// Rutas existentes
	// Ruta para obtener la información de la empresa (email y contraseña)
	r.GET("/api/company/get-email-password", controllers.GetEmailPassword)
	// Ruta para cambiar la contraseña
	r.POST("/api/company/change-password", controllers.ChangeCompanyPassword)

	//ruta para obtener la informacion del cliente (email y contraseña)
	r.GET("/api/customer/get-email-password", controllers.GetEmailPassword2)
	// Ruta para cambiar la contraseña
	r.POST("/api/customer/change-password", controllers.ChangeCustomerPassword)

	err = r.Run(":8081")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
	"github.com/jhoancamilorayomejia/TGacueducto/models"
	"golang.org/x/crypto/bcrypt"
)

// Company representa la estructura de la tabla company en la base de datos
type Company struct {
	IDcompany int    `json:"idcompany"`
	Nit       string `json:"nit"`
	Nombre    string `json:"Name"`
	Localidad string `json:"address"`
	Phone     string `json:"Phone"`
	Email     string `json:"Email"`
	Password  string `json:"password"`
	SecretKey string `json:"secret_key"`
}

// GetCompanies maneja la solicitud para obtener todos los registros de la tabla company
func GetCompanies(c *gin.Context) {
	var companies []Company

	rows, err := db.DB.Query("SELECT idcompany, nit, name, address, phone, email FROM company")
	if err != nil {
		log.Printf("Error querying company table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying company table"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var company Company
		err := rows.Scan(&company.IDcompany, &company.Nit, &company.Nombre, &company.Localidad, &company.Phone, &company.Email)
		if err != nil {
			log.Printf("Error scanning company row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning company row"})
			return
		}
		companies = append(companies, company)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over rows"})
		return
	}

	c.JSON(http.StatusOK, companies)
}

// UpdateCompany maneja la solicitud para actualizar un registro de la tabla company
func UpdateCompany(c *gin.Context) {
	var company Company

	// Bind the JSON payload to the company struct
	if err := c.BindJSON(&company); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Ensure that the IDcompany field is provided
	if company.IDcompany == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	// Update the company record in the database
	query := `
		UPDATE company
		SET nit = $1, name = $2, address = $3, phone = $4, email = $5
		WHERE idcompany = $6
	`
	_, err := db.DB.Exec(query, company.Nit, company.Nombre, company.Localidad, company.Phone, company.Email, company.IDcompany)
	if err != nil {
		log.Printf("Error updating company: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully"})
}

// CreateCompany maneja la solicitud para crear una nueva empresa
func CreateCompany(c *gin.Context) {
	var company models.Company // Usamos la estructura del archivo models

	// Vincular los datos del JSON al struct `Company`
	if err := c.BindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Encriptar la contraseña antes de guardarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(company.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar la contraseña"})
		return
	}
	company.Password = string(hashedPassword)

	// Inserta la nueva empresa en la tabla company
	_, err = db.DB.Exec(`
		INSERT INTO company (nit, name, address, phone, email, password, secret_key)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, company.IDnit, company.Nombre, company.Address, company.Phone, company.Email, company.Password, company.SecretKey)
	if err != nil {
		log.Printf("Error inserting company: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar la Empresa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empresa registrada con éxito"})
}

/*
   Campo aparte

*/
// Es para poder que aparezca todas las entidades y seleccionar al que se quiere registrar al cleinte
// UserSwitch representa la estructura simplificada para la selección de entidad
type UserSwitch struct {
	IDcompany int    `json:"idcompany"`
	Name      string `json:"name"`
}

// GetSwitch maneja la solicitud para obtener los idcompany y name de la tabla company
func GetSwitch(c *gin.Context) {
	var userSwitches []UserSwitch

	rows, err := db.DB.Query("SELECT idcompany, name FROM company")
	if err != nil {
		log.Printf("Error querying company table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying company table"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userSwitch UserSwitch
		err := rows.Scan(&userSwitch.IDcompany, &userSwitch.Name)
		if err != nil {
			log.Printf("Error scanning company row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning company row"})
			return
		}
		userSwitches = append(userSwitches, userSwitch)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over rows"})
		return
	}

	c.JSON(http.StatusOK, userSwitches)
}

// DeleteCompany maneja la solicitud para eliminar una empresa por su ID
func DeleteCompany(c *gin.Context) {
	idcompany := c.Param("idcompany")

	// Ejecutar la consulta SQL para eliminar el registro
	_, err := db.DB.Exec("DELETE FROM company WHERE idcompany = $1", idcompany)
	if err != nil {
		log.Printf("Error deleting company: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la Entidad."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entidad eliminada exitosamente."})
}

// ensayo para Info company (desde admin)
// Handler para obtener los detalles de una empresa por ID
func GetCompanyByID(c *gin.Context) {
	idcompany := c.Param("idcompany")
	var company Company
	err := db.DB.QueryRow("SELECT nit, name, address, phone, email FROM company WHERE idcompany = $1", idcompany).
		Scan(&company.Nit, &company.Nombre, &company.Localidad, &company.Phone, &company.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching company details"})
		return
	}

	c.JSON(http.StatusOK, company)
}

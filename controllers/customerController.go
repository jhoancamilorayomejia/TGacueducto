package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
)

// Usuario representa la estructura de la tabla usuarios en la base de datos
type Customer struct {
	IDcustomer int    `json:"idcustomer"`
	IDcompany  int    `json:"idcompany"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Localidad  string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

// GetUsuarios maneja la solicitud para obtener todos los registros de la tabla usuarios
func GetUsuarios(c *gin.Context) {
	var usuarios []Customer

	rows, err := db.DB.Query("SELECT idcustomer, idcompany, name, last_name, address, phone, email FROM customer")
	if err != nil {
		log.Printf("Error querying usuario table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying usuario table"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var usuario Customer
		err := rows.Scan(&usuario.IDcustomer, &usuario.IDcompany, &usuario.Name, &usuario.LastName, &usuario.Localidad, &usuario.Phone, &usuario.Email)
		if err != nil {
			log.Printf("Error scanning usuario row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning usuario row"})
			return
		}
		usuarios = append(usuarios, usuario)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over rows"})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

// UpdateCustomer maneja la solicitud para actualizar un cliente en la base de datos
func UpdateCustomer(c *gin.Context) {
	var customer Customer
	if err := c.BindJSON(&customer); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate that IDcustomer is present
	if customer.IDcustomer == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDcustomer is required"})
		return
	}

	// Prepare the SQL statement for updating
	stmt, err := db.DB.Prepare(`
		UPDATE customer
		SET idcompany = $1, name = $2, last_name = $3, address = $4, phone = $5, email = $6
		WHERE idcustomer = $7
	`)
	if err != nil {
		log.Printf("Error preparing update statement: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparing update statement"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(customer.IDcompany, customer.Name, customer.LastName, customer.Localidad, customer.Phone, customer.Email, customer.IDcustomer)
	if err != nil {
		log.Printf("Error executing update statement: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error executing update statement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}

// RegisterUsuario maneja la solicitud para registrar un nuevo usuario
/* func RegisterUsuario(c *gin.Context) {
	var user Customer

	// Parsear el cuerpo de la solicitud
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Insertar en la base de datos
	query := `INSERT INTO customer(idcompany, name, last_name, address, phone, email, password)
	          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING idcustomer`

	err = db.DB.QueryRow(query, user.IDcompany, user.Name, user.LastName, user.Localidad, user.Phone, user.Email, user.Password).Scan(&user.IDcustomer)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Responder con el nuevo usuario creado
	c.JSON(http.StatusOK, user)
}
*/

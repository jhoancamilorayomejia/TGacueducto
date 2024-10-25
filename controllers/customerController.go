package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
	"github.com/jhoancamilorayomejia/TGacueducto/models"
	"golang.org/x/crypto/bcrypt"
)

// Usuario representa la estructura de la tabla usuarios en la base de datos
type Customer struct {
	IDcustomer int    `json:"idcustomer"`
	IDcompany  int    `json:"idcompany"`
	Cedula     string `json:"cedula"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Localidad  string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	SecretKey  string `json:"secret_key"`
}

// GetUsuarios maneja la solicitud para obtener los registros de la tabla usuarios filtrados por idcompany
func GetUsuarios(c *gin.Context) {
	var usuarios []Customer

	// Obtén el idcompany del parámetro de consulta
	idcompany := c.Query("idcompany")

	// Prepara la consulta SQL
	query := "SELECT idcustomer, idcompany, cedula, name, last_name, address, phone, email FROM customer WHERE idcompany = $1"
	rows, err := db.DB.Query(query, idcompany)
	if err != nil {
		log.Printf("Error querying usuario table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying usuario table"})
		return
	}
	defer rows.Close()

	// Itera sobre los resultados de la consulta
	for rows.Next() {
		var usuario Customer
		err := rows.Scan(&usuario.IDcustomer, &usuario.IDcompany, &usuario.Cedula, &usuario.Name, &usuario.LastName, &usuario.Localidad, &usuario.Phone, &usuario.Email)
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

type Customer2 struct {
	IDcustomer int    `json:"idcustomer"`
	IDcompany  int    `json:"idcompany"`
	Cedula     string `json:"cedula"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Localidad  string `json:"address"` // Localidad en lugar de address para consistencia con JSON
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

// GetUsuariosPorIDCompany maneja la solicitud para obtener los registros de la tabla customer relacionados con un idcompany específico
func GetUsuariosPorIDCompany(c *gin.Context) {
	idCompany := c.Param("idcompany") // Obtiene el ID de la empresa desde los parámetros de la ruta
	var allusuarios []Customer2

	// Realiza la consulta a la base de datos
	rows, err := db.DB.Query("SELECT idcustomer, idcompany, cedula, name, last_name, address, phone, email FROM customer WHERE idcompany = $1", idCompany)
	if err != nil {
		log.Printf("Error querying customer table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying customer table"})
		return
	}
	defer rows.Close()

	// Itera sobre los resultados
	for rows.Next() {
		var usuario Customer2
		err := rows.Scan(&usuario.IDcustomer, &usuario.IDcompany, &usuario.Cedula, &usuario.Name, &usuario.LastName, &usuario.Localidad, &usuario.Phone, &usuario.Email)
		if err != nil {
			log.Printf("Error scanning customer row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning customer row"})
			return
		}
		allusuarios = append(allusuarios, usuario)
	}

	// Verifica si hubo un error al iterar sobre los rows
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over rows"})
		return
	}

	// Responde con los datos obtenidos
	c.JSON(http.StatusOK, gin.H{"usuarios": allusuarios})
}

// RegisterCustomer maneja la solicitud para registrar un nuevo cliente
func RegisterCustomer(c *gin.Context) {
	var client Customer
	if err := c.BindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si ya existe un cliente con el mismo correo electrónico
	var emailExists bool
	err := db.DB.QueryRow(`SELECT EXISTS (SELECT 1 FROM customer WHERE email = $1)`, client.Email).Scan(&emailExists)
	if err != nil {
		log.Printf("Error verificando correo electrónico: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar el correo electrónico"})
		return
	}
	if emailExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El correo electrónico ya está registrado"})
		return
	}

	// Verificar si ya existe un cliente con la misma cédula
	var cedulaExists bool
	err = db.DB.QueryRow(`SELECT EXISTS (SELECT 1 FROM customer WHERE cedula = $1)`, client.Cedula).Scan(&cedulaExists)
	if err != nil {
		log.Printf("Error verificando cédula: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar la cédula"})
		return
	}
	if cedulaExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La cédula ya está registrada"})
		return
	}

	// Generar un secret_key único
	secretKey := models.GenerateRandomHexString(32) // Genera una cadena hex con 32 bytes
	client.SecretKey = secretKey

	// Verificar si la compañía existe
	var companyExists bool
	err = db.DB.QueryRow(`SELECT EXISTS (SELECT 1 FROM company WHERE idcompany = $1)`, client.IDcompany).Scan(&companyExists)
	if err != nil || !companyExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La compañía seleccionada no existe"})
		return
	}

	// Encriptar la contraseña antes de guardarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(client.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar la contraseña"})
		return
	}
	client.Password = string(hashedPassword)

	// Insertar el cliente en la base de datos
	_, err = db.DB.Exec(`
		INSERT INTO customer (idcompany, cedula, name, last_name, address, phone, email, password, secret_key)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, client.IDcompany, client.Cedula, client.Name, client.LastName, client.Localidad, client.Phone, client.Email, client.Password, client.SecretKey)
	if err != nil {
		log.Printf("Error inserting customer: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar el cliente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente registrado con éxito"})
}

// DeleteFacture maneja la solicitud para eliminar una factura por su ID
func DeleteCustomer(c *gin.Context) {
	idcustomer := c.Param("idcustomer") // Obtiene el ID del cliente desde la URL

	// Ejecutar la consulta SQL para eliminar el registro
	_, err := db.DB.Exec("DELETE FROM customer WHERE idcustomer = $1", idcustomer)
	if err != nil {
		log.Printf("Error al eliminar el cliente: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el cliente."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado exitosamente."})
}

package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
	"github.com/jhoancamilorayomejia/TGacueducto/models"
	"golang.org/x/crypto/bcrypt"
)

// Admin representa la estructura de la tabla admin en la base de datos
type Admin struct {
	IDadmin   int    `json:"idadmin"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	SecretKey string `json:"secret_key"`
	//Tipouser  string `json:"tipouser"`
}

// GetAdmins maneja la solicitud para obtener todos los registros de la tabla admin
func GetAdmins(c *gin.Context) {
	var admins []Admin

	// Consulta sin CASE
	rows, err := db.DB.Query(`
		SELECT idadmin, nombre, apellido, email
		FROM admin
	`)
	if err != nil {
		log.Printf("Error querying admin table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener la lista de administradores."})
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	// Iterar sobre los resultados
	for rows.Next() {
		var admin Admin
		err := rows.Scan(&admin.IDadmin, &admin.Nombre, &admin.Apellido, &admin.Email)
		if err != nil {
			log.Printf("Error scanning admin row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los datos de los administradores."})
			return
		}

		// Lógica para cambiar Tipouser a "Administrador" si es 1
		//if admin.Tipouser == "1" {
		//	admin.Tipouser = "Administrador"
		//}

		admins = append(admins, admin)
	}

	// Verificar errores durante la iteración
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Hubo un problema al leer los datos de los administradores."})
		return
	}

	// Responder con la lista de administradores
	c.JSON(http.StatusOK, admins)
}

// UpdateAdmin maneja la solicitud para actualizar un administrador
func UpdateAdmin(c *gin.Context) {
	idadmin := c.Param("idadmin")

	var admin Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	_, err := db.DB.Exec(`
        UPDATE admin
        SET nombre = $1, apellido = $2, email = $3
        WHERE idadmin = $4
    `, admin.Nombre, admin.Apellido, admin.Email, idadmin)
	if err != nil {
		log.Printf("Error updating admin: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el administrador."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Administrador actualizado con éxito"})
}

// DeleteAdmin maneja la solicitud para eliminar un administrador por su ID
func DeleteAdmin(c *gin.Context) {
	idadmin := c.Param("idadmin")

	// Ejecutar la consulta SQL para eliminar el registro
	_, err := db.DB.Exec("DELETE FROM admin WHERE idadmin = $1", idadmin)
	if err != nil {
		log.Printf("Error deleting admin: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el administrador."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Administrador eliminado exitosamente."})
}

// CreateAdmin maneja la solicitud para crear un nuevo administrador
func CreateAdmin(c *gin.Context) {
	var admin Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Generar un secret_key único
	secretKey := models.GenerateRandomHexString(32) // Genera una cadena hex con 32 bytes
	admin.SecretKey = secretKey

	// Encriptar la contraseña antes de guardarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar la contraseña"})
		return
	}
	admin.Password = string(hashedPassword)

	_, err = db.DB.Exec(`
		INSERT INTO admin (nombre, apellido, email, password, secret_key)
		VALUES ($1, $2, $3, $4, $5)
	`, admin.Nombre, admin.Apellido, admin.Email, admin.Password, admin.SecretKey)
	if err != nil {
		log.Printf("Error inserting admin: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar el administrador"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Administrador registrado con éxito", "secret_key": secretKey})
}

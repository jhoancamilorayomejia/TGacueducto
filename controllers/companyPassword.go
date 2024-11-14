package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
	"golang.org/x/crypto/bcrypt"
)

// ChangePasswordRequest estructura para manejar el cambio de contraseña
type ChangePasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

// CompanyPassword representa la estructura de la tabla company en la base de datos
type CompanyPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetEmailPassword maneja la solicitud para obtener el email y la contraseña de una empresa específica
func GetEmailPassword(c *gin.Context) {
	userEmail := c.Query("email")

	var company CompanyPassword
	// Realiza la consulta para obtener el email y la contraseña de la empresa
	err := db.DB.QueryRow("SELECT email, password FROM company WHERE email = $1", userEmail).Scan(&company.Email, &company.Password)
	if err != nil {
		log.Printf("Error querying company by email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying company by email"})
		return
	}

	c.JSON(http.StatusOK, company)
}

// ChangeCompanyPassword maneja la solicitud para actualizar la contraseña de la empresa
func ChangeCompanyPassword(c *gin.Context) {
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Encripta la nueva contraseña usando bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error al encriptar la contraseña: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encriptar la contraseña"})
		return
	}

	// Actualiza la contraseña en la base de datos
	_, err = db.DB.Exec("UPDATE company SET password = $1 WHERE email = $2", hashedPassword, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Compañía no encontrada"})
		} else {
			log.Printf("Error al actualizar la contraseña: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la contraseña"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contraseña actualizada exitosamente"})
}

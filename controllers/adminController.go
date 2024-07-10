package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
)

// Admin representa la estructura de la tabla admin en la base de datos
type Admin struct {
	IDadmin  int    `json:"idadmin"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Tipouser string `json:"tipouser"`
}

// GetAdmins maneja la solicitud para obtener todos los registros de la tabla admin
func GetAdmins(c *gin.Context) {
	var admins []Admin

	rows, err := db.DB.Query("SELECT idadmin, nombre, apellido, email, password, tipouser FROM admin")
	if err != nil {
		log.Printf("Error querying admin table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying admin table"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var admin Admin
		err := rows.Scan(&admin.IDadmin, &admin.Nombre, &admin.Apellido, &admin.Email, &admin.Password, &admin.Tipouser)
		if err != nil {
			log.Printf("Error scanning admin row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning admin row"})
			return
		}
		admins = append(admins, admin)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over rows"})
		return
	}

	c.JSON(http.StatusOK, admins)
}

package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
)

// Usuario representa la estructura de la tabla company en la base de datos
type Usuario struct {
	IDusuario int    `json:"idcompany"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Localidad string `json:"address"`
	Phone     string `json:"Phone"`
	Email     string `json:"Email"`
}

// GetUsuarios maneja la solicitud para obtener todos los registros de la tabla company
func GetUsuarios(c *gin.Context) {
	var usuarios []Usuario

	rows, err := db.DB.Query("SELECT iduser, name, last_name, address, phone, email FROM usuarios")
	if err != nil {
		log.Printf("Error querying usuario table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying company table"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var usuario Usuario
		err := rows.Scan(&usuario.IDusuario, &usuario.Name, &usuario.LastName, &usuario.Localidad, &usuario.Phone, &usuario.Email)
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

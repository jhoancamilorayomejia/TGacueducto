// controllers/adminController.go
package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
)

type Admin struct {
	ID       int    `json:"IDadmin"`
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Email    string `json:"email"`
	Password string `json:"password"`
	TipoUser string `json:"TipoUser"`
}

func GetAdmins(c *gin.Context) {
	var admins []Admin

	rows, err := db.DB.Query("SELECT IDadmin, Nombre, Apellido, email, password, TipoUser FROM admin")
	if err != nil {
		log.Fatalf("Error querying admin table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying admin table"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var admin Admin
		err := rows.Scan(&admin.ID, &admin.Nombre, &admin.Apellido, &admin.Email, &admin.Password, &admin.TipoUser)
		if err != nil {
			log.Fatalf("Error scanning admin row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning admin row"})
			return
		}
		admins = append(admins, admin)
	}

	c.JSON(http.StatusOK, admins)
}

package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
)

// Company representa la estructura de la tabla company en la base de datos
type Company struct {
	IDcompany int    `json:"idcompany"`
	Nit       string `json:"nit"`
	Nombre    string `json:"Name"`
	Localidad string `json:"address"`
	Phone     string `json:"Phone"`
	Email     string `json:"Email"`
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

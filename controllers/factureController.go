package controllers

import (
	"net/http"

	"github.com/jhoancamilorayomejia/TGacueducto/db"
	"github.com/jhoancamilorayomejia/TGacueducto/models"

	"github.com/gin-gonic/gin"
)

// Facture estructura para la factura
type Facture struct {
	IDfacture     int    `json:"idfacture"`
	IDcompany     int    `json:"idcompany"`
	IDcustomer    int    `json:"idcustomer"`
	FactureNumber string `json:"facturenumber"`
	DateCreation  string `json:"datecreation"`
	DatePayment   string `json:"datepayment"`
	TotalPay      string `json:"totalpay"`
}

// GetAllFactures obtiene todas las facturas de la base de datos
func GetAllFactures(c *gin.Context) {
	var facturas []Facture

	rows, err := db.DB.Query("SELECT idfacture, idcompany, idcustomer, facturenumber, datecreation, datepayment, totalpay FROM facture")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching facturas"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var factura Facture
		err := rows.Scan(&factura.IDfacture, &factura.IDcompany, &factura.IDcustomer, &factura.FactureNumber, &factura.DateCreation, &factura.DatePayment, &factura.TotalPay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning factura"})
			return
		}
		facturas = append(facturas, factura)
	}

	c.JSON(http.StatusOK, facturas)
}

// GetFacturesByCustomerID obtiene las facturas por ID de cliente
/*func GetFacturesByCustomerID(c *gin.Context) {
	idCustomer := c.Param("idcustomer") // Obtener el ID del cliente desde la URL
	var facturas []Facture

	rows, err := db.DB.Query("SELECT idfacture, idcompany, idcustomer, facturenumber, datecreation, datepayment, totalpay FROM facture WHERE idcustomer = ?", idCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching facturas"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var factura Facture
		err := rows.Scan(&factura.IDfacture, &factura.IDcompany, &factura.IDcustomer, &factura.FactureNumber, &factura.DateCreation, &factura.DatePayment, &factura.TotalPay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning factura"})
			return
		}
		facturas = append(facturas, factura)
	}

	c.JSON(http.StatusOK, facturas)
} */

// CreateFacture maneja la creación de una nueva factura
func CreateFacture(c *gin.Context) {
	var input Facture

	// Validar los datos recibidos desde el frontend
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear una nueva instancia del modelo Facture (ajusta según el nombre de tu modelo)
	facture := models.Facture{
		IDcompany:     input.IDcompany,
		IDcustomer:    input.IDcustomer,
		FactureNumber: input.FactureNumber,
		DateCreation:  input.DateCreation,
		DatePayment:   input.DatePayment,
		TotalPay:      input.TotalPay,
	}

	// Insertar la nueva factura en la base de datos
	query := `INSERT INTO facture (idcompany, idcustomer, facturenumber, datecreation, datepayment, totalpay)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING idfacture`

	err := db.DB.QueryRow(query, facture.IDcompany, facture.IDcustomer, facture.FactureNumber, facture.DateCreation, facture.DatePayment, facture.TotalPay).Scan(&facture.IDfacture)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la factura"})
		return
	}

	// Retornar la nueva factura creada
	c.JSON(http.StatusCreated, gin.H{
		"message": "Factura creada exitosamente",
		"facture": facture,
	})
}

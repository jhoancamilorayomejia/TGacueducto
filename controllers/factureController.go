package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/db"
)

// Facture representa una factura
type Facture struct {
	IDfacture     int       `json:"idfacture"`
	IDcompany     int       `json:"idcompany"`
	IDcustomer    int       `json:"idcustomer"`
	FactureNumber string    `json:"facturenumber"`
	DateCreation  time.Time `json:"datecreation"`
	DatePayment   time.Time `json:"datepayment"` // Cambiado a DueDate
	TotalPay      float64   `json:"totalpay"`    // Monto total a pagar
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
func GetFacturesByCustomerID(c *gin.Context) {
	idCustomer := c.Param("idcustomer") // Obtener el ID del cliente desde la URL
	var facturas []Facture

	// Cambia la consulta SQL para filtrar por ID de cliente
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
}

//Para las descarga de los PDF

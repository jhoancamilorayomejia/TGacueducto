package models

import "time"

// Facture representa una factura
type Facture struct {
	IDfacture     int       `json:"idfacture"`
	IDcompany     int       `json:"idcompany"`
	IDcustomer    int       `json:"idcustomer"`
	FactureNumber string    `json:"facturenumber"`
	DateCreation  time.Time `json:"datecreation"`
	DatePayment   time.Time `json:"datepayment"`
	TotalPay      float64   `json:"totalpay"` // Cambiado a float64
}

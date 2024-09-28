package models

// Facture representa una factura
type Facture struct {
	IDfacture     int    `json:"idfacture"`
	IDcompany     int    `json:"idcompany"`
	IDcustomer    int    `json:"idcustomer"`
	FactureNumber string `json:"facturenumber"`
	DateCreation  string `json:"datecreation"` // Cambiado a string
	DatePayment   string `json:"datepayment"`  // Cambiado a string
	TotalPay      string `json:"totalpay"`     // Monto total a pagar
}

package models

type Facture struct {
	IDfacture     int    `json:"idfacture"`
	IDcompany     int    `json:"idcompany"`
	IDcustomer    int    `json:"idcustomer"`
	FactureNumber string `json:"facturenumber"`
	DateCreation  string `json:"datecreation"`
	DatePayment   string `json:"datepayment"`
	MeterBefore   string `json:"meterbefore"`
	MeterAfter    string `json:"meterafter"`
	Consumer      string `json:"consumer"`
	CodFacture    string `json:"codfacture"`
	DateLimit     string `json:"datelimit"`
	StatusFacture string `json:"statusfacture"`
	TotalPay      string `json:"totalpay"`
	CustomerEmail string `json:"customerEmail"` // Nuevo campo agregado
}

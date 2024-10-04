package models

// Admin representa la estructura de la tabla admin en la base de datos. Para luego iniciar session
type Company struct {
	IDcompany int    `json:"idcompany"`
	IDnit     string `json:"nit"`
	Nombre    string `json:"Name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	SecretKey string `json:"secret_key"`
}

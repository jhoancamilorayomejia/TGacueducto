package models

// User representa la estructura de la tabla admin en la base de datos. Para luego iniciar session
type Customer struct {
	IDcustomer int    `json:"idcustomer"`
	Cedula     string `json:"cedula"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	SecretKey  string `json:"secret_key"`
}

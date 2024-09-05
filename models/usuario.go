package models

// User representa la estructura de la tabla admin en la base de datos. Para luego iniciar session
type Usuario struct {
	IDusuario int    `json:"idusuario"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Tipouser  string `json:"idtipouser"`
}

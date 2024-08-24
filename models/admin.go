package models

// Admin representa la estructura de la tabla company en la base de datos
type Admin struct {
	IDadmin   int    `json:"idadmin"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Tipouser  string `json:"tipouser"`
	SecretKey string `json:"secret_key"`
}

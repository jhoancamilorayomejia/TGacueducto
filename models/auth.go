package models

type Auth struct {
	ID       int    `json:"id"`
	Cedula   string `json:"cedula"`
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	UserType string `json:"userType"`
}

package models

type Admin struct {
	ID    int    `db:"IDadmin"`
	Name  string `db:"Nombre"`
	Email string `db:"email"`
}

package storage

import (
	"github.com/jhoancamilorayomejia/TGacueducto/db"
	"github.com/jhoancamilorayomejia/TGacueducto/models"
)

// para company tendria que crear otra funcion de aqui.
// GetAdminByEmail
func GetAdminByEmail(email string) (*models.Admin, error) {
	rows, err := db.DB.Query(`SELECT idadmin, nombre, password, secret_key FROM admin WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	admin := &models.Admin{}

	for rows.Next() {
		err := rows.Scan(&admin.IDadmin, &admin.Nombre, &admin.Password, &admin.SecretKey)
		if err != nil {
			return nil, err
		}
	}

	return admin, nil
}

// GetAdminByID
func GetAdminByID(id int, email string) (*models.Admin, error) {
	rows, err := db.DB.Query(`SELECT secret_key FROM admin WHERE idadmin = $1 AND email = $2`, id, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	admin := &models.Admin{}

	for rows.Next() {
		err := rows.Scan(&admin.SecretKey)
		if err != nil {
			return nil, err
		}
	}

	return admin, nil
}

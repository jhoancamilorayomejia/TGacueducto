package storage

import (
	"github.com/jhoancamilorayomejia/TGacueducto/db"
	"github.com/jhoancamilorayomejia/TGacueducto/models"
)

// para company tendria que crear otra funcion de aqui.
// GetAdminByEmail
func GetCompanyByEmail(email string) (*models.Company, error) {
	rows, err := db.DB.Query(`SELECT idcompany, nit, name, password, secret_key FROM company WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	company := &models.Company{}

	for rows.Next() {
		err := rows.Scan(&company.IDcompany, &company.IDnit, &company.Nombre, &company.Password, &company.SecretKey)
		if err != nil {
			return nil, err
		}
	}

	return company, nil
}

func GetCompanyByID(id int, email string) (*models.Company, error) {
	rows, err := db.DB.Query(`SELECT secret_key FROM company WHERE idcompany = $1 AND email = $2`, id, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	company := &models.Company{}

	for rows.Next() {
		err := rows.Scan(&company.SecretKey)
		if err != nil {
			return nil, err
		}
	}

	return company, nil
}

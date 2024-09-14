package storage

import (
	"fmt"

	"github.com/jhoancamilorayomejia/TGacueducto/db"
	"github.com/jhoancamilorayomejia/TGacueducto/models"
)

// GetCustomerByEmail obtiene un cliente a partir de su correo electrónico
func GetCustomerByEmail(email string) (*models.Customer, error) {
	rows, err := db.DB.Query(`SELECT idcustomer, password, secret_key FROM customer WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customer := &models.Customer{}

	for rows.Next() {
		err := rows.Scan(&customer.IDcustomer, &customer.Password, &customer.SecretKey)
		if err != nil {
			return nil, err
		}
	}

	return customer, nil
}

// GetCustomerByID obtiene un cliente a partir de su ID y correo electrónico
func GetCustomerByID(id int, email string) (*models.Customer, error) {
	rows, err := db.DB.Query(`SELECT secret_key FROM customer WHERE idcustomer = $1 AND email = $2`, id, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customer := &models.Customer{}

	for rows.Next() {
		err := rows.Scan(&customer.SecretKey)
		if err != nil {
			return nil, err
		}
	}

	return customer, nil
}

// DeleteCustomerByID elimina un cliente a partir de su ID
func DeleteCustomerByID(id int) error {
	query := `DELETE FROM customer WHERE idcustomer = $1`
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no customer found with ID %d", id)
	}

	return nil
}

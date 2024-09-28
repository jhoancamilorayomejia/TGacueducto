// src/controllers/paymentController.go
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jhoancamilorayomejia/TGacueducto/services" // Asegúrate de que el path sea correcto
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Price string `json:"price"`
	}

	// Decodificar el JSON recibido
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear el pago en PayPal
	paymentID, err := services.CreatePayPalPayment(req.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar el ID del pago como respuesta
	response := map[string]string{"paymentID": paymentID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

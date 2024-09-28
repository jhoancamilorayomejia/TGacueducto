// src/services/paypalService.go
package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	payPalAPI    = "https://api.sandbox.paypal.com/v1/payments/payment" // Usa la URL de producción para entorno real
	clientID     = "YOUR_PAYPAL_CLIENT_ID"
	clientSecret = "YOUR_PAYPAL_CLIENT_SECRET"
)

type PayPalPayment struct {
	Intent string `json:"intent"`
	Payer  struct {
		PaymentMethod string `json:"payment_method"`
	} `json:"payer"`
	Transactions []struct {
		Amount struct {
			Total    string `json:"total"`
			Currency string `json:"currency"`
		} `json:"amount"`
		Description string `json:"description"`
	} `json:"transactions"`
	RedirectUrls struct {
		ReturnURL string `json:"return_url"`
		CancelURL string `json:"cancel_url"`
	} `json:"redirect_urls"`
}

type PayPalResponse struct {
	ID string `json:"id"`
}

func CreatePayPalPayment(price string) (string, error) {
	payment := PayPalPayment{
		Intent: "sale",
	}
	payment.Payer.PaymentMethod = "paypal"
	payment.Transactions = []struct {
		Amount struct {
			Total    string `json:"total"`
			Currency string `json:"currency"`
		} `json:"amount"`
		Description string `json:"description"`
	}{{
		Amount: struct {
			Total    string `json:"total"`
			Currency string `json:"currency"`
		}{
			Total:    price,
			Currency: "USD", // Cambia según tu moneda
		},
		Description: "Pago por servicio",
	}}
	payment.RedirectUrls.ReturnURL = "http://localhost:8080/success" // Cambia según tu URL de éxito
	payment.RedirectUrls.CancelURL = "http://localhost:8080/cancel"  // Cambia según tu URL de cancelación

	// Crear la solicitud de pago
	body, err := json.Marshal(payment)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", payPalAPI, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response PayPalResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.ID, nil
}

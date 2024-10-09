package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

// CreatePreference genera la preferencia de pago
func CreatePreference(c *gin.Context) {
	// Cargar datos enviados desde el frontend
	var requestData struct {
		Amount     float64 `json:"amount"`
		FactureCod string  `json:"FactureCod"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Configurar SDK de Mercado Pago
	cfg, err := config.New("APP_USR-2877031606482659-100419-b3d37070db91f99a4971fe85f8b0171e-2007177437") // Reemplaza con tu Access Token de Mercado Pago
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Configuration error"})
		return
	}

	client := preference.NewClient(cfg)

	// Crear preferencia con los datos recibidos
	request := preference.Request{
		Items: []preference.ItemRequest{
			{
				Title:     fmt.Sprintf("Pago de Servicio Acueducto - Factura: %s", requestData.FactureCod),
				Quantity:  1,
				UnitPrice: requestData.Amount,
				//Description: fmt.Sprintf("Pago de Servicio - Factura: %s", requestData.FactureCod),
			},
		},
		BackURLs: &preference.BackURLsRequest{
			Success: "http://localhost:8080/success",
			Failure: "http://localhost:8080/failure",
			Pending: "http://localhost:8080/pending",
		},
	}

	// Enviar preferencia a Mercado Pago
	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating preference"})
		return
	}

	// Devolver el Preference ID al frontend
	c.JSON(http.StatusOK, gin.H{"preferenceId": resource.ID})
}

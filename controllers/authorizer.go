package controllers

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/TGacueducto/models"
	"github.com/jhoancamilorayomejia/TGacueducto/security"
	"github.com/jhoancamilorayomejia/TGacueducto/storage"
)

// Funcion Login
func Login(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error getting body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	auth, err := models.Parser[models.Auth](body)
	if err != nil {
		log.Printf("Error parsing body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var apiKey *storage.APIKey
	var apiKeyResponse *storage.APIKeyResponse
	var userType string
	var userID int        // Para almacenar el ID del usuario
	var userName string   // Para almacenar el nombre del usuario
	var userNit string    //Para almacenar el Nit en company
	var userCedula string //Para almacenar la cedula en customer

	switch auth.UserType {
	case "admin":
		admin, err := storage.GetAdminByEmail(auth.Email)
		if err != nil {
			log.Printf("Error getting admin: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		err = security.VerifyPassword(admin.Password, auth.Password)
		if err != nil {
			log.Printf("Error verifying password: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
			return
		}

		apiKey, err = storage.NewAPIKey(admin, "", auth.Email)
		if err != nil {
			log.Printf("Error creating api key: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		userType = "admin"
		userID = admin.IDadmin  // Suponiendo que `admin` tiene un campo ID
		userName = admin.Nombre // Suponiendo que `admin` tiene un campo Nombre

	case "company":
		company, err := storage.GetCompanyByEmail(auth.Email)
		if err != nil {
			log.Printf("Error getting company: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		err = security.VerifyPassword(company.Password, auth.Password)
		if err != nil {
			log.Printf("Error verifying password: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
			return
		}

		apiKey, err = storage.NewCompanyAPIKey(company, "", auth.Email)
		if err != nil {
			log.Printf("Error creating api key: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		userType = "company"
		userID = company.IDcompany // Suponiendo que `company` tiene un campo ID
		userName = company.Nombre  // Suponiendo que `company` tiene un campo Nombre
		userNit = company.IDnit

	case "customer":
		customer, err := storage.GetCustomerByEmail(auth.Email)
		if err != nil {
			log.Printf("Error getting customer: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		err = security.VerifyPassword(customer.Password, auth.Password)
		if err != nil {
			log.Printf("Error verifying password: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
			return
		}

		apiKey, err = storage.NewCustomerAPIKey(customer, "", auth.Email)
		if err != nil {
			log.Printf("Error creating api key: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		userType = "customer"
		userID = customer.IDcustomer // Suponiendo que `customer` tiene un campo ID
		userName = customer.Name     // Suponiendo que `customer` tiene un campo Nombre
		userCedula = customer.Cedula

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type"})
		return
	}

	apiKeyResponse, err = apiKey.GenerateAPIKey(c)
	if err != nil {
		log.Printf("Error generating api key: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": apiKeyResponse.AccessToken,
		"userType":     userType,
		"userID":       userID,   // ID del usuario
		"userName":     userName, // Nombre del usuario
		"userNit":      userNit,
		"userCedula":   userCedula,
		"message":      "user logged successfully",
	})
}

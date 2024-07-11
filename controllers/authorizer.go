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

// Login
func Login(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error getting body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	auth, err := models.Parser[models.Auth](body)
	if err != nil {
		log.Printf("Error parssing body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	admin, err := storage.GetAdminByEmail(auth.Email)
	if err != nil {
		log.Printf("Error getting secret key: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	err = security.VerifyPassword(admin.Password, auth.Password)
	if err != nil {
		log.Printf("Error verifying password: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	apiKey, err := storage.NewAPIKey(admin, "", auth.Email)
	if err != nil {
		log.Printf("Error creating api key: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	apiKeyResponse, err := apiKey.GenerateAPIKey(c)
	if err != nil {
		log.Printf("Error generatting api key: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": apiKeyResponse.AccessToken,
		"message":      "user logged successfully",
	})
}

func CheckAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	_, err := storage.VerifyAPIKey(c, authHeader)
	if err != nil && storage.ErrNoHMACSigningMethod.Error() == err.Error() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unexpected api algorithm received"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User unauthorized"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

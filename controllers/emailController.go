package controllers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func SendEmail(c *gin.Context) {
	// Parsear los datos del formulario
	to := c.PostForm("to")
	subject := c.PostForm("subject")
	body := c.PostForm("body")

	// Validar los campos requeridos
	if to == "" || subject == "" || body == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan campos requeridos."})
		return
	}

	// Obtener los archivos PDF
	pdfFile, err := c.FormFile("pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo obtener el archivo PDF principal."})
		return
	}

	pdf2File, err := c.FormFile("pdf2")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo obtener el segundo archivo PDF."})
		return
	}

	// Crear el mensaje de correo
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Función para adjuntar un archivo
	attachFile := func(file *multipart.FileHeader) error {
		// Abrir el archivo
		openedFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("error al abrir el archivo %s: %v", file.Filename, err)
		}
		defer openedFile.Close()

		// Leer el contenido del archivo
		buffer := make([]byte, file.Size)
		_, err = openedFile.Read(buffer)
		if err != nil {
			return fmt.Errorf("error al leer el archivo %s: %v", file.Filename, err)
		}

		// Adjuntar el archivo al correo
		m.Attach(file.Filename, gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(buffer)
			return err
		}))
		return nil
	}

	// Adjuntar ambos archivos
	if err := attachFile(pdfFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := attachFile(pdf2File); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Configurar el dialer SMTP
	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		getEnvAsInt("SMTP_PORT"),
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
	)

	// Enviar el correo
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error al enviar el correo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar el correo."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Correo enviado exitosamente con ambos PDFs adjuntos."})
}

// Función auxiliar para convertir variables de entorno a int
func getEnvAsInt(name string) int {
	value := os.Getenv(name)
	var intValue int
	fmt.Sscanf(value, "%d", &intValue)
	return intValue
}

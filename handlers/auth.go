package handlers

import (
	"go-test/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Rutas Auth
func (h *Handler) LoginHandler(c *gin.Context) {

	// struct only use for this handler
	// credentials
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Check if the user put information correct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	// Chek if teh credentials exist and mount the bills into the user
	var user User
	err := h.DB.Where("email = ?", loginData.Email).Preload("Bills").First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// User not found
			c.JSON(404, gin.H{"error": "Credentials not valid"})
			return
		}
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// comparamos el hash con el password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credentials not valid"})
		return
	}

	// Creamos la llave secreta
	privateKey, err := utils.GetJWTPrivateKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating the privateKey"})
		return
	}

	// Verificar que privateKey no sea nulo
	if privateKey == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error null privateKey"})
		return
	}

	// Leemos el valor de el archivo config
	env_expiration := viper.GetInt("TOKEN_EXPIRATION")

	// Verificar si las claves esperadas están configuradas
	if !viper.IsSet("TOKEN_EXPIRATION") {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error viper not working"})
		return
	}
	// Generar el token JWT
	tokenString, err := utils.GenerateJWTToken(privateKey, loginData.Email, time.Duration(env_expiration)) // 3600 segundos = 1 hora
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating the token"})
		return
	}

	// Devolvemos el token y el ID del user
	c.JSON(200, gin.H{
		"id":    &user.ID,
		"token": tokenString,
	})
}

func (h *Handler) RegisterHandler(c *gin.Context) {
	h.NewUserHandler(c)
}

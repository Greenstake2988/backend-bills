package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credentials not valid"})
	}

	// Devolvemos el token
	c.JSON(200, &user)
}

func (h *Handler) RegisterHandler(c *gin.Context) {
	h.NewUserHandler(c)
}

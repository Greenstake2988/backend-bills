package handlers

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

// Rutas Login
func (app *App) LoginUser(c *gin.Context) {

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
	err := app.DB.Where("email = ? AND password = ?", loginData.Email, loginData.Password).Preload("Bills").First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// User not found
			c.JSON(404, gin.H{"error": "Credentials not valid"})
			return
		}
		c.JSON(500, gin.H{"error": "Internal server error"})
	}

	// Devolvemos el token
	c.JSON(200, &user)
}

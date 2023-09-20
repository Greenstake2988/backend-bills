package handlers

import (
	"backend-bills/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) Verification(c *gin.Context) {

	var verificationData struct {
		Email            string `json:"email"             binding:"required"`
		VerificationCode string `json:"verification_code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&verificationData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	var user models.User
	err := h.DB.Where("email = ?", verificationData.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// User not found
			c.JSON(404, gin.H{"error": "Credentials not valid"})
			return
		}
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if verificationData.VerificationCode != user.VerificationCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Code verification wrong.",
		})
		return
	}

	user.VerificationStatusEmail = true
	// Guardamos los cambios
	if err := h.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "User verification completed.",
	})

}

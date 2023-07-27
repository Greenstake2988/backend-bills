package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"go-test/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Rutas Login
func (app *App) LoginUser(c *gin.Context) {

	// Obtener la clave secreta desechamos el error
	jwtSecret, _ := utils.GetJWTPrivateKey()

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

	// Generamos el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	// Firmamos el token con la clave secreta
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating the token"})
		return
	}

	// Devolvemos el token
	c.JSON(200, gin.H{"token": tokenString})
}

// Midleware autenticador
func (app *App) AuthMiddleware(c *gin.Context) {
	// Obtener la clave secreta desechamos el error
	jwtSecret, _ := utils.GetJWTPrivateKey()

	// Obtener el token del header
	tokenString := c.GetHeader("Authorization")

	// Verificar que el token no este en blanco
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		c.Abort()
		return
	}

	// Validar el token y obtener datos de el usuario
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		c.Abort()
		return
	}
	// Agregar los datos del usuario  al contexto para que esten disponibles
	claims, _ := token.Claims.(jwt.MapClaims)
	c.Set("userEmail", claims["email"])

	// Continuar con el siguiente handler
	c.Next()
}

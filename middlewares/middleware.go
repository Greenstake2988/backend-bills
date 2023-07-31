package middlewares

import (
	"fmt"
	"backend-bills/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Middleware de autenticación JWT
func AuthMiddleware(c *gin.Context) {
	// Obtener el token del encabezado de autorización
	authHeader := c.GetHeader("Authorization")

	// Extraer el token del encabezado (el valor del token vendrá en formato "Bearer <token>")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	// Recuperamos la llave
	privateKey, err := utils.GetJWTPrivateKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Key"})
	}

	// Verificar si se proporcionó un token
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
		c.Abort()
		return
	}

	// Validar el token y obtener los datos del usuario
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("método de firma inválido: %v", token.Header["alg"])
		}
		return &privateKey.PublicKey, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Error al desencriptar y validar el token: %v", err)})
		c.Abort()
		return
	}

	// Agregar los datos del usuario al contexto para que estén disponibles en otras rutas
	claims, _ := token.Claims.(jwt.MapClaims)
	c.Set("userEmail", claims["email"])

	// Continuar con el siguiente handler
	c.Next()
}

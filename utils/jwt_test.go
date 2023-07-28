package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)


func TestGetJWTPrivateKey(t *testing.T) {

	privateKey, err := GetJWTPrivateKey()
	if err != nil {
		t.Fatalf("Error al obtener la clave privada: %v", err)
	}

	// Verificar que privateKey no sea nulo
	if privateKey == nil {
		t.Fatal("La clave privada no debería ser nula")
	}

	// Generar el token JWT
	userEmail := "ejemplo@gmail.com"
	// Definir el tiempo de expiración del token (por ejemplo, 1 hora)
	env_expiration := viper.GetInt("TOKEN_EXPIRATION")
	// Generar el token JWT
	tokenString, err := GenerateJWTToken(privateKey, userEmail, time.Duration(env_expiration)) 
	if err != nil {
		t.Fatalf("Error al generar el token: %v", err)
	}

	// Esperar un tiempo para que el token expire
	//time.Sleep(6 * time.Second)

	// Desencriptar y validar el token utilizando la clave pública
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("Método de firma inválido: %v", token.Header["alg"])
		}
		return &privateKey.PublicKey, nil
	})

	// Verificar que no haya ocurrido un error en el desencriptado y validación del token
	if err != nil || !token.Valid {
		t.Fatalf("Error al desencriptar y validar el token: %v", err)
	}

	// Verificar que los datos del token sean correctos
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		t.Fatalf("Error al obtener las claims del token")
	}

	// Verificar que el email en las claims coincida con el email utilizado para generar el token
	if email, found := claims["email"]; !found || email != userEmail {
		t.Fatalf("Email en las claims del token no coincide con el email utilizado para generar el token")
	}

}

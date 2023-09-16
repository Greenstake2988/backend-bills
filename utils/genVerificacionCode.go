package utils

import (
	"math/rand"
	"strconv"
)

// generateVerificationCode genera un código de verificación de 4 dígitos.
func GenerateVerificationCode() string {
	min := 1000
	max := 9999
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

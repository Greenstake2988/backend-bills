// utils.go

package utils

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// ValidatePassword valida si la contraseña cumple con los requisitos mínimos.
// La contraseña debe tener al menos 8 caracteres, al menos una letra mayúscula,
// al menos una letra minúscula y al menos un número.
func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("la contraseña debe tener al menos 8 caracteres")
	}

	// Usamos expresiones regulares para validar que haya al menos una letra mayúscula,
	// una letra minúscula y un número en la contraseña.
	containsUpperCase := regexp.MustCompile(`[A-Z]`).MatchString
	containsLowerCase := regexp.MustCompile(`[a-z]`).MatchString
	containsNumber := regexp.MustCompile(`[0-9]`).MatchString

	if !containsUpperCase(password) {
		return errors.New("la contraseña debe contener al menos una letra mayúscula")
	}

	if !containsLowerCase(password) {
		return errors.New("la contraseña debe contener al menos una letra minúscula")
	}

	if !containsNumber(password) {
		return errors.New("la contraseña debe contener al menos un número")
	}

	return nil
}

func HashAndValidatePassword(password string) (string, error) {

	// Validar la contraseña
	if err := validatePassword(password); err != nil {
		return "", err
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error al generar el hash de la contraseña: %v", err)
	}

	return string(hashedPassword), nil
}

// utils.go

package utils

import (
	"backend-bills/models"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	ErrorPassword8caracteres = 3
	ErrorPassword1Mayuscula  = 4
	ErrorPassword1Minuscula  = 5
	ErrorPassword1numero     = 6
	ErrorPasswordHash        = 7
)

// ValidatePassword valida si la contraseña cumple con los requisitos mínimos.
// La contraseña debe tener al menos 8 caracteres, al menos una letra mayúscula,
// al menos una letra minúscula y al menos un número.
func validatePassword(password string) []models.APIError {

	var ValidationErrors []models.APIError

	if len(password) < 8 {

		ValidationErrors = append(ValidationErrors, models.APIError{Code: ErrorPassword8caracteres, Message: "la contraseña debe tener al menos 8 caracteres"})
	}

	// Usamos expresiones regulares para validar que haya al menos una letra mayúscula,
	// una letra minúscula y un número en la contraseña.
	containsUpperCase := regexp.MustCompile(`[A-Z]`).MatchString
	containsLowerCase := regexp.MustCompile(`[a-z]`).MatchString
	containsNumber := regexp.MustCompile(`[0-9]`).MatchString

	if !containsUpperCase(password) {
		ValidationErrors = append(ValidationErrors, models.APIError{Code: ErrorPassword1Mayuscula, Message: "la contraseña debe contener al menos una letra mayúscula"})
	}

	if !containsLowerCase(password) {
		ValidationErrors = append(ValidationErrors, models.APIError{Code: ErrorPassword1Minuscula, Message: "la contraseña debe contener al menos una letra minúscula"})
	}

	if !containsNumber(password) {
		ValidationErrors = append(ValidationErrors, models.APIError{Code: ErrorPassword1numero, Message: "la contraseña debe contener al menos un número"})
	}

	return ValidationErrors
}

func HashAndValidatePassword(password string) (string, []models.APIError) {

	var ValidationErrors []models.APIError
	// Validar la contraseña
	if err := validatePassword(password); err != nil {
		return "", err
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ValidationErrors = append(ValidationErrors, models.APIError{Code: ErrorPasswordHash, Message: fmt.Sprintf("error al generar el hash de la contraseña: %v", err)})
		return "", ValidationErrors
	}

	return string(hashedPassword), nil
}

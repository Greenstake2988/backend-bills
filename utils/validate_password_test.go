package utils

import "testing"

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		wantErr  bool
	}{
		{password: "Abc12345", wantErr: false}, // Cumple con los requisitos
		{password: "abc12345", wantErr: true},  // No contiene mayúsculas
		{password: "ABC12345", wantErr: true},  // No contiene minúsculas
		{password: "Abcdefgh", wantErr: true},  // No contiene números
		{password: "12345678", wantErr: true},  // No contiene letras
		{password: "Abc12", wantErr: true},     // Menos de 8 caracteres
	}

	for _, test := range tests {
		err := ValidatePassword(test.password)
		if (err != nil) != test.wantErr {
			t.Errorf("Para el password '%s', esperado error: %v, obtenido error: %v", test.password, test.wantErr, err)
		}
	}
}

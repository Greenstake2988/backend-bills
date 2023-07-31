package utils

import (
	"time"
)

// CustomDate es una estructura personalizada que contiene un campo time.Time
type DateOnly struct {
	time.Time
}

// Deserializar el tipo ala fecha con el formato que buscamos
func (do *DateOnly) UnmarshalJSON(b []byte) error {
	// Aquí puedes manejar diferentes formatos de fecha según tus necesidades
	parsedDate, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	do.Time = parsedDate
	return nil
}

// MarshalJSON implementa la interfaz json.Marshaler para CustomDate
func (cd *DateOnly) MarshalJSON() ([]byte, error) {
	// Aquí puedes definir el formato en el que deseas que se muestre la fecha
	return []byte(`"` + cd.Time.Format("2006-01-02") + `"`), nil
}

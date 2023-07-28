package handlers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

// Conexion Base de Datos
func (h *Handler) ConnectDB() {

	var err error
	//Conexion sqlite3
	h.DB, err = gorm.Open(sqlite.Open("bills.sqlite"), &gorm.Config{})

	if err != nil {
		panic("error al conectar ala base de datos")
	}

	// AutoMigrate intenta crear la tabala si no existe
	err = h.DB.AutoMigrate(&User{}, &Bill{})

	// Habilitamos la funcion para poder usar claves foraneas
	h.DB.Exec("PRAGMA foreign_keys = ON;")

	if err != nil {
		panic("Error al crear la tabla en la base de datos")
	}
}

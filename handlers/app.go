package handlers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

// Conexion Base de Datos
func (app *App) ConnectDB() {

	var err error
	//Conexion sqlite3
	app.DB, err = gorm.Open(sqlite.Open("bills.sqlite"), &gorm.Config{})

	if err != nil {
		panic("error al conectar ala base de datos")
	}

	// AutoMigrate intenta crear la tabala si no existe
	err = app.DB.AutoMigrate(&User{}, &Bill{})

	// Habilitamos la funcion para poder usar claves foraneas
	app.DB.Exec("PRAGMA foreign_keys = ON;")

	if err != nil {
		panic("Error al crear la tabla en la base de datos")
	}
}

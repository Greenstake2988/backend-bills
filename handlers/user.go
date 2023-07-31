package handlers

import (
	"backend-bills/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string    `json:"email" binding:"required" gorm:"unique"`
	Password string    `json:"password" binding:"required"`
	Date     time.Time `json:"date" gorm:"type:date"`
	Bills    []Bill    `json:"bills" gorm:"constraint:OnDelete:CASCADE"`
}

// Rutas Users
func (h *Handler) GetUserHandler(c *gin.Context) {
	var user User

	userID := c.Param("id")

	// Check if the "id" parameter is empty
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// find the first value in the data base with userID
	if err := h.DB.Preload("Bills").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
func (h *Handler) UsersHandler(c *gin.Context) {
	// fecthar los datos de la base de datos
	var users []User
	// SELECT * FROM users;
	// SELECT * FROM bills WHERE user_id IN (1,2,3,4);
	h.DB.Preload("Bills").Find(&users)

	// Set the "Access-Control-Allow-Origin" header to allow all origins (*)
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(200, gin.H{
		"users": users,
	})
}
func (h *Handler) NewUserHandler(c *gin.Context) {
	var newUser User

	// Convierte el Json en el tipo de objeto que necesitamos
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Invalido"})
		return
	}

	// Si la fecha no fue proporcionada o esta en blano le pones la fecha de Ahora
	if newUser.Date.IsZero() {
		newUser.Date = time.Now()
	}

	// Verificar si el correo electrónico ya existe en la base de datos
	var existingUser User
	if err := h.DB.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		// Correo electrónico ya existente, devolver mensaje de error en formato JSON
		c.JSON(http.StatusBadRequest, gin.H{"error": "El correo ya existe elige otro."})
		return
	}

	// Validamos y hasehamos el password
	hashedValidatedPassword, err := utils.HashAndValidatePassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Pasamos el hased password al usuario nuevo
	newUser.Password = hashedValidatedPassword

	// Aqui guardaremos en la base de datos
	if err := h.DB.Create(&newUser).Error; err != nil {

		// Si no es un error de clave externa, devolver otro mensaje de error genérico
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "JSON Recibido", "data": newUser.Email})
}
func (h *Handler) DeleteUserHandler(c *gin.Context) {
	var user User

	userID := c.Param("id")

	// Check if the "id" parameter is empty
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// find the first value in the data base with userID
	if err := h.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// borramos user y bills asociados
	// unscoped para poder borrar definitivo
	if err := h.DB.Unscoped().Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User delete successfully"})
}
func (h *Handler) UpdateUser(c *gin.Context) {

	// credentials
	var updateData struct {
		Password string         `json:"password" `
		Date     utils.DateOnly `json:"date"`
	}

	//Recuperar el id de la ruta
	userID := c.Param("id")
	// Check if the "id" parameter is empty
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Invalido"})
		return
	}
	// Convierte el Json en el tipo de objeto que necesitamos
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear un usario en blanco
	var newUser User
	// ligar el usuario nuevo con los valores de la base de datos
	if err := h.DB.First(&newUser, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Si recibimos una nueva fecha la actualizamos
	if !updateData.Date.IsZero() {
		newUser.Date = updateData.Date.Time
	}

	// Si Recibimos el password lo intercambiamos
	if updateData.Password != "" {
		// Validamos y hasehamos el password
		hashedValidatedPassword, err := utils.HashAndValidatePassword(updateData.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// actualizamos al usuario
		newUser.Password = hashedValidatedPassword
	}

	// Guardamos los cambios
	if err := h.DB.Save(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado con exito"})
}

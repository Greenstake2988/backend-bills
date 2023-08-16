package handlers

import (
	"backend-bills/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
	Bills    []Bill `json:"bills" gorm:"constraint:OnDelete:CASCADE"`
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
	var errors []string

	// Convierte el Json en el tipo de objeto que necesitamos
	if err := c.ShouldBindJSON(&newUser); err != nil {
		// Checar si el error es de la validacion de campos de la libreria
		if verr, ok := err.(validator.ValidationErrors); ok {

			for _, e := range verr {
				errors = append(errors, "El "+e.Field()+" no es Valido")
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		errors = append(errors, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	// Verificar si el correo electrónico ya existe en la base de datos
	var existingUser User
	if err := h.DB.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		// Correo electrónico ya existente, devolver mensaje de error en formato JSON
		errors = append(errors, "El correo ya existe elige otro.")
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	// Validamos y hasehamos el password
	hashedValidatedPassword, err := utils.HashAndValidatePassword(newUser.Password)
	if err != nil {
		errors = append(errors, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errors})
		return
	}
	// Pasamos el hased password al usuario nuevo
	newUser.Password = hashedValidatedPassword

	// Aqui guardaremos en la base de datos
	if err := h.DB.Create(&newUser).Error; err != nil {

		// Si no es un error de clave externa, devolver otro mensaje de error genérico
		errors = append(errors, "Error al crear el usuario")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errors})

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

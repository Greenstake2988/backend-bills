package handlers

import (
	"go-test/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required" gorm:"unique"`
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

	// Convierte el Json en el tipo de objeto que necesitamos
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Invalido"})
		return
	}

	// Verificar si el correo electrónico ya existe en la base de datos
	var existingUser User
	if err := h.DB.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		// Correo electrónico ya existente, devolver mensaje de error en formato JSON
		c.JSON(http.StatusBadRequest, gin.H{"error": "El correo ya existe elige otro."})
		return
	}

	// Validar la contraseña antes de guardarla en la base de datos
	if err := utils.ValidatePassword(newUser.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash de la contraseña antes de guardarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear el password"})
		return
	}
	// Pasamos el hased password al usuario nuevo
	newUser.Password = string(hashedPassword)

	// Aqui guardaremos en la base de datos
	if err := h.DB.Create(&newUser).Error; err != nil {

		// Si no es un error de clave externa, devolver otro mensaje de error genérico
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "JSON Recibido", "data": newUser.Email})
}
func (handler *Handler) DeleteUserHandler(c *gin.Context) {
	var user User

	userID := c.Param("id")

	// Check if the "id" parameter is empty
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// find the first value in the data base with userID
	if err := handler.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// borramos user y bills asociados
	// unscoped para poder borrar definitivo
	if err := handler.DB.Unscoped().Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User delete successfully"})
}

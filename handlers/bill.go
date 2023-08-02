package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	UserID  uint      `json:"user_id" binding:"required"`
	Concept string    `json:"concept" binding:"required"`
	Price   float32   `json:"price" binding:"required"`
	Date    time.Time `json:"date" gorm:"type:date"`
}

// Rutas Bills
func (h *Handler) GetBillHandler(c *gin.Context) {
	var bill Bill

	billID := c.Param("id")

	// Check if the "id" parameter is empty
	if billID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalida"})
		return
	}

	// find the first value in the data base with billID
	if err := h.DB.First(&bill, billID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bill no encontrado"})
		return
	}

	c.JSON(http.StatusOK, bill)
}
func (h *Handler) BillsHandler(c *gin.Context) {
	// fecthar los datos de la base de datos
	var bills []Bill
	h.DB.Find(&bills)

	// Set the "Access-Control-Allow-Origin" header to allow all origins (*)
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(200, gin.H{
		"bills": bills,
	})
}
func (h *Handler) NewBillHandler(c *gin.Context) {
	var newBill Bill

	// Convierte el Json en el tipo de objeto que necesitamos
	if err := c.ShouldBindJSON(&newBill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Si la fecha no fue proporcionada o esta en blano le pones la fecha de Ahora
	if newBill.Date.IsZero() {
		newBill.Date = time.Now()
	}

	// Aqui guardaremos en la base de datos
	if err := h.DB.Create(&newBill).Error; err != nil {
		//Verificamos si cumple con la restricciond e llave foranea
		if isForeignKeyConstraintError(err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "el user_id debe ser de un usuario real"})
			return
		}
		// Si no es un error de clave externa, devolver otro mensaje de error genérico
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la Bills"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "JSON Recibido", "data": newBill.Concept})
}
func (h *Handler) DeleteBillHandler(c *gin.Context) {
	var bill Bill

	billID := c.Param("id")

	// Check if the "id" parameter is empty
	if billID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}

	// find the first value in the data base with billID
	if err := h.DB.First(&bill, billID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bill not found"})
		return
	}

	// delete the bill is found
	if err := h.DB.Delete(&bill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Bill"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bill delete successfully"})
}

func isForeignKeyConstraintError(err error) bool {
	sqliteErr, ok := err.(sqlite3.Error)
	if !ok {
		return false
	}

	return sqliteErr.Code == sqlite3.ErrConstraint
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	UserID  uint      `json:"user_id" binding:"required"`
	Concept string    `json:"concept" binding:"required"`
	Price   float32   `json:"price" binding:"required"`
	Date    time.Time `json:"date" gorm:"type:date"`
}

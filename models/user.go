package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email                   string `json:"email" binding:"required,email" gorm:"unique"`
	Password                string `json:"password" binding:"required"`
	VerificationCode        string `json:"verification_code"`
	VerificationStatusEmail bool   `json:"verification_status_email"`
	Bills                   []Bill `json:"bills" gorm:"constraint:OnDelete:CASCADE"`
}

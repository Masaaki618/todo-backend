package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null;uniqueIndex:idx_user_email"`
	Password string `json:"password" gorm:"not null"`
}

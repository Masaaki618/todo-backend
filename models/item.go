package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Title string `json:"title" gorm:"not null"`
	Done  bool   `json:"done" gorm:"not null:default:false"`
}

package main

import (
	"todo-backend/infra"
	"todo-backend/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic("failed to migrate database:")
	}
}

package main

import (
	"log"
	"os"
	"todo-backend/controllers"
	"todo-backend/models"
	"todo-backend/repositories"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{Id: 1, Title: "筋トレ", Done: false},
		{Id: 2, Title: "サッカー", Done: true},
		{Id: 3, Title: "スノボ", Done: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	group := r.Group("api/v1/items")
	group.GET("", itemController.FindAll)
	group.GET("/:id", itemController.FindByID)
	group.POST("", itemController.Create)
	group.PUT("/:id", itemController.Update)
	group.DELETE("/:id", itemController.Delete)

	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = "0.0.0.0:8080"
	}

	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

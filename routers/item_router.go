package routers

import (
	"todo-backend/controllers"
	"todo-backend/repositories"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupItemRoutes(r *gin.Engine, db *gorm.DB) {
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	group := r.Group("api/v1/items")
	group.GET("", itemController.FindAll)
	group.GET("/:id", itemController.FindByID)
	group.POST("", itemController.Create)
	group.PUT("/:id", itemController.Update)
	group.DELETE("/:id", itemController.Delete)
}

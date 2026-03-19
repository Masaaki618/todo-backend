package routers

import (
	"todo-backend/controllers"
	"todo-backend/infra"
	"todo-backend/repositories"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	//items := []models.Item{
	//	{Title: "筋トレ", Done: false},
	//	{Title: "サッカー", Done: true},
	//	{Title: "スノボ", Done: false},
	//}

	//itemRepository := repositories.NewItemMemoryRepository(items)
	db := infra.SetupDB()
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	group := r.Group("api/v1/items")
	group.GET("", itemController.FindAll)
	group.GET("/:id", itemController.FindByID)
	group.POST("", itemController.Create)
	group.PUT("/:id", itemController.Update)
	group.DELETE("/:id", itemController.Delete)

	return r
}

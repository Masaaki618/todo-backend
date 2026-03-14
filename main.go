package main

import (
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
	r.Run("localhost:8080") // デフォルトで0.0.0.0:8080で待機します
	// loggerとrecoveryミドルウェア付きGinルーター作成

	// 簡単なGETエンドポイント定義

	// ポート8080でサーバー起動（デフォルト）
	// 0.0.0.0:8080（Windowsではlocalhost:8080）で待機
}

package controllers

import (
	"net/http"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
)

type IItemController interface {
	FindAll(c *gin.Context)
}

type ItemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemController {
	return &ItemController{service: service}
}

func (i *ItemController) FindAll(ctx *gin.Context) {
	items, err := i.service.FindALL()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

package controllers

import (
	"net/http"
	"strconv"
	"todo-backend/dto"
	"todo-backend/models"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
)

type IItemController interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	Create(c *gin.Context)
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

func (i *ItemController) FindByID(ctx *gin.Context) {
	var item *models.Item
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid id"})
		return
	}

	item, err = i.service.FindByID(uint(id))
	if err != nil {
		if err.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

func (i *ItemController) Create(ctx *gin.Context) {
	var input dto.CreateItemInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem, err := i.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newItem})
}

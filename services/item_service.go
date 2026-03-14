package services

import (
	"todo-backend/models"
	"todo-backend/repositories"
)

type IItemService interface {
	FindALL() (*[]models.Item, error)
	FindByID(id uint) (*models.Item, error)
}

type ItemService struct {
	service repositories.IItemRepository
}

func NewItemService(service repositories.IItemRepository) IItemService {
	return &ItemService{service: service}
}

func (s *ItemService) FindALL() (*[]models.Item, error) {
	return s.service.FindAll()
}

func (s *ItemService) FindByID(id uint) (*models.Item, error) {
	return s.service.FindByID(id)
}

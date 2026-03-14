package services

import (
	"todo-backend/dto"
	"todo-backend/models"
	"todo-backend/repositories"
)

type IItemService interface {
	FindALL() (*[]models.Item, error)
	FindByID(id uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
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

func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	return s.service.Create(models.Item{
		Title: createItemInput.Title,
	})
}

func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	item, err := s.service.FindByID(itemId)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Title != nil {
		item.Title = *updateItemInput.Title
	}

	if updateItemInput.Done != nil {
		item.Done = *updateItemInput.Done
	}

	return s.service.Update(*item)
}

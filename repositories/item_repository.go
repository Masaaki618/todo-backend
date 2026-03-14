package repositories

import (
	"errors"
	"todo-backend/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindByID(id uint) (*models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindByID(id uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.Id == id {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}

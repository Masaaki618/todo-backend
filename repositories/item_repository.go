package repositories

import (
	"errors"
	"todo-backend/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindByID(id uint) (*models.Item, error)
	Create(item models.Item) (*models.Item, error)
	Update(item models.Item) (*models.Item, error)
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

func (r *ItemMemoryRepository) Create(item models.Item) (*models.Item, error) {
	item.Id = uint(len(r.items) + 1)
	r.items = append(r.items, item)
	return &item, nil
}

func (r *ItemMemoryRepository) Update(upDateItem models.Item) (*models.Item, error) {
	for i, item := range r.items {
		if item.Id == upDateItem.Id {
			r.items[i] = upDateItem
			return &upDateItem, nil
		}
	}
	return nil, errors.New("Item not found")
}

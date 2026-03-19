package repositories

import (
	"errors"
	"todo-backend/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindByID(id uint) (*models.Item, error)
	Create(item models.Item) (*models.Item, error)
	Update(item models.Item) (*models.Item, error)
	Delete(id uint) error
}

type ItemMemoryRepository struct {
	items []models.Item
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindByID(id uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (r *ItemMemoryRepository) Create(item models.Item) (*models.Item, error) {
	item.ID = uint(len(r.items) + 1)
	r.items = append(r.items, item)
	return &item, nil
}

func (r *ItemMemoryRepository) Update(upDateItem models.Item) (*models.Item, error) {
	for i, item := range r.items {
		if item.ID == upDateItem.ID {
			r.items[i] = upDateItem
			return &upDateItem, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (r *ItemMemoryRepository) Delete(id uint) error {
	for i, item := range r.items {
		if item.ID == id {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("Item not found")
}

func (i ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := i.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

func (i ItemRepository) FindByID(id uint) (*models.Item, error) {
	var item models.Item
	result := i.db.First(&item, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("Item not found")
		}

		return nil, result.Error
	}
	return &item, nil
}

func (i ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := i.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newItem, nil
}

func (i ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	result := i.db.Save(&updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}

func (i ItemRepository) Delete(id uint) error {
	deleteItem, err := i.FindByID(id)
	if err != nil {
		return err
	}
	result := i.db.Delete(&deleteItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

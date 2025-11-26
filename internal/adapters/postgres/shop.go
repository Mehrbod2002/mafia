package postgres

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"

	"gorm.io/gorm"
)

type shopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ports.ShopRepository {
	return &shopRepository{db: db}
}

func (r *shopRepository) List() ([]domain.ShopItem, error) {
	var items []domain.ShopItem
	err := r.db.Find(&items).Error
	return items, err
}

func (r *shopRepository) FindByID(id uint) (*domain.ShopItem, error) {
	var item domain.ShopItem
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *shopRepository) Create(item *domain.ShopItem) error {
	return r.db.Create(item).Error
}

func (r *shopRepository) Update(item *domain.ShopItem) error {
	return r.db.Save(item).Error
}

func (r *shopRepository) Delete(id uint) error {
	return r.db.Delete(&domain.ShopItem{}, id).Error
}

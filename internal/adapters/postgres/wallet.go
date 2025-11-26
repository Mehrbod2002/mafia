package postgres

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"

	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) ports.WalletRepository {
	return &walletRepository{db}
}

func (r *walletRepository) Create(w *domain.Wallet) error {
	return r.db.Create(w).Error
}

func (r *walletRepository) FindByUserID(id uint) (*domain.Wallet, error) {
	var w domain.Wallet
	err := r.db.Where("user_id = ?", id).First(&w).Error
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *walletRepository) Update(w *domain.Wallet) error {
	return r.db.Save(w).Error
}

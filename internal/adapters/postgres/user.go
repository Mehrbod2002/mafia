package postgres

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(u *domain.User) error {
	return r.db.Create(u).Error
}

func (r *userRepository) FindByPhone(phone string) (*domain.User, error) {
	var u domain.User
	err := r.db.Where("phone = ?", phone).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Update(u *domain.User) error {
	return r.db.Save(u).Error
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var u domain.User
	err := r.db.Preload("Profile").Preload("Wallet").First(&u, id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) UpdateOTP(id uint, otp string, expires time.Time) error {
	return r.db.Model(&domain.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"otp":         otp,
		"otp_expires": expires,
	}).Error
}

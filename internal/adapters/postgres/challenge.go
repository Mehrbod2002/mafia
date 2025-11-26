package postgres

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"

	"gorm.io/gorm"
)

type challengeRepository struct {
	db *gorm.DB
}

func NewChallengeRepository(db *gorm.DB) ports.ChallengeRepository {
	return &challengeRepository{db}
}

func (r *challengeRepository) Create(c *domain.Challenge) error {
	return r.db.Create(c).Error
}

func (r *challengeRepository) FindByID(id uint) (*domain.Challenge, error) {
	var c domain.Challenge
	err := r.db.First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *challengeRepository) List() ([]domain.Challenge, error) {
	var challenges []domain.Challenge
	err := r.db.Find(&challenges).Error
	return challenges, err
}

package postgres

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"gorm.io/gorm"
)

type ${repo}Repository struct{ db *gorm.DB }

func New${repo}Repository(db *gorm.DB) ports.${repo}Repository {
	return &${repo}Repository{db}
}

func (r *${repo}Repository) Create(m *domain.$repo) error { return r.db.Create(m).Error }
func (r *${repo}Repository) FindByID(id uint) (*domain.$repo, error) {
	var m domain.$repo
	err := r.db.First(&m, id).Error
	return &m, err
}

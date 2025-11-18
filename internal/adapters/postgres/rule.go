package postgres

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"

	"gorm.io/gorm"
)

type ruleRepository struct {
	db *gorm.DB
}

func NewRuleRepository(db *gorm.DB) ports.RuleRepository {
	return &ruleRepository{db: db}
}

func (r *ruleRepository) Create(rule *domain.GameRule) error {
	return r.db.Create(rule).Error
}

func (r *ruleRepository) List() ([]domain.GameRule, error) {
	var rules []domain.GameRule
	err := r.db.Find(&rules).Error
	return rules, err
}

func (r *ruleRepository) Update(rule *domain.GameRule) error {
	return r.db.Save(rule).Error
}

func (r *ruleRepository) Delete(id uint) error {
	return r.db.Delete(&domain.GameRule{}, id).Error
}

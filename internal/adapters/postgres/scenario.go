package postgres

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"

	"gorm.io/gorm"
)

type scenarioRepository struct {
	db *gorm.DB
}

func NewScenarioRepository(db *gorm.DB) ports.ScenarioRepository {
	return &scenarioRepository{db: db}
}

func (r *scenarioRepository) Create(s *domain.Scenario) error {
	return r.db.Create(s).Error
}

func (r *scenarioRepository) List() ([]domain.Scenario, error) {
	var scenarios []domain.Scenario
	err := r.db.Find(&scenarios).Error
	return scenarios, err
}

func (r *scenarioRepository) Update(s *domain.Scenario) error {
	return r.db.Save(s).Error
}

func (r *scenarioRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Scenario{}, id).Error
}

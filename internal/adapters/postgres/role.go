package postgres

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) ports.RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) Create(role *domain.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) FindByID(id uint) (*domain.Role, error) {
	var role domain.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) List() ([]domain.Role, error) {
	var roles []domain.Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *roleRepository) Update(role *domain.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Role{}, id).Error
}

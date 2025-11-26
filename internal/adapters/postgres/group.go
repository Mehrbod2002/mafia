package postgres

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"

	"gorm.io/gorm"
)

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) ports.GroupRepository {
	return &groupRepository{db}
}

func (r *groupRepository) Create(g *domain.Group) error {
	return r.db.Create(g).Error
}

func (r *groupRepository) FindByID(id uint) (*domain.Group, error) {
	var g domain.Group
	err := r.db.Preload("Members").First(&g, id).Error
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *groupRepository) Update(g *domain.Group) error {
	return r.db.Save(g).Error
}

func (r *groupRepository) AddMember(groupID, userID uint) error {
	return r.db.Exec("INSERT INTO group_members (group_id, user_id) VALUES (?, ?)", groupID, userID).Error
}

func (r *groupRepository) RemoveMember(groupID, userID uint) error {
	return r.db.Exec("DELETE FROM group_members WHERE group_id = ? AND user_id = ?", groupID, userID).Error
}

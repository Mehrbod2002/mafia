package services

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type groupService struct {
	groupRepo ports.GroupRepository
	userRepo  ports.UserRepository
}

func NewGroupService(groupRepo ports.GroupRepository, userRepo ports.UserRepository) ports.GroupService {
	return &groupService{groupRepo, userRepo}
}

func (s *groupService) CreateGroup(ownerID uint, name string) (*domain.Group, error) {
	group := &domain.Group{Name: name, OwnerID: ownerID}
	err := s.groupRepo.Create(group)
	return group, err
}

func (s *groupService) Invite(groupID, userID uint) error {
	return s.groupRepo.AddMember(groupID, userID)
}

func (s *groupService) Kick(groupID, userID uint) error {
	return s.groupRepo.RemoveMember(groupID, userID)
}

func (s *groupService) GetStats(groupID uint) (map[string]interface{}, error) {
	group, err := s.groupRepo.FindByID(groupID)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"members": len(group.Members),
		"points":  group.Points,
	}, nil
}

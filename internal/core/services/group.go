package services

import (
	"context"
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type groupService struct {
	groupRepo ports.GroupRepository
	userRepo  ports.UserRepository
	events    ports.EventBus
}

func NewGroupService(groupRepo ports.GroupRepository, userRepo ports.UserRepository, events ports.EventBus) ports.GroupService {
	return &groupService{groupRepo: groupRepo, userRepo: userRepo, events: events}
}

func (s *groupService) CreateGroup(ownerID uint, name string) (*domain.Group, error) {
	group := &domain.Group{Name: name, OwnerID: ownerID}
	if err := s.groupRepo.Create(group); err != nil {
		return nil, err
	}
	if s.events != nil {
		s.events.Publish(context.Background(), "group.created", group)
	}
	return group, nil
}

func (s *groupService) Invite(groupID, userID uint) error {
	if err := s.groupRepo.AddMember(groupID, userID); err != nil {
		return err
	}
	if s.events != nil {
		s.events.Publish(context.Background(), "group.member_added", map[string]uint{"group_id": groupID, "user_id": userID})
	}
	return nil
}

func (s *groupService) Kick(groupID, userID uint) error {
	if err := s.groupRepo.RemoveMember(groupID, userID); err != nil {
		return err
	}
	if s.events != nil {
		s.events.Publish(context.Background(), "group.member_removed", map[string]uint{"group_id": groupID, "user_id": userID})
	}
	return nil
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

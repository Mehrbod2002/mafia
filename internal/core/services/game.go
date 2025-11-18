package services

import (
	"fmt"
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"math/rand"
)

type gameService struct {
	roomRepo ports.RoomRepository
	roleRepo ports.RoleRepository
	userRepo ports.UserRepository
}

func NewGameService(roomRepo ports.RoomRepository, roleRepo ports.RoleRepository, userRepo ports.UserRepository) ports.GameService {
	return &gameService{roomRepo, roleRepo, userRepo}
}

func (s *gameService) CreateRoom(hostID uint, roomType string) (*domain.GameRoom, error) {
	room := &domain.GameRoom{
		HostID: hostID,
		Type:   roomType,
		Code:   randString(6),
	}
	err := s.roomRepo.Create(room)
	return room, err
}

func (s *gameService) ListRooms() ([]domain.GameRoom, error) {
	return s.roomRepo.ListWaiting()
}

func (s *gameService) JoinRoom(roomID, userID uint) error {
	room, err := s.roomRepo.FindByID(roomID)
	if err != nil || len(room.Players) >= 20 {
		return fmt.Errorf("cannot join")
	}
	return s.roomRepo.AddPlayer(roomID, userID)
}

func (s *gameService) LeaveRoom(roomID, userID uint) error {
	return s.roomRepo.RemovePlayer(roomID, userID)
}

func (s *gameService) StartGame(roomID uint) error {
	room, err := s.roomRepo.FindByID(roomID)
	if err != nil || len(room.Players) < 6 {
		return fmt.Errorf("not enough players")
	}
	assignRoles(room, s.roleRepo)
	room.Status = "playing"
	room.Phase = "night"
	return s.roomRepo.Update(room)
}

func assignRoles(room *domain.GameRoom, roleRepo ports.RoleRepository) {
	// TODO: implement role assignment logic
}

func (s *gameService) Vote(roomID, userID, targetID uint) error {
	return nil
}

func (s *gameService) UseAbility(roomID, userID uint, ability string, targetID uint) error {
	return nil
}

func randString(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

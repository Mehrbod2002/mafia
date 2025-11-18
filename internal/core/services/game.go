package services

import (
	"context"
	"encoding/json"
	"fmt"
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"math/rand"
	"time"
)

type gameService struct {
	roomRepo ports.RoomRepository
	roleRepo ports.RoleRepository
	userRepo ports.UserRepository
	events   ports.EventBus
}

func NewGameService(roomRepo ports.RoomRepository, roleRepo ports.RoleRepository, userRepo ports.UserRepository, events ports.EventBus) ports.GameService {
	return &gameService{roomRepo: roomRepo, roleRepo: roleRepo, userRepo: userRepo, events: events}
}

func (s *gameService) CreateRoom(hostID uint, roomType string) (*domain.GameRoom, error) {
	room := &domain.GameRoom{
		HostID: hostID,
		Type:   roomType,
		Code:   randString(6),
	}
	if err := s.roomRepo.Create(room); err != nil {
		return nil, err
	}
	if s.events != nil {
		s.events.Publish(context.Background(), "game.room_created", room)
	}
	return room, nil
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
	assignments, err := assignRoles(room, s.roleRepo)
	if err != nil {
		return err
	}
	room.Status = "playing"
	room.Phase = "night"
	room.DayCount = 1
	room.Results = assignments
	if err := s.roomRepo.Update(room); err != nil {
		return err
	}
	if s.events != nil {
		s.events.Publish(context.Background(), "game.started", room)
	}
	return nil
}

func assignRoles(room *domain.GameRoom, roleRepo ports.RoleRepository) (string, error) {
	roles, err := roleRepo.List()
	if err != nil {
		return "", err
	}
	var pool []string
	for _, r := range roles {
		count := r.MaxCount
		if count == 0 {
			count = 1
		}
		for i := 0; i < count; i++ {
			pool = append(pool, r.Name)
		}
	}
	if len(pool) < len(room.Players) {
		for len(pool) < len(room.Players) {
			pool = append(pool, "Villager")
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pool), func(i, j int) { pool[i], pool[j] = pool[j], pool[i] })

	assignments := make(map[uint]string)
	for idx, p := range room.Players {
		assignments[p.ID] = pool[idx%len(pool)]
	}
	data, _ := json.Marshal(assignments)
	return string(data), nil
}

func (s *gameService) Vote(roomID, userID, targetID uint) error {
	room, err := s.roomRepo.FindByID(roomID)
	if err != nil {
		return err
	}
	votes := map[string]uint{"voter": userID, "target": targetID}
	entry, _ := json.Marshal(votes)
	room.Results += string(entry) + "\n"
	return s.roomRepo.Update(room)
}

func (s *gameService) UseAbility(roomID, userID uint, ability string, targetID uint) error {
	room, err := s.roomRepo.FindByID(roomID)
	if err != nil {
		return err
	}
	abilityLog := map[string]interface{}{
		"user":    userID,
		"ability": ability,
		"target":  targetID,
	}
	entry, _ := json.Marshal(abilityLog)
	room.Results += string(entry) + "\n"
	return s.roomRepo.Update(room)
}

func (s *gameService) AdvancePhase(roomID uint) (*domain.GameRoom, error) {
	room, err := s.roomRepo.FindByID(roomID)
	if err != nil {
		return nil, err
	}
	if room.Phase == "night" {
		room.Phase = "day"
	} else {
		room.Phase = "night"
		room.DayCount++
	}
	if err := s.roomRepo.Update(room); err != nil {
		return nil, err
	}
	if s.events != nil {
		s.events.Publish(context.Background(), "game.phase_changed", room)
	}
	return room, nil
}

func randString(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

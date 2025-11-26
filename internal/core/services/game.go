package services

import (
	"context"
	"fmt"
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"math/rand"
	"time"
)

type gameService struct {
	roomRepo  ports.RoomRepository
	roleRepo  ports.RoleRepository
	userRepo  ports.UserRepository
	events    ports.EventBus
	abilities map[string]domain.AbilityOption
}

func NewGameService(roomRepo ports.RoomRepository, roleRepo ports.RoleRepository, userRepo ports.UserRepository, events ports.EventBus) ports.GameService {
	return &gameService{roomRepo: roomRepo, roleRepo: roleRepo, userRepo: userRepo, events: events, abilities: domain.AbilityIndex()}
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

	room.Status = "playing"
	room.Phase = "night"
	room.DayCount = 1

	state, err := s.assignRoles(room)
	if err != nil {
		return err
	}

	if err := s.saveGameState(room, state); err != nil {
		return err
	}
	if s.events != nil {
		s.events.Publish(context.Background(), "game.started", room)
	}
	return nil
}

func (s *gameService) assignRoles(room *domain.GameRoom) (*domain.GameState, error) {
	roles, err := s.roleRepo.List()
	if err != nil {
		return nil, err
	}

	roleIndex := make(map[string]domain.Role)
	var pool []string
	for _, r := range roles {
		roleIndex[r.Name] = r
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

	state := domain.NewGameState(room.Phase, room.DayCount)
	for idx, p := range room.Players {
		roleName := pool[idx%len(pool)]
		role := roleIndex[roleName]
		assignment := domain.PlayerAssignment{
			Role:          roleName,
			Team:          role.Team,
			Abilities:     append([]string{}, role.Abilities...),
			Alive:         true,
			UsedAbilities: map[string]domain.AbilityUsage{},
		}
		state.Assignments[p.ID] = assignment
	}
	return state, nil
}

func (s *gameService) Vote(roomID, userID, targetID uint) error {
	room, err := s.roomRepo.FindByID(roomID)
	if err != nil {
		return err
	}

	if room.Phase != "day" {
		return fmt.Errorf("votes are only allowed during the day phase")
	}

	state, err := s.loadGameState(room)
	if err != nil {
		return err
	}

	voter, ok := state.Assignments[userID]
	if !ok || !voter.Alive {
		return fmt.Errorf("voter is not active in this game")
	}

	if _, ok := state.Assignments[targetID]; !ok {
		return fmt.Errorf("invalid vote target")
	}

	vote := domain.VoteLog{
		Voter:     userID,
		Target:    targetID,
		Phase:     room.Phase,
		Day:       room.DayCount,
		Timestamp: time.Now(),
	}
	state.Votes = append(state.Votes, vote)

	return s.saveGameState(room, state)
}

func (s *gameService) UseAbility(roomID, userID uint, ability string, targetID uint) error {
	room, err := s.roomRepo.FindByID(roomID)
	if err != nil {
		return err
	}

	if room.Status != "playing" {
		return fmt.Errorf("game has not started")
	}

	state, err := s.loadGameState(room)
	if err != nil {
		return err
	}

	player, ok := state.Assignments[userID]
	if !ok || !player.Alive {
		return fmt.Errorf("player not active in this room")
	}

	definition, ok := s.abilities[ability]
	if !ok {
		return fmt.Errorf("unknown ability")
	}

	if len(player.Abilities) > 0 && !containsAbility(player.Abilities, ability) {
		return fmt.Errorf("ability not available to this role")
	}

	if definition.Phase != "both" && definition.Phase != room.Phase {
		return fmt.Errorf("ability can only be used during %s", definition.Phase)
	}

	if definition.Side != "" && definition.Side != "neutral" && player.Team != "" && player.Team != definition.Side {
		return fmt.Errorf("ability side does not match player team")
	}

	if targetID != 0 {
		if target, ok := state.Assignments[targetID]; !ok || !target.Alive {
			return fmt.Errorf("invalid target")
		}
	}

	if player.UsedAbilities == nil {
		player.UsedAbilities = map[string]domain.AbilityUsage{}
	}
	if usage, ok := player.UsedAbilities[ability]; ok && usage.Day == state.DayCount && usage.Phase == room.Phase {
		return fmt.Errorf("ability already used this %s", room.Phase)
	}

	player.UsedAbilities[ability] = domain.AbilityUsage{Day: state.DayCount, Phase: room.Phase}
	state.Assignments[userID] = player

	log := domain.AbilityAction{
		UserID:    userID,
		Ability:   ability,
		TargetID:  targetID,
		Phase:     room.Phase,
		Day:       room.DayCount,
		Timestamp: time.Now(),
	}
	state.Abilities = append(state.Abilities, log)

	return s.saveGameState(room, state)
}

func (s *gameService) AdvancePhase(roomID uint) (*domain.GameRoom, error) {
	room, err := s.roomRepo.FindByID(roomID)
	if err != nil {
		return nil, err
	}

	state, err := s.loadGameState(room)
	if err != nil {
		return nil, err
	}

	if room.Phase == "night" {
		room.Phase = "day"
	} else {
		room.Phase = "night"
		room.DayCount++
	}

	state.Phase = room.Phase
	state.DayCount = room.DayCount

	if err := s.saveGameState(room, state); err != nil {
		return nil, err
	}
	if s.events != nil {
		s.events.Publish(context.Background(), "game.phase_changed", room)
	}
	return room, nil
}

func (s *gameService) loadGameState(room *domain.GameRoom) (*domain.GameState, error) {
	state, err := domain.ParseGameState(room.Results)
	if err != nil {
		return nil, err
	}
	if state.Phase == "" {
		state.Phase = room.Phase
	}
	if state.DayCount == 0 {
		state.DayCount = room.DayCount
	}
	return state, nil
}

func (s *gameService) saveGameState(room *domain.GameRoom, state *domain.GameState) error {
	state.Phase = room.Phase
	state.DayCount = room.DayCount

	serialized, err := state.Serialize()
	if err != nil {
		return err
	}
	room.Results = serialized
	return s.roomRepo.Update(room)
}

func randString(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func containsAbility(list []string, ability string) bool {
	for _, a := range list {
		if a == ability {
			return true
		}
	}
	return false
}

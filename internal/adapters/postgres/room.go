package postgres

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) ports.RoomRepository {
	return &roomRepository{db}
}

func (r *roomRepository) Create(room *domain.GameRoom) error {
	return r.db.Create(room).Error
}

func (r *roomRepository) FindByID(id uint) (*domain.GameRoom, error) {
	var room domain.GameRoom
	err := r.db.Preload("Players").First(&room, id).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) ListWaiting() ([]domain.GameRoom, error) {
	var rooms []domain.GameRoom
	err := r.db.Where("status = ?", "waiting").Find(&rooms).Error
	return rooms, err
}

func (r *roomRepository) Update(room *domain.GameRoom) error {
	return r.db.Save(room).Error
}

func (r *roomRepository) AddPlayer(roomID, userID uint) error {
	return r.db.Exec("INSERT INTO room_players (game_room_id, user_id) VALUES (?, ?)", roomID, userID).Error
}

func (r *roomRepository) RemovePlayer(roomID, userID uint) error {
	return r.db.Exec("DELETE FROM room_players WHERE game_room_id = ? AND user_id = ?", roomID, userID).Error
}

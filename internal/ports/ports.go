package ports

import (
	"mafia/internal/core/domain"
	"time"
)

type UserRepository interface {
	Create(*domain.User) error
	FindByPhone(string) (*domain.User, error)
	Update(*domain.User) error
	FindByID(uint) (*domain.User, error)
	UpdateOTP(id uint, otp string, expires time.Time) error
}

type WalletRepository interface {
	Create(*domain.Wallet) error
	FindByUserID(uint) (*domain.Wallet, error)
	Update(*domain.Wallet) error
}

type ChallengeRepository interface {
	Create(*domain.Challenge) error
	FindByID(uint) (*domain.Challenge, error)
	List() ([]domain.Challenge, error)
}

type GroupRepository interface {
	Create(*domain.Group) error
	FindByID(uint) (*domain.Group, error)
	Update(*domain.Group) error
	AddMember(groupID, userID uint) error
	RemoveMember(groupID, userID uint) error
}

type RoomRepository interface {
	Create(*domain.GameRoom) error
	FindByID(uint) (*domain.GameRoom, error)
	ListWaiting() ([]domain.GameRoom, error)
	Update(*domain.GameRoom) error
	AddPlayer(roomID, userID uint) error
	RemovePlayer(roomID, userID uint) error
}

type RoleRepository interface {
	Create(*domain.Role) error
	FindByID(uint) (*domain.Role, error)
	List() ([]domain.Role, error)
	Update(*domain.Role) error
	Delete(id uint) error
}

type Repositories struct {
	User      UserRepository
	Wallet    WalletRepository
	Challenge ChallengeRepository
	Group     GroupRepository
	Room      RoomRepository
	Role      RoleRepository
}

type UserService interface {
	Register(phone string) error
	VerifyOTP(domain.VerifyOTPRequest) (string, uint, error)
	Login(phone string) error
	ValidateToken(token string) (uint, error)
	IsAdmin(id uint) (bool, error)
	GetProfile(id uint) (*domain.Profile, error)
	UpdateProfile(id uint, req domain.UpdateProfileRequest) (*domain.Profile, error)
	BanUser(id uint) error
	SuspendUser(id uint) error
}

type WalletService interface {
	GetWallet(userID uint) (*domain.Wallet, error)
	InitiatePurchase(userID uint, planID string) (string, error)
}

type ChallengeService interface {
	List() ([]domain.Challenge, error)
	Complete(challengeID, userID uint) error
}

type GroupService interface {
	CreateGroup(ownerID uint, name string) (*domain.Group, error)
	Invite(groupID, userID uint) error
	Kick(groupID, userID uint) error
	GetStats(groupID uint) (map[string]interface{}, error)
}

type GameService interface {
	CreateRoom(hostID uint, roomType string) (*domain.GameRoom, error)
	ListRooms() ([]domain.GameRoom, error)
	JoinRoom(roomID, userID uint) error
	LeaveRoom(roomID, userID uint) error
	StartGame(roomID uint) error
	Vote(roomID, userID, targetID uint) error
	UseAbility(roomID, userID uint, ability string, targetID uint) error
}

type Services struct {
	User      UserService
	Wallet    WalletService
	Challenge ChallengeService
	Group     GroupService
	Game      GameService
}

// SFU represents the WebRTC bridge used by websocket handlers.
type SFU interface{}

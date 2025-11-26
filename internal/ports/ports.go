package ports

import (
	"context"
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

type ShopRepository interface {
	List() ([]domain.ShopItem, error)
	FindByID(uint) (*domain.ShopItem, error)
	Create(*domain.ShopItem) error
	Update(*domain.ShopItem) error
	Delete(id uint) error
}

type RuleRepository interface {
	Create(*domain.GameRule) error
	List() ([]domain.GameRule, error)
	Update(*domain.GameRule) error
	Delete(id uint) error
}

type ScenarioRepository interface {
	Create(*domain.Scenario) error
	List() ([]domain.Scenario, error)
	Update(*domain.Scenario) error
	Delete(id uint) error
}

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration)
	Get(ctx context.Context, key string) (interface{}, bool)
	Delete(ctx context.Context, key string)
}

type Queue interface {
	Enqueue(func()) error
	Close()
}

type EventBus interface {
	Publish(ctx context.Context, topic string, payload interface{})
	Subscribe(topic string, handler func(ctx context.Context, payload interface{}))
}

type NotificationSender interface {
	Send(userID uint, channel, message string) error
}

type PaymentProvider interface {
	CreatePaymentURL(userID uint, planID string) (string, error)
}

type Infrastructure struct {
	Cache         Cache
	Queue         Queue
	Events        EventBus
	Notifications NotificationSender
	Payments      PaymentProvider
}

type Repositories struct {
	User      UserRepository
	Wallet    WalletRepository
	Challenge ChallengeRepository
	Group     GroupRepository
	Room      RoomRepository
	Role      RoleRepository
	Shop      ShopRepository
	Rule      RuleRepository
	Scenario  ScenarioRepository
}

type UserService interface {
	Register(phone string) error
	VerifyOTP(domain.VerifyOTPRequest) (string, uint, error)
	Login(phone string) error
	ValidateToken(token string) (uint, error)
	IsAdmin(id uint) (bool, error)
	GetProfile(id uint) (*domain.Profile, error)
	UpdateProfile(id uint, req domain.UpdateProfileRequest) (*domain.Profile, error)
	GetDashboard(id uint) (map[string]interface{}, error)
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
	AdvancePhase(roomID uint) (*domain.GameRoom, error)
	Vote(roomID, userID, targetID uint) error
	UseAbility(roomID, userID uint, ability string, targetID uint) error
}

type ShopService interface {
	ListItems() ([]domain.ShopItem, error)
	PurchaseItem(userID, itemID uint) (*domain.ShopItem, error)
	CreateItem(item domain.ShopItem) (*domain.ShopItem, error)
	UpdateItem(item domain.ShopItem) (*domain.ShopItem, error)
	DeleteItem(id uint) error
}

type AdminService interface {
        CreateRole(req domain.CreateRoleRequest) (*domain.Role, error)
        UpdateRole(id uint, req domain.CreateRoleRequest) (*domain.Role, error)
        DeleteRole(id uint) error
        ListRoles() ([]domain.Role, error)
        ListAbilities() []domain.AbilityOption
        CreateRule(req domain.RuleRequest) (*domain.GameRule, error)
        UpdateRule(id uint, req domain.RuleRequest) (*domain.GameRule, error)
        DeleteRule(id uint) error
        ListRules() ([]domain.GameRule, error)
        CreateScenario(req domain.ScenarioRequest) (*domain.Scenario, error)
	UpdateScenario(id uint, req domain.ScenarioRequest) (*domain.Scenario, error)
	DeleteScenario(id uint) error
	ListScenarios() ([]domain.Scenario, error)
}

type Services struct {
	User      UserService
	Wallet    WalletService
	Challenge ChallengeService
	Group     GroupService
	Game      GameService
	Shop      ShopService
	Admin     AdminService
}

// SFU represents the WebRTC bridge used by websocket handlers.
type SFU interface{}

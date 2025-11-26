package domain

import (
	"time"
)

type User struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	Phone      string     `json:"phone" gorm:"unique;not null"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	OTP        string     `json:"otp"`
	OTPExpires time.Time  `json:"otp_expires"`
	Role       string     `json:"role" gorm:"default:user"`
	Status     string     `json:"status" gorm:"default:active"`
	Profile    Profile    `json:"profile"`
	Wallet     Wallet     `json:"wallet"`
}

type Profile struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	UserID    uint       `json:"user_id"`
	Name      string     `json:"name"`
	Avatar    string     `json:"avatar"`
	Gender    string     `json:"gender"`
	Age       int        `json:"age"`
	Level     int        `json:"level"`
	Wins      int        `json:"wins"`
	Losses    int        `json:"losses"`
	PlayTime  int        `json:"play_time"`
	Friends   int        `json:"friends"`
	Medals    []string   `json:"medals" gorm:"serializer:json"`
}

type Role struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Name        string     `json:"name" gorm:"unique"`
	Description string     `json:"description"`
	Abilities   []string   `json:"abilities" gorm:"serializer:json"`
	Team        string     `json:"team"`
	MaxCount    int        `json:"max_count"`
}

type GameRoom struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Code      string     `json:"code" gorm:"unique"`
	Type      string     `json:"type"`
	HostID    uint       `json:"host_id"`
	Players   []User     `json:"players" gorm:"many2many:room_players;"`
	Status    string     `json:"status" gorm:"default:waiting"`
	Phase     string     `json:"phase"`
	DayCount  int        `json:"day_count"`
	Winner    string     `json:"winner"`
	Results   string     `json:"results" gorm:"type:json"`
}

type Group struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
	OwnerID   uint       `json:"owner_id"`
	Members   []User     `json:"members" gorm:"many2many:group_members;"`
	Points    int        `json:"points"`
}

type Wallet struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	UserID    uint       `json:"user_id"`
	Coins     int        `json:"coins"`
	Diamonds  int        `json:"diamonds"`
}

type Transaction struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	UserID      uint       `json:"user_id"`
	Type        string     `json:"type"`
	Amount      int        `json:"amount"`
	Currency    string     `json:"currency"`
	Description string     `json:"description"`
}

type Challenge struct {
	ID             uint       `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	RewardCoins    int        `json:"reward_coins"`
	RewardDiamonds int        `json:"reward_diamonds"`
	Target         string     `json:"target" gorm:"type:json"`
	Period         string     `json:"period"`
}

type Report struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	ReporterID uint       `json:"reporter_id"`
	TargetID   uint       `json:"target_id"`
	RoomID     uint       `json:"room_id"`
	Reason     string     `json:"reason"`
	Status     string     `json:"status" gorm:"default:pending"`
}

type Term struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Content   string     `json:"content"`
	Version   string     `json:"version"`
	Required  bool       `json:"required"`
}

type ShopItem struct {
	ID        uint                   `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	DeletedAt *time.Time             `json:"deleted_at,omitempty"`
	Name      string                 `json:"name"`
	Type      string                 `json:"type"`
	Price     int                    `json:"price"`
	Currency  string                 `json:"currency"`
	Stock     int                    `json:"stock"`
	Metadata  map[string]interface{} `json:"metadata" gorm:"serializer:json"`
}

type GameRule struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Phase       string     `json:"phase"`
	Enabled     bool       `json:"enabled"`
}

type Scenario struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Rules       []string   `json:"rules" gorm:"serializer:json"`
	Roles       []string   `json:"roles" gorm:"serializer:json"`
}

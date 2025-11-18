package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone      string `gorm:"unique;not null"`
	Username   string
	Password   string
	OTP        string
	OTPExpires time.Time
	Role       string `gorm:"default:user"`
	Status     string `gorm:"default:active"`
	Profile    Profile
	Wallet     Wallet
}

type Profile struct {
	gorm.Model
	UserID   uint
	Name     string
	Avatar   string
	Gender   string
	Age      int
	Level    int
	Wins     int
	Losses   int
	PlayTime int
	Friends  int
	Medals   []string `gorm:"serializer:json"`
}

type Role struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
	Abilities   []string `gorm:"serializer:json"`
	Team        string
	MaxCount    int
}

type GameRoom struct {
	gorm.Model
	Code     string `gorm:"unique"`
	Type     string
	HostID   uint
	Players  []User `gorm:"many2many:room_players;"`
	Status   string `gorm:"default:waiting"`
	Phase    string
	DayCount int
	Winner   string
	Results  string `gorm:"type:json"`
}

type Group struct {
	gorm.Model
	Name    string
	OwnerID uint
	Members []User `gorm:"many2many:group_members;"`
	Points  int
}

type Wallet struct {
	gorm.Model
	UserID   uint
	Coins    int
	Diamonds int
}

type Transaction struct {
	gorm.Model
	UserID      uint
	Type        string
	Amount      int
	Currency    string
	Description string
}

type Challenge struct {
	gorm.Model
	Title          string
	Description    string
	RewardCoins    int
	RewardDiamonds int
	Target         string `gorm:"type:json"`
	Period         string
}

type Report struct {
	gorm.Model
	ReporterID uint
	TargetID   uint
	RoomID     uint
	Reason     string
	Status     string `gorm:"default:pending"`
}

type Term struct {
	gorm.Model
	Content  string
	Version  string
	Required bool
}

// ShopItem represents items available for purchase in the in-game store.
type ShopItem struct {
	gorm.Model
	Name     string
	Type     string
	Price    int
	Currency string
	Stock    int
	Metadata map[string]interface{} `gorm:"serializer:json"`
}

// GameRule captures configurable rule toggles for day/night flow or special cases.
type GameRule struct {
	gorm.Model
	Name        string
	Description string
	Phase       string
	Enabled     bool
}

// Scenario lets admins bundle rules and roles into curated presets.
type Scenario struct {
	gorm.Model
	Name        string
	Description string
	Rules       []string `gorm:"serializer:json"`
	Roles       []string `gorm:"serializer:json"`
}

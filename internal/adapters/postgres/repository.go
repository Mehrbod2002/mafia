package postgres

import (
	"mafia/internal/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// New preserves the existing bootstrap signature while ensuring migrations run.
func New(dsn string) *gorm.DB {
	return NewPostgresRepository(dsn)
}

func NewPostgresRepository(dsn string) *gorm.DB {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&domain.User{}, &domain.Profile{}, &domain.Role{}, &domain.GameRoom{},
		&domain.Group{}, &domain.Wallet{}, &domain.Transaction{}, &domain.Challenge{},
		&domain.Report{}, &domain.Term{}, &domain.ShopItem{}, &domain.GameRule{}, &domain.Scenario{},
	)
	return db
}

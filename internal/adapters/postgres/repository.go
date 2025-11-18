package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mafia/internal/core/domain"
)

var db *gorm.DB

func NewPostgresRepository(dsn string) *gorm.DB {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&domain.User{}, &domain.Profile{}, &domain.Role{}, &domain.GameRoom{},
		&domain.Group{}, &domain.Wallet{}, &domain.Transaction{}, &domain.Challenge{},
		&domain.Report{}, &domain.Term{},
	)
	return db
}

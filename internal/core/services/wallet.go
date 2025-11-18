package services

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type walletService struct {
	walletRepo ports.WalletRepository
}

func NewWalletService(walletRepo ports.WalletRepository) ports.WalletService {
	return &walletService{walletRepo}
}

func (s *walletService) GetWallet(userID uint) (*domain.Wallet, error) {
	return s.walletRepo.FindByUserID(userID)
}

func (s *walletService) InitiatePurchase(userID uint, planID string) (string, error) {
	return "https://zarinpal.com/pg/StartPay/12345", nil
}

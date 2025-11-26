package services

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type walletService struct {
	walletRepo ports.WalletRepository
	payments   ports.PaymentProvider
}

func NewWalletService(walletRepo ports.WalletRepository, payments ports.PaymentProvider) ports.WalletService {
	return &walletService{walletRepo: walletRepo, payments: payments}
}

func (s *walletService) GetWallet(userID uint) (*domain.Wallet, error) {
	return s.walletRepo.FindByUserID(userID)
}

func (s *walletService) InitiatePurchase(userID uint, planID string) (string, error) {
	if s.payments != nil {
		return s.payments.CreatePaymentURL(userID, planID)
	}
	return "", nil
}

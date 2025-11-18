package services

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type challengeService struct {
	challengeRepo ports.ChallengeRepository
	userRepo      ports.UserRepository
	walletRepo    ports.WalletRepository
}

func NewChallengeService(challengeRepo ports.ChallengeRepository, userRepo ports.UserRepository, walletRepo ports.WalletRepository) ports.ChallengeService {
	return &challengeService{challengeRepo, userRepo, walletRepo}
}

func (s *challengeService) List() ([]domain.Challenge, error) {
	return s.challengeRepo.List()
}

func (s *challengeService) Complete(challengeID, userID uint) error {
	challenge, err := s.challengeRepo.FindByID(challengeID)
	if err != nil {
		return err
	}
	wallet, err := s.walletRepo.FindByUserID(userID)
	if err != nil {
		return err
	}
	wallet.Coins += challenge.RewardCoins
	wallet.Diamonds += challenge.RewardDiamonds
	return s.walletRepo.Update(wallet)
}

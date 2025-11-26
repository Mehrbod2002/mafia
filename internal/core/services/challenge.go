package services

import (
	"context"
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type challengeService struct {
	challengeRepo ports.ChallengeRepository
	userRepo      ports.UserRepository
	walletRepo    ports.WalletRepository
	events        ports.EventBus
}

func NewChallengeService(challengeRepo ports.ChallengeRepository, userRepo ports.UserRepository, walletRepo ports.WalletRepository, events ports.EventBus) ports.ChallengeService {
	return &challengeService{challengeRepo: challengeRepo, userRepo: userRepo, walletRepo: walletRepo, events: events}
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
	if err := s.walletRepo.Update(wallet); err != nil {
		return err
	}
	if s.events != nil {
		s.events.Publish(context.Background(), "challenge.completed", map[string]uint{"challenge_id": challengeID, "user_id": userID})
	}
	return nil
}

package services

import "mafia/internal/ports"

func NewServices(repos ports.Repositories, _ interface{}, _ ports.SFU) ports.Services {
	user := NewUserService(repos.User, repos.Wallet)
	wallet := NewWalletService(repos.Wallet)
	challenge := NewChallengeService(repos.Challenge, repos.User, repos.Wallet)
	group := NewGroupService(repos.Group, repos.User)
	game := NewGameService(repos.Room, repos.Role, repos.User)

	return ports.Services{
		User:      user,
		Wallet:    wallet,
		Challenge: challenge,
		Group:     group,
		Game:      game,
	}
}

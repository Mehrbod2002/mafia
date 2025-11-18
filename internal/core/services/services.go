package services

import "mafia/internal/ports"

func NewServices(repos ports.Repositories, infra ports.Infrastructure, _ ports.SFU) ports.Services {
	user := NewUserService(repos.User, repos.Wallet, infra)
	wallet := NewWalletService(repos.Wallet, infra.Payments)
	challenge := NewChallengeService(repos.Challenge, repos.User, repos.Wallet, infra.Events)
	group := NewGroupService(repos.Group, repos.User, infra.Events)
	game := NewGameService(repos.Room, repos.Role, repos.User, infra.Events)
	shop := NewShopService(repos.Shop, repos.Wallet)
	admin := NewAdminService(repos.Role, repos.Rule, repos.Scenario)

	return ports.Services{
		User:      user,
		Wallet:    wallet,
		Challenge: challenge,
		Group:     group,
		Game:      game,
		Shop:      shop,
		Admin:     admin,
	}
}

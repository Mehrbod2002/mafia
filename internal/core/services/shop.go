package services

import (
	"fmt"
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type shopService struct {
	shopRepo   ports.ShopRepository
	walletRepo ports.WalletRepository
}

func NewShopService(shopRepo ports.ShopRepository, walletRepo ports.WalletRepository) ports.ShopService {
	return &shopService{shopRepo: shopRepo, walletRepo: walletRepo}
}

func (s *shopService) ListItems() ([]domain.ShopItem, error) {
	return s.shopRepo.List()
}

func (s *shopService) PurchaseItem(userID, itemID uint) (*domain.ShopItem, error) {
	item, err := s.shopRepo.FindByID(itemID)
	if err != nil {
		return nil, err
	}
	if item.Stock == 0 {
		return nil, fmt.Errorf("out of stock")
	}
	wallet, err := s.walletRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	switch item.Currency {
	case "coins":
		if wallet.Coins < item.Price {
			return nil, fmt.Errorf("insufficient coins")
		}
		wallet.Coins -= item.Price
	default:
		if wallet.Diamonds < item.Price {
			return nil, fmt.Errorf("insufficient diamonds")
		}
		wallet.Diamonds -= item.Price
	}
	item.Stock--
	if err := s.walletRepo.Update(wallet); err != nil {
		return nil, err
	}
	if err := s.shopRepo.Update(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *shopService) CreateItem(item domain.ShopItem) (*domain.ShopItem, error) {
	if err := s.shopRepo.Create(&item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *shopService) UpdateItem(item domain.ShopItem) (*domain.ShopItem, error) {
	if err := s.shopRepo.Update(&item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *shopService) DeleteItem(id uint) error {
	return s.shopRepo.Delete(id)
}

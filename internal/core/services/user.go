package services

import (
	"crypto/rand"
	"fmt"
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret")

type userService struct {
	userRepo   ports.UserRepository
	walletRepo ports.WalletRepository
}

func NewUserService(userRepo ports.UserRepository, walletRepo ports.WalletRepository) ports.UserService {
	return &userService{userRepo, walletRepo}
}

func (s *userService) Register(phone string) error {
	_, err := s.userRepo.FindByPhone(phone)
	if err == nil {
		return fmt.Errorf("user exists")
	}
	otp := generateOTP()
	user := &domain.User{
		Phone:      phone,
		OTP:        otp,
		OTPExpires: time.Now().Add(5 * time.Minute),
	}
	return s.userRepo.Create(user)
}

func (s *userService) VerifyOTP(req domain.VerifyOTPRequest) (string, uint, error) {
	user, err := s.userRepo.FindByPhone(req.Phone)
	if err != nil || user.OTP != req.OTP || time.Now().After(user.OTPExpires) {
		return "", 0, fmt.Errorf("invalid otp")
	}
	user.OTP = ""
	s.userRepo.Update(user)

	profile := domain.Profile{UserID: user.ID, Name: req.Name, Age: req.Age, Gender: req.Gender}
	s.userRepo.Update(user)

	wallet := domain.Wallet{UserID: user.ID, Coins: 100, Diamonds: 10}
	s.walletRepo.Create(&wallet)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString(jwtKey)
	return tokenStr, user.ID, nil
}

func (s *userService) Login(phone string) error {
	_, err := s.userRepo.FindByPhone(phone)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	otp := generateOTP()
	return s.userRepo.UpdateOTP(0, otp, time.Now().Add(5*time.Minute))
}

func (s *userService) ValidateToken(token string) (uint, error) {
	claims := jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !t.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	return uint(claims["user_id"].(float64)), nil
}

func (s *userService) IsAdmin(id uint) (bool, error) {
	user, err := s.userRepo.FindByID(id)
	return user.Role == "admin", err
}

func (s *userService) GetProfile(id uint) (*domain.Profile, error) {
	user, err := s.userRepo.FindByID(id)
	return &user.Profile, err
}

func (s *userService) UpdateProfile(id uint, req domain.UpdateProfileRequest) (*domain.Profile, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	user.Profile.Name = req.Name
	user.Profile.Avatar = req.Avatar
	s.userRepo.Update(user)
	return &user.Profile, nil
}

func (s *userService) BanUser(id uint) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	user.Status = "banned"
	return s.userRepo.Update(user)
}

func (s *userService) SuspendUser(id uint) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	user.Status = "suspended"
	return s.userRepo.Update(user)
}

func generateOTP() string {
	b := make([]byte, 3)
	rand.Read(b)
	return fmt.Sprintf("%06d", int(b[0])<<16|int(b[1])<<8|int(b[2])%1000000)
}

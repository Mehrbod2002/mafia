package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret")

type userService struct {
	userRepo      ports.UserRepository
	walletRepo    ports.WalletRepository
	cache         ports.Cache
	queue         ports.Queue
	events        ports.EventBus
	notifications ports.NotificationSender
}

func NewUserService(userRepo ports.UserRepository, walletRepo ports.WalletRepository, infra ports.Infrastructure) ports.UserService {
	return &userService{userRepo: userRepo, walletRepo: walletRepo, cache: infra.Cache, queue: infra.Queue, events: infra.Events, notifications: infra.Notifications}
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
	if err := s.userRepo.Create(user); err != nil {
		return err
	}
	if s.cache != nil {
		s.cache.Set(context.Background(), otpCacheKey(phone), otp, 5*time.Minute)
	}
	if s.queue != nil {
		_ = s.queue.Enqueue(func() {
			if s.notifications != nil {
				_ = s.notifications.Send(user.ID, "sms", fmt.Sprintf("Your verification code is %s", otp))
			}
			if s.events != nil {
				s.events.Publish(context.Background(), "user.registered", user)
			}
		})
	}
	return nil
}

func (s *userService) VerifyOTP(req domain.VerifyOTPRequest) (string, uint, error) {
	user, err := s.userRepo.FindByPhone(req.Phone)
	if err != nil {
		return "", 0, fmt.Errorf("invalid otp")
	}
	cachedOTP := user.OTP
	if s.cache != nil {
		if val, ok := s.cache.Get(context.Background(), otpCacheKey(req.Phone)); ok {
			if code, ok := val.(string); ok {
				cachedOTP = code
			}
		}
	}
	if cachedOTP != req.OTP || time.Now().After(user.OTPExpires) {
		return "", 0, fmt.Errorf("invalid otp")
	}
	user.OTP = ""
	s.userRepo.Update(user)

	s.userRepo.Update(user)

	wallet := domain.Wallet{UserID: user.ID, Coins: 100, Diamonds: 10}
	s.walletRepo.Create(&wallet)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString(jwtKey)

	if s.queue != nil {
		_ = s.queue.Enqueue(func() {
			if s.events != nil {
				s.events.Publish(context.Background(), "user.verified", user)
			}
			if s.notifications != nil {
				_ = s.notifications.Send(user.ID, "in-app", "Welcome to Mafia! Your account is verified.")
			}
		})
	}

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

func (s *userService) GetDashboard(id uint) (map[string]interface{}, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	wallet, err := s.walletRepo.FindByUserID(id)
	if err != nil {
		return nil, err
	}
	summary := map[string]interface{}{
		"profile": user.Profile,
		"wallet":  wallet,
		"stats": map[string]int{
			"wins":      user.Profile.Wins,
			"losses":    user.Profile.Losses,
			"play_time": user.Profile.PlayTime,
		},
	}
	return summary, nil
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

func otpCacheKey(phone string) string {
	return fmt.Sprintf("otp:%s", phone)
}

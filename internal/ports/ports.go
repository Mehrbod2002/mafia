package ports

import "mafia/internal/core/domain"

type Repository interface{}
type Service interface{}

$(for svc in User Game Room Group Wallet Challenge Role Report Term Leaderboard League Achievement Badge Friend Block Chat Voice Stats Analytics Notification Payment; do echo "type ${svc}Repository interface { Create(*domain.$svc) error; FindByID(uint) (*domain.$svc, error); }"; done)

$(for svc in User Game Room Group Wallet Challenge Role Report Term Leaderboard League Achievement Badge Friend Block Chat Voice Stats Analytics Notification Payment; do echo "type ${svc}Service interface { /* methods */ }"; done)

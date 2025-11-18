package services

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type ${svc}Service struct {
	repo ports.${svc}Repository
	// deps
}

func New${svc}Service(repo ports.${svc}Repository) ports.${svc}Service {
	return &${svc}Service{repo}
}

// Methods stubbed

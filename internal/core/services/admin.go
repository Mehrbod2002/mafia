package services

import (
	"fmt"
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

type adminService struct {
	roleRepo     ports.RoleRepository
	ruleRepo     ports.RuleRepository
	scenarioRepo ports.ScenarioRepository
}

func NewAdminService(roleRepo ports.RoleRepository, ruleRepo ports.RuleRepository, scenarioRepo ports.ScenarioRepository) ports.AdminService {
	return &adminService{roleRepo: roleRepo, ruleRepo: ruleRepo, scenarioRepo: scenarioRepo}
}

func (s *adminService) CreateRole(req domain.CreateRoleRequest) (*domain.Role, error) {
	role := domain.Role{Name: req.Name, Description: req.Description, Abilities: req.Abilities, Team: req.Team, MaxCount: req.MaxCount}
	if err := s.roleRepo.Create(&role); err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *adminService) UpdateRole(id uint, req domain.CreateRoleRequest) (*domain.Role, error) {
	role, err := s.roleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	role.Name = req.Name
	role.Description = req.Description
	role.Abilities = req.Abilities
	role.Team = req.Team
	role.MaxCount = req.MaxCount
	if err := s.roleRepo.Update(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *adminService) DeleteRole(id uint) error {
	return s.roleRepo.Delete(id)
}

func (s *adminService) ListRoles() ([]domain.Role, error) {
	return s.roleRepo.List()
}

func (s *adminService) CreateRule(req domain.RuleRequest) (*domain.GameRule, error) {
	rule := domain.GameRule{Name: req.Name, Description: req.Description, Phase: req.Phase, Enabled: req.Enabled}
	if err := s.ruleRepo.Create(&rule); err != nil {
		return nil, err
	}
	return &rule, nil
}

func (s *adminService) UpdateRule(id uint, req domain.RuleRequest) (*domain.GameRule, error) {
	rules, err := s.ruleRepo.List()
	if err != nil {
		return nil, err
	}
	var target *domain.GameRule
	for i := range rules {
		if rules[i].ID == id {
			target = &rules[i]
			break
		}
	}
	if target == nil {
		return nil, fmt.Errorf("rule not found")
	}
	target.Name = req.Name
	target.Description = req.Description
	target.Phase = req.Phase
	target.Enabled = req.Enabled
	if err := s.ruleRepo.Update(target); err != nil {
		return nil, err
	}
	return target, nil
}

func (s *adminService) DeleteRule(id uint) error {
	return s.ruleRepo.Delete(id)
}

func (s *adminService) ListRules() ([]domain.GameRule, error) {
	return s.ruleRepo.List()
}

func (s *adminService) CreateScenario(req domain.ScenarioRequest) (*domain.Scenario, error) {
	scenario := domain.Scenario{Name: req.Name, Description: req.Description, Rules: req.Rules, Roles: req.Roles}
	if err := s.scenarioRepo.Create(&scenario); err != nil {
		return nil, err
	}
	return &scenario, nil
}

func (s *adminService) UpdateScenario(id uint, req domain.ScenarioRequest) (*domain.Scenario, error) {
	scenarios, err := s.scenarioRepo.List()
	if err != nil {
		return nil, err
	}
	var target *domain.Scenario
	for i := range scenarios {
		if scenarios[i].ID == id {
			target = &scenarios[i]
			break
		}
	}
	if target == nil {
		return nil, fmt.Errorf("scenario not found")
	}
	target.Name = req.Name
	target.Description = req.Description
	target.Rules = req.Rules
	target.Roles = req.Roles
	if err := s.scenarioRepo.Update(target); err != nil {
		return nil, err
	}
	return target, nil
}

func (s *adminService) DeleteScenario(id uint) error {
	return s.scenarioRepo.Delete(id)
}

func (s *adminService) ListScenarios() ([]domain.Scenario, error) {
	return s.scenarioRepo.List()
}

package domain

type RegisterRequest struct {
	Phone string `json:"phone" binding:"required"`
}

type VerifyOTPRequest struct {
	Phone  string `json:"phone" binding:"required"`
	OTP    string `json:"otp" binding:"required"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type LoginRequest struct {
	Phone string `json:"phone" binding:"required"`
}

type UpdateProfileRequest struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type PurchaseRequest struct {
	PlanID string `json:"plan_id"`
}

type CreateRoomRequest struct {
	Type string `json:"type"`
}

type VoteRequest struct {
	TargetID uint `json:"target_id"`
}

type AbilityRequest struct {
	Ability  string `json:"ability"`
	TargetID uint   `json:"target_id"`
}

type WSMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type PurchaseItemRequest struct {
	ItemID uint `json:"item_id" binding:"required"`
}

type CreateRoleRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Abilities   []string `json:"abilities"`
	Team        string   `json:"team"`
	MaxCount    int      `json:"max_count"`
}

type RuleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Phase       string `json:"phase"`
	Enabled     bool   `json:"enabled"`
}

type ScenarioRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Rules       []string `json:"rules"`
	Roles       []string `json:"roles"`
}

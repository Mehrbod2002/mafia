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

package payment

import "fmt"

// Provider creates payment initiation URLs.
type Provider interface {
	CreatePaymentURL(userID uint, planID string) (string, error)
}

// ZarinpalProvider crafts sandbox URLs without contacting the gateway.
type ZarinpalProvider struct {
	baseURL string
}

// NewZarinpalProvider constructs a provider with the given base URL (defaults applied when empty).
func NewZarinpalProvider(baseURL string) *ZarinpalProvider {
	if baseURL == "" {
		baseURL = "https://zarinpal.com"
	}
	return &ZarinpalProvider{baseURL: baseURL}
}

// CreatePaymentURL builds a deterministic payment URL for the client to redirect to.
func (p *ZarinpalProvider) CreatePaymentURL(userID uint, planID string) (string, error) {
	return fmt.Sprintf("%s/pg/StartPay/%d-%s", p.baseURL, userID, planID), nil
}

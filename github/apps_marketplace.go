package github

import (
	"context"
)

type MarketplaceService struct {
	client *Client

	Stubbed bool
}

type MarketplacePlan struct {
	URL                 *string `json:"url,omitempty"`
	AccountsURL         *string `json:"accounts_url,omitempty"`
	ID                  *int64  `json:"id,omitempty"`
	Number              *int    `json:"number,omitempty"`
	Name                *string `json:"name,omitempty"`
	Description         *string `json:"description,omitempty"`
	MonthlyPriceInCents *int    `json:"monthly_price_in_cents,omitempty"`
	YearlyPriceInCents  *int    `json:"yearly_price_in_cents,omitempty"`

	PriceModel *string   `json:"price_model,omitempty"`
	UnitName   *string   `json:"unit_name,omitempty"`
	Bullets    *[]string `json:"bullets,omitempty"`

	State        *string `json:"state,omitempty"`
	HasFreeTrial *bool   `json:"has_free_trial,omitempty"`
}

type MarketplacePurchase struct {
	Account *MarketplacePurchaseAccount `json:"account,omitempty"`

	BillingCycle    *string          `json:"billing_cycle,omitempty"`
	NextBillingDate *Timestamp       `json:"next_billing_date,omitempty"`
	UnitCount       *int             `json:"unit_count,omitempty"`
	Plan            *MarketplacePlan `json:"plan,omitempty"`
	OnFreeTrial     *bool            `json:"on_free_trial,omitempty"`
	FreeTrialEndsOn *Timestamp       `json:"free_trial_ends_on,omitempty"`
	UpdatedAt       *Timestamp       `json:"updated_at,omitempty"`
}

type MarketplacePendingChange struct {
	EffectiveDate *Timestamp       `json:"effective_date,omitempty"`
	UnitCount     *int             `json:"unit_count,omitempty"`
	ID            *int64           `json:"id,omitempty"`
	Plan          *MarketplacePlan `json:"plan,omitempty"`
}

type MarketplacePlanAccount struct {
	URL                      *string                   `json:"url,omitempty"`
	Type                     *string                   `json:"type,omitempty"`
	ID                       *int64                    `json:"id,omitempty"`
	Login                    *string                   `json:"login,omitempty"`
	OrganizationBillingEmail *string                   `json:"organization_billing_email,omitempty"`
	MarketplacePurchase      *MarketplacePurchase      `json:"marketplace_purchase,omitempty"`
	MarketplacePendingChange *MarketplacePendingChange `json:"marketplace_pending_change,omitempty"`
}

type MarketplacePurchaseAccount struct {
	URL                      *string `json:"url,omitempty"`
	Type                     *string `json:"type,omitempty"`
	ID                       *int64  `json:"id,omitempty"`
	Login                    *string `json:"login,omitempty"`
	OrganizationBillingEmail *string `json:"organization_billing_email,omitempty"`
	Email                    *string `json:"email,omitempty"`
	NodeID                   *string `json:"node_id,omitempty"`
}

//meta:operation GET /marketplace_listing/plans
//meta:operation GET /marketplace_listing/stubbed/plans
func (s *MarketplaceService) ListPlans(ctx context.Context, opts *ListOptions) ([]*MarketplacePlan, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /marketplace_listing/plans/{plan_id}/accounts
//meta:operation GET /marketplace_listing/stubbed/plans/{plan_id}/accounts
func (s *MarketplaceService) ListPlanAccountsForPlan(ctx context.Context, planID int64, opts *ListOptions) ([]*MarketplacePlanAccount, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /marketplace_listing/accounts/{account_id}
//meta:operation GET /marketplace_listing/stubbed/accounts/{account_id}
func (s *MarketplaceService) GetPlanAccountForAccount(ctx context.Context, accountID int64) (*MarketplacePlanAccount, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/marketplace_purchases
//meta:operation GET /user/marketplace_purchases/stubbed
func (s *MarketplaceService) ListMarketplacePurchasesForUser(ctx context.Context, opts *ListOptions) ([]*MarketplacePurchase, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *MarketplaceService) marketplaceURI(endpoint string) string {
	_ = "STUB: not implemented"
	return ""
}

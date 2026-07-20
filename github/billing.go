package github

import (
	"context"
)

type BillingService service

type MinutesUsedBreakdown = map[string]int

type PackagesBilling struct {
	TotalGigabytesBandwidthUsed     int `json:"total_gigabytes_bandwidth_used"`
	TotalPaidGigabytesBandwidthUsed int `json:"total_paid_gigabytes_bandwidth_used"`
	IncludedGigabytesBandwidth      int `json:"included_gigabytes_bandwidth"`
}

type StorageBilling struct {
	DaysLeftInBillingCycle       int `json:"days_left_in_billing_cycle"`
	EstimatedPaidStorageForMonth int `json:"estimated_paid_storage_for_month"`
	EstimatedStorageForMonth     int `json:"estimated_storage_for_month"`
}

type ActiveCommittersListOptions struct {
	AdvancedSecurityProduct *string `url:"advanced_security_product,omitempty"`

	ListOptions
}

type ActiveCommitters struct {
	TotalAdvancedSecurityCommitters     *int                          `json:"total_advanced_security_committers,omitempty"`
	TotalCount                          *int                          `json:"total_count,omitempty"`
	MaximumAdvancedSecurityCommitters   *int                          `json:"maximum_advanced_security_committers,omitempty"`
	PurchasedAdvancedSecurityCommitters *int                          `json:"purchased_advanced_security_committers,omitempty"`
	Repositories                        []*RepositoryActiveCommitters `json:"repositories"`
}

type RepositoryActiveCommitters struct {
	Name                                string                                 `json:"name"`
	AdvancedSecurityCommitters          int                                    `json:"advanced_security_committers"`
	AdvancedSecurityCommittersBreakdown []*AdvancedSecurityCommittersBreakdown `json:"advanced_security_committers_breakdown"`
}

type AdvancedSecurityCommittersBreakdown struct {
	UserLogin       string `json:"user_login"`
	LastPushedDate  string `json:"last_pushed_date"`
	LastPushedEmail string `json:"last_pushed_email"`
}

type UsageReportOptions struct {
	Year *int `url:"year,omitempty"`

	Month *int `url:"month,omitempty"`

	Day *int `url:"day,omitempty"`

	Hour *int `url:"hour,omitempty"`
}

type PremiumRequestUsageReportOptions struct {
	Year *int `url:"year,omitempty"`

	Month *int `url:"month,omitempty"`

	Day *int `url:"day,omitempty"`

	User *string `url:"user,omitempty"`

	Model *string `url:"model,omitempty"`

	Product *string `url:"product,omitempty"`
}

type UsageItem struct {
	Date           string  `json:"date"`
	Product        string  `json:"product"`
	SKU            string  `json:"sku"`
	Quantity       float64 `json:"quantity"`
	UnitType       string  `json:"unitType"`
	PricePerUnit   float64 `json:"pricePerUnit"`
	GrossAmount    float64 `json:"grossAmount"`
	DiscountAmount float64 `json:"discountAmount"`
	NetAmount      float64 `json:"netAmount"`
	RepositoryName *string `json:"repositoryName,omitempty"`

	OrganizationName *string `json:"organizationName,omitempty"`
}

type UsageReport struct {
	UsageItems []*UsageItem `json:"usageItems,omitempty"`
}

type PremiumRequestUsageItem struct {
	Product          string  `json:"product"`
	SKU              string  `json:"sku"`
	Model            string  `json:"model"`
	UnitType         string  `json:"unitType"`
	PricePerUnit     float64 `json:"pricePerUnit"`
	GrossQuantity    float64 `json:"grossQuantity"`
	GrossAmount      float64 `json:"grossAmount"`
	DiscountQuantity float64 `json:"discountQuantity"`
	DiscountAmount   float64 `json:"discountAmount"`
	NetQuantity      float64 `json:"netQuantity"`
	NetAmount        float64 `json:"netAmount"`
}

type PremiumRequestUsageTimePeriod struct {
	Year  int  `json:"year"`
	Month *int `json:"month,omitempty"`
	Day   *int `json:"day,omitempty"`
}

type PremiumRequestUsageReport struct {
	TimePeriod PremiumRequestUsageTimePeriod `json:"timePeriod"`

	Organization *string `json:"organization,omitempty"`

	User       *string                    `json:"user,omitempty"`
	Product    *string                    `json:"product,omitempty"`
	Model      *string                    `json:"model,omitempty"`
	UsageItems []*PremiumRequestUsageItem `json:"usageItems"`
}

//meta:operation GET /orgs/{org}/settings/billing/packages
func (s *BillingService) GetOrganizationPackagesBilling(ctx context.Context, org string) (*PackagesBilling, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/settings/billing/shared-storage
func (s *BillingService) GetOrganizationStorageBilling(ctx context.Context, org string) (*StorageBilling, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/settings/billing/advanced-security
func (s *BillingService) GetOrganizationAdvancedSecurityActiveCommitters(ctx context.Context, org string, opts *ActiveCommittersListOptions) (*ActiveCommitters, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/settings/billing/packages
func (s *BillingService) GetPackagesBilling(ctx context.Context, user string) (*PackagesBilling, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/settings/billing/shared-storage
func (s *BillingService) GetStorageBilling(ctx context.Context, user string) (*StorageBilling, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{org}/settings/billing/usage
func (s *BillingService) GetOrganizationUsageReport(ctx context.Context, org string, opts *UsageReportOptions) (*UsageReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/settings/billing/usage
func (s *BillingService) GetUsageReport(ctx context.Context, user string, opts *UsageReportOptions) (*UsageReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{org}/settings/billing/premium_request/usage
func (s *BillingService) GetOrganizationPremiumRequestUsageReport(ctx context.Context, org string, opts *PremiumRequestUsageReportOptions) (*PremiumRequestUsageReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/settings/billing/premium_request/usage
func (s *BillingService) GetPremiumRequestUsageReport(ctx context.Context, user string, opts *PremiumRequestUsageReportOptions) (*PremiumRequestUsageReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{org}/settings/billing/ai_credit/usage
func (s *BillingService) GetOrgAICreditUsage(ctx context.Context, org string, opts *PremiumRequestUsageReportOptions) (*PremiumRequestUsageReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/settings/billing/ai_credit/usage
func (s *BillingService) GetUserAICreditUsage(ctx context.Context, username string, opts *PremiumRequestUsageReportOptions) (*PremiumRequestUsageReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

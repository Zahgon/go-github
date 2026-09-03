package github

import (
	"context"
)

type EnterpriseUsageReportOptions struct {
	Year         int    `url:"year,omitempty"`
	Month        int    `url:"month,omitempty"`
	Day          int    `url:"day,omitempty"`
	CostCenterID string `url:"cost_center_id,omitempty"`
}

type EnterprisePremiumRequestUsageReportOptions struct {
	EnterpriseUsageReportOptions
	Organization string `url:"organization,omitempty"`
	User         string `url:"user,omitempty"`
	Model        string `url:"model,omitempty"`
	Product      string `url:"product,omitempty"`
}

type EnterpriseUsageSummaryOptions struct {
	EnterpriseUsageReportOptions
	Organization string `url:"organization,omitempty"`
	Repository   string `url:"repository,omitempty"`
	Product      string `url:"product,omitempty"`
	SKU          string `url:"sku,omitempty"`
}

type EnterpriseUsageTimePeriod struct {
	Year  int  `json:"year"`
	Month *int `json:"month,omitempty"`
	Day   *int `json:"day,omitempty"`
}

type EnterpriseAggregatedUsageItem struct {
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

type EnterpriseAggregatedUsageReport struct {
	TimePeriod   EnterpriseUsageTimePeriod        `json:"timePeriod"`
	Enterprise   string                           `json:"enterprise"`
	Organization *string                          `json:"organization,omitempty"`
	User         *string                          `json:"user,omitempty"`
	Product      *string                          `json:"product,omitempty"`
	Model        *string                          `json:"model,omitempty"`
	CostCenter   *BillingCostCenter               `json:"costCenter,omitempty"`
	UsageItems   []*EnterpriseAggregatedUsageItem `json:"usageItems"`
}

type EnterpriseUsageSummaryItem struct {
	Product          string  `json:"product"`
	SKU              string  `json:"sku"`
	UnitType         string  `json:"unitType"`
	PricePerUnit     float64 `json:"pricePerUnit"`
	GrossQuantity    float64 `json:"grossQuantity"`
	GrossAmount      float64 `json:"grossAmount"`
	DiscountQuantity float64 `json:"discountQuantity"`
	DiscountAmount   float64 `json:"discountAmount"`
	NetQuantity      float64 `json:"netQuantity"`
	NetAmount        float64 `json:"netAmount"`
}

type EnterpriseUsageSummaryReport struct {
	TimePeriod   EnterpriseUsageTimePeriod     `json:"timePeriod"`
	Enterprise   string                        `json:"enterprise"`
	Organization *string                       `json:"organization,omitempty"`
	Repository   *string                       `json:"repository,omitempty"`
	Product      *string                       `json:"product,omitempty"`
	SKU          *string                       `json:"sku,omitempty"`
	CostCenter   *BillingCostCenter            `json:"costCenter,omitempty"`
	UsageItems   []*EnterpriseUsageSummaryItem `json:"usageItems"`
}

type BillingCostCenter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EnterpriseUsageItem struct {
	Date             string  `json:"date"`
	Product          string  `json:"product"`
	SKU              string  `json:"sku"`
	Quantity         float64 `json:"quantity"`
	UnitType         string  `json:"unitType"`
	PricePerUnit     float64 `json:"pricePerUnit"`
	GrossAmount      float64 `json:"grossAmount"`
	DiscountAmount   float64 `json:"discountAmount"`
	NetAmount        float64 `json:"netAmount"`
	OrganizationName string  `json:"organizationName"`
	RepositoryName   *string `json:"repositoryName,omitempty"`
}

type EnterpriseUsageReport struct {
	UsageItems []*EnterpriseUsageItem `json:"usageItems,omitempty"`
}

//meta:operation GET /enterprises/{enterprise}/settings/billing/usage
func (s *EnterpriseService) GetUsageReport(ctx context.Context, enterprise string, opts *EnterpriseUsageReportOptions) (*EnterpriseUsageReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/settings/billing/usage/summary
func (s *EnterpriseService) GetUsageSummary(ctx context.Context, enterprise string, opts *EnterpriseUsageSummaryOptions) (*EnterpriseUsageSummaryReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/settings/billing/premium_request/usage
func (s *EnterpriseService) GetPremiumRequestUsageReport(ctx context.Context, enterprise string, opts *EnterprisePremiumRequestUsageReportOptions) (*EnterpriseAggregatedUsageReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/settings/billing/ai_credit/usage
func (s *EnterpriseService) GetAICreditUsage(ctx context.Context, enterprise string, opts *EnterprisePremiumRequestUsageReportOptions) (*EnterpriseAggregatedUsageReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

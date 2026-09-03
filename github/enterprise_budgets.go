package github

import (
	"context"
)

const (
	BudgetScopeEnterprise   = "enterprise"
	BudgetScopeOrganization = "organization"
	BudgetScopeRepository   = "repository"
	BudgetScopeCostCenter   = "cost_center"
)

const (
	BudgetTypeProductPricing = "ProductPricing"
	BudgetTypeSkuPricing     = "SkuPricing"
)

type EnterpriseBudgetAlerting struct {
	WillAlert       *bool    `json:"will_alert,omitempty"`
	AlertRecipients []string `json:"alert_recipients,omitempty"`
}

type EnterpriseBudget struct {
	ID                  *string                   `json:"id,omitempty"`
	BudgetType          *string                   `json:"budget_type,omitempty"`
	BudgetProductSKU    *string                   `json:"budget_product_sku,omitempty"`
	BudgetScope         *string                   `json:"budget_scope,omitempty"`
	BudgetEntityName    *string                   `json:"budget_entity_name,omitempty"`
	BudgetAmount        *int                      `json:"budget_amount,omitempty"`
	PreventFurtherUsage *bool                     `json:"prevent_further_usage,omitempty"`
	BudgetAlerting      *EnterpriseBudgetAlerting `json:"budget_alerting,omitempty"`
}

func (b EnterpriseBudget) String() string { _ = "STUB: not implemented"; return "" }

type EnterpriseListBudgets struct {
	Budgets     []*EnterpriseBudget `json:"budgets"`
	HasNextPage *bool               `json:"has_next_page,omitempty"`
	TotalCount  *int                `json:"total_count,omitempty"`
}

type EnterpriseCreateBudget struct {
	BudgetAmount        int                       `json:"budget_amount"`
	PreventFurtherUsage bool                      `json:"prevent_further_usage"`
	BudgetAlerting      *EnterpriseBudgetAlerting `json:"budget_alerting"`
	BudgetScope         string                    `json:"budget_scope"`
	BudgetEntityName    *string                   `json:"budget_entity_name,omitempty"`
	BudgetType          string                    `json:"budget_type"`
	BudgetProductSKU    *string                   `json:"budget_product_sku,omitempty"`
}

type EnterpriseUpdateBudget struct {
	BudgetAmount        *int                      `json:"budget_amount,omitempty"`
	PreventFurtherUsage *bool                     `json:"prevent_further_usage,omitempty"`
	BudgetAlerting      *EnterpriseBudgetAlerting `json:"budget_alerting,omitempty"`
	BudgetScope         *string                   `json:"budget_scope,omitempty"`
	BudgetEntityName    *string                   `json:"budget_entity_name,omitempty"`
	BudgetType          *string                   `json:"budget_type,omitempty"`
	BudgetProductSKU    *string                   `json:"budget_product_sku,omitempty"`
}

type EnterpriseCreateOrUpdateBudgetResponse struct {
	Message string            `json:"message"`
	Budget  *EnterpriseBudget `json:"budget"`
}

type EnterpriseDeleteBudgetResponse struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

//meta:operation GET /enterprises/{enterprise}/settings/billing/budgets
func (s *EnterpriseService) ListBudgets(ctx context.Context, enterprise string) (*EnterpriseListBudgets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/settings/billing/budgets
func (s *EnterpriseService) CreateBudget(ctx context.Context, enterprise string, body EnterpriseCreateBudget) (*EnterpriseCreateOrUpdateBudgetResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/settings/billing/budgets/{budget_id}
func (s *EnterpriseService) GetBudget(ctx context.Context, enterprise, budgetID string) (*EnterpriseBudget, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/settings/billing/budgets/{budget_id}
func (s *EnterpriseService) UpdateBudget(ctx context.Context, enterprise, budgetID string, body EnterpriseUpdateBudget) (*EnterpriseCreateOrUpdateBudgetResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/settings/billing/budgets/{budget_id}
func (s *EnterpriseService) DeleteBudget(ctx context.Context, enterprise, budgetID string) (*EnterpriseDeleteBudgetResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

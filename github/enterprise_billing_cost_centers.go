package github

import (
	"context"
)

type CostCenter struct {
	ID                string                `json:"id"`
	Name              string                `json:"name"`
	Resources         []*CostCenterResource `json:"resources"`
	State             *string               `json:"state,omitempty"`
	AzureSubscription *string               `json:"azure_subscription,omitempty"`
}

type CostCenterResource struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type CostCenters struct {
	CostCenters []*CostCenter `json:"costCenters,omitempty"`
}

type ListCostCenterOptions struct {
	State *string `url:"state,omitempty"`
}

type CostCenterRequest struct {
	Name string `json:"name"`
}

type CostCenterResourceRequest struct {
	Users         []string `json:"users,omitempty"`
	Organizations []string `json:"organizations,omitempty"`
	Repositories  []string `json:"repositories,omitempty"`
}

type AddResourcesToCostCenterResponse struct {
	Message             *string               `json:"message,omitempty"`
	ReassignedResources []*ReassignedResource `json:"reassigned_resources,omitempty"`
}

type ReassignedResource struct {
	ResourceType       *string `json:"resource_type,omitempty"`
	Name               *string `json:"name,omitempty"`
	PreviousCostCenter *string `json:"previous_cost_center,omitempty"`
}

type RemoveResourcesFromCostCenterResponse struct {
	Message *string `json:"message,omitempty"`
}

type DeleteCostCenterResponse struct {
	Message         string `json:"message"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	CostCenterState string `json:"costCenterState"`
}

//meta:operation GET /enterprises/{enterprise}/settings/billing/cost-centers
func (s *EnterpriseService) ListCostCenters(ctx context.Context, enterprise string, opts *ListCostCenterOptions) (*CostCenters, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/settings/billing/cost-centers
func (s *EnterpriseService) CreateCostCenter(ctx context.Context, enterprise string, body CostCenterRequest) (*CostCenter, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/settings/billing/cost-centers/{cost_center_id}
func (s *EnterpriseService) GetCostCenter(ctx context.Context, enterprise, costCenterID string) (*CostCenter, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/settings/billing/cost-centers/{cost_center_id}
func (s *EnterpriseService) UpdateCostCenter(ctx context.Context, enterprise, costCenterID string, body CostCenterRequest) (*CostCenter, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/settings/billing/cost-centers/{cost_center_id}
func (s *EnterpriseService) DeleteCostCenter(ctx context.Context, enterprise, costCenterID string) (*DeleteCostCenterResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/settings/billing/cost-centers/{cost_center_id}/resource
func (s *EnterpriseService) AddResourcesToCostCenter(ctx context.Context, enterprise, costCenterID string, body CostCenterResourceRequest) (*AddResourcesToCostCenterResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/settings/billing/cost-centers/{cost_center_id}/resource
func (s *EnterpriseService) RemoveResourcesFromCostCenter(ctx context.Context, enterprise, costCenterID string, resources CostCenterResourceRequest) (*RemoveResourcesFromCostCenterResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

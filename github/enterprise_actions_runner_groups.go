package github

import (
	"context"
)

type ListOrganizations struct {
	TotalCount    *int            `json:"total_count,omitempty"`
	Organizations []*Organization `json:"organizations"`
}

type EnterpriseRunnerGroup struct {
	ID                           *int64   `json:"id,omitempty"`
	Name                         *string  `json:"name,omitempty"`
	Visibility                   *string  `json:"visibility,omitempty"`
	Default                      *bool    `json:"default,omitempty"`
	SelectedOrganizationsURL     *string  `json:"selected_organizations_url,omitempty"`
	RunnersURL                   *string  `json:"runners_url,omitempty"`
	HostedRunnersURL             *string  `json:"hosted_runners_url,omitempty"`
	NetworkConfigurationID       *string  `json:"network_configuration_id,omitempty"`
	Inherited                    *bool    `json:"inherited,omitempty"`
	AllowsPublicRepositories     *bool    `json:"allows_public_repositories,omitempty"`
	RestrictedToWorkflows        *bool    `json:"restricted_to_workflows,omitempty"`
	SelectedWorkflows            []string `json:"selected_workflows,omitempty"`
	WorkflowRestrictionsReadOnly *bool    `json:"workflow_restrictions_read_only,omitempty"`
}

type EnterpriseRunnerGroups struct {
	TotalCount   *int                     `json:"total_count,omitempty"`
	RunnerGroups []*EnterpriseRunnerGroup `json:"runner_groups"`
}

type CreateEnterpriseRunnerGroupRequest struct {
	Name       *string `json:"name,omitempty"`
	Visibility *string `json:"visibility,omitempty"`

	SelectedOrganizationIDs []int64 `json:"selected_organization_ids,omitempty"`

	Runners []int64 `json:"runners,omitempty"`

	AllowsPublicRepositories *bool `json:"allows_public_repositories,omitempty"`

	RestrictedToWorkflows *bool `json:"restricted_to_workflows,omitempty"`

	SelectedWorkflows []string `json:"selected_workflows,omitempty"`

	NetworkConfigurationID *string `json:"network_configuration_id,omitempty"`
}

type UpdateEnterpriseRunnerGroupRequest struct {
	Name                     *string  `json:"name,omitempty"`
	Visibility               *string  `json:"visibility,omitempty"`
	AllowsPublicRepositories *bool    `json:"allows_public_repositories,omitempty"`
	RestrictedToWorkflows    *bool    `json:"restricted_to_workflows,omitempty"`
	SelectedWorkflows        []string `json:"selected_workflows,omitempty"`
	NetworkConfigurationID   *string  `json:"network_configuration_id,omitempty"`
}

type SetOrgAccessRunnerGroupRequest struct {
	SelectedOrganizationIDs []int64 `json:"selected_organization_ids"`
}

type ListEnterpriseRunnerGroupOptions struct {
	ListOptions

	VisibleToOrganization string `url:"visible_to_organization,omitempty"`
}

//meta:operation GET /enterprises/{enterprise}/actions/runner-groups
func (s *EnterpriseService) ListRunnerGroups(ctx context.Context, enterprise string, opts *ListEnterpriseRunnerGroupOptions) (*EnterpriseRunnerGroups, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}
func (s *EnterpriseService) GetEnterpriseRunnerGroup(ctx context.Context, enterprise string, groupID int64) (*EnterpriseRunnerGroup, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}
func (s *EnterpriseService) DeleteEnterpriseRunnerGroup(ctx context.Context, enterprise string, groupID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /enterprises/{enterprise}/actions/runner-groups
func (s *EnterpriseService) CreateEnterpriseRunnerGroup(ctx context.Context, enterprise string, body CreateEnterpriseRunnerGroupRequest) (*EnterpriseRunnerGroup, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}
func (s *EnterpriseService) UpdateEnterpriseRunnerGroup(ctx context.Context, enterprise string, groupID int64, body UpdateEnterpriseRunnerGroupRequest) (*EnterpriseRunnerGroup, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations
func (s *EnterpriseService) ListOrganizationAccessRunnerGroup(ctx context.Context, enterprise string, groupID int64, opts *ListOptions) (*ListOrganizations, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations
func (s *EnterpriseService) SetOrganizationAccessRunnerGroup(ctx context.Context, enterprise string, groupID int64, body SetOrgAccessRunnerGroupRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations/{org_id}
func (s *EnterpriseService) AddOrganizationAccessRunnerGroup(ctx context.Context, enterprise string, groupID, orgID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations/{org_id}
func (s *EnterpriseService) RemoveOrganizationAccessRunnerGroup(ctx context.Context, enterprise string, groupID, orgID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners
func (s *EnterpriseService) ListRunnerGroupRunners(ctx context.Context, enterprise string, groupID int64, opts *ListOptions) (*Runners, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners
func (s *EnterpriseService) SetRunnerGroupRunners(ctx context.Context, enterprise string, groupID int64, body SetRunnerGroupRunnersRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners/{runner_id}
func (s *EnterpriseService) AddRunnerGroupRunners(ctx context.Context, enterprise string, groupID, runnerID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners/{runner_id}
func (s *EnterpriseService) RemoveRunnerGroupRunners(ctx context.Context, enterprise string, groupID, runnerID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

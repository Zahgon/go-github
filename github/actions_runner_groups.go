package github

import (
	"context"
)

type RunnerGroup struct {
	ID                           *int64   `json:"id,omitempty"`
	Name                         *string  `json:"name,omitempty"`
	Visibility                   *string  `json:"visibility,omitempty"`
	Default                      *bool    `json:"default,omitempty"`
	SelectedRepositoriesURL      *string  `json:"selected_repositories_url,omitempty"`
	RunnersURL                   *string  `json:"runners_url,omitempty"`
	HostedRunnersURL             *string  `json:"hosted_runners_url,omitempty"`
	NetworkConfigurationID       *string  `json:"network_configuration_id,omitempty"`
	Inherited                    *bool    `json:"inherited,omitempty"`
	AllowsPublicRepositories     *bool    `json:"allows_public_repositories,omitempty"`
	RestrictedToWorkflows        *bool    `json:"restricted_to_workflows,omitempty"`
	SelectedWorkflows            []string `json:"selected_workflows,omitempty"`
	WorkflowRestrictionsReadOnly *bool    `json:"workflow_restrictions_read_only,omitempty"`
}

type RunnerGroups struct {
	TotalCount   int            `json:"total_count"`
	RunnerGroups []*RunnerGroup `json:"runner_groups"`
}

type CreateRunnerGroupRequest struct {
	Name       *string `json:"name,omitempty"`
	Visibility *string `json:"visibility,omitempty"`

	SelectedRepositoryIDs []int64 `json:"selected_repository_ids,omitempty"`

	Runners []int64 `json:"runners,omitempty"`

	AllowsPublicRepositories *bool `json:"allows_public_repositories,omitempty"`

	RestrictedToWorkflows *bool `json:"restricted_to_workflows,omitempty"`

	SelectedWorkflows []string `json:"selected_workflows,omitempty"`

	NetworkConfigurationID *string `json:"network_configuration_id,omitempty"`
}

type UpdateRunnerGroupRequest struct {
	Name                     *string  `json:"name,omitempty"`
	Visibility               *string  `json:"visibility,omitempty"`
	AllowsPublicRepositories *bool    `json:"allows_public_repositories,omitempty"`
	RestrictedToWorkflows    *bool    `json:"restricted_to_workflows,omitempty"`
	SelectedWorkflows        []string `json:"selected_workflows,omitempty"`
	NetworkConfigurationID   *string  `json:"network_configuration_id,omitempty"`
}

type SetRepoAccessRunnerGroupRequest struct {
	SelectedRepositoryIDs []int64 `json:"selected_repository_ids"`
}

type SetRunnerGroupRunnersRequest struct {
	Runners []int64 `json:"runners"`
}

type ListOrgRunnerGroupOptions struct {
	ListOptions

	VisibleToRepository string `url:"visible_to_repository,omitempty"`
}

//meta:operation GET /orgs/{org}/actions/runner-groups
func (s *ActionsService) ListOrganizationRunnerGroups(ctx context.Context, org string, opts *ListOrgRunnerGroupOptions) (*RunnerGroups, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/runner-groups/{runner_group_id}
func (s *ActionsService) GetOrganizationRunnerGroup(ctx context.Context, org string, groupID int64) (*RunnerGroup, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/runner-groups/{runner_group_id}
func (s *ActionsService) DeleteOrganizationRunnerGroup(ctx context.Context, org string, groupID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /orgs/{org}/actions/runner-groups
func (s *ActionsService) CreateOrganizationRunnerGroup(ctx context.Context, org string, body CreateRunnerGroupRequest) (*RunnerGroup, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/actions/runner-groups/{runner_group_id}
func (s *ActionsService) UpdateOrganizationRunnerGroup(ctx context.Context, org string, groupID int64, body UpdateRunnerGroupRequest) (*RunnerGroup, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/runner-groups/{runner_group_id}/repositories
func (s *ActionsService) ListRepositoryAccessRunnerGroup(ctx context.Context, org string, groupID int64, opts *ListOptions) (*ListRepositories, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/runner-groups/{runner_group_id}/repositories
func (s *ActionsService) SetRepositoryAccessRunnerGroup(ctx context.Context, org string, groupID int64, body SetRepoAccessRunnerGroupRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/actions/runner-groups/{runner_group_id}/repositories/{repository_id}
func (s *ActionsService) AddRepositoryAccessRunnerGroup(ctx context.Context, org string, groupID, repoID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/runner-groups/{runner_group_id}/repositories/{repository_id}
func (s *ActionsService) RemoveRepositoryAccessRunnerGroup(ctx context.Context, org string, groupID, repoID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/actions/runner-groups/{runner_group_id}/hosted-runners
func (s *ActionsService) ListRunnerGroupHostedRunners(ctx context.Context, org string, groupID int64, opts *ListOptions) (*HostedRunners, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/runner-groups/{runner_group_id}/runners
func (s *ActionsService) ListRunnerGroupRunners(ctx context.Context, org string, groupID int64, opts *ListOptions) (*Runners, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/runner-groups/{runner_group_id}/runners
func (s *ActionsService) SetRunnerGroupRunners(ctx context.Context, org string, groupID int64, body SetRunnerGroupRunnersRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/actions/runner-groups/{runner_group_id}/runners/{runner_id}
func (s *ActionsService) AddRunnerGroupRunners(ctx context.Context, org string, groupID, runnerID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/runner-groups/{runner_group_id}/runners/{runner_id}
func (s *ActionsService) RemoveRunnerGroupRunners(ctx context.Context, org string, groupID, runnerID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

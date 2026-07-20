package github

import (
	"context"
)

type ActionsPermissions struct {
	EnabledRepositories *string `json:"enabled_repositories,omitempty"`
	AllowedActions      *string `json:"allowed_actions,omitempty"`
	SelectedActionsURL  *string `json:"selected_actions_url,omitempty"`
	SHAPinningRequired  *bool   `json:"sha_pinning_required,omitempty"`
}

func (a ActionsPermissions) String() string { _ = "STUB: not implemented"; return "" }

type ActionsEnabledOnOrgRepos struct {
	TotalCount   int           `json:"total_count"`
	Repositories []*Repository `json:"repositories"`
}

type ActionsAllowed struct {
	GithubOwnedAllowed *bool `json:"github_owned_allowed,omitempty"`
	VerifiedAllowed    *bool `json:"verified_allowed,omitempty"`

	PatternsAllowed []string `json:"patterns_allowed,omitzero"`
}

func (a ActionsAllowed) String() string { _ = "STUB: not implemented"; return "" }

type DefaultWorkflowPermissionOrganization struct {
	DefaultWorkflowPermissions   *string `json:"default_workflow_permissions,omitempty"`
	CanApprovePullRequestReviews *bool   `json:"can_approve_pull_request_reviews,omitempty"`
}

type SelfHostedRunnersSettingsOrganization struct {
	EnabledRepositories     *string `json:"enabled_repositories,omitempty"`
	SelectedRepositoriesURL *string `json:"selected_repositories_url,omitempty"`
}

func (s SelfHostedRunnersSettingsOrganization) String() string {
	_ = "STUB: not implemented"
	return ""
}

type SelfHostedRunnersSettingsOrganizationOpt struct {
	EnabledRepositories *string `json:"enabled_repositories,omitempty"`
}

//meta:operation GET /orgs/{org}/actions/permissions
func (s *ActionsService) GetActionsPermissions(ctx context.Context, org string) (*ActionsPermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions
func (s *ActionsService) UpdateActionsPermissions(ctx context.Context, org string, body ActionsPermissions) (*ActionsPermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/permissions/repositories
func (s *ActionsService) ListEnabledReposInOrg(ctx context.Context, owner string, opts *ListOptions) (*ActionsEnabledOnOrgRepos, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/repositories
func (s *ActionsService) SetEnabledReposInOrg(ctx context.Context, owner string, repositoryIDs []int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/repositories/{repository_id}
func (s *ActionsService) AddEnabledReposInOrg(ctx context.Context, owner string, repositoryID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/permissions/repositories/{repository_id}
func (s *ActionsService) RemoveEnabledReposInOrg(ctx context.Context, owner string, repositoryID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/actions/permissions/selected-actions
func (s *ActionsService) GetActionsAllowed(ctx context.Context, org string) (*ActionsAllowed, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/selected-actions
func (s *ActionsService) UpdateActionsAllowed(ctx context.Context, org string, body ActionsAllowed) (*ActionsAllowed, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/permissions/workflow
func (s *ActionsService) GetDefaultWorkflowPermissionsInOrganization(ctx context.Context, org string) (*DefaultWorkflowPermissionOrganization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/workflow
func (s *ActionsService) UpdateDefaultWorkflowPermissionsInOrganization(ctx context.Context, org string, body DefaultWorkflowPermissionOrganization) (*DefaultWorkflowPermissionOrganization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/permissions/artifact-and-log-retention
func (s *ActionsService) GetArtifactAndLogRetentionPeriodInOrganization(ctx context.Context, org string) (*ArtifactPeriod, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/artifact-and-log-retention
func (s *ActionsService) UpdateArtifactAndLogRetentionPeriodInOrganization(ctx context.Context, org string, body ArtifactPeriodOpt) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/actions/permissions/self-hosted-runners
func (s *ActionsService) GetSelfHostedRunnersSettingsInOrganization(ctx context.Context, org string) (*SelfHostedRunnersSettingsOrganization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/self-hosted-runners
func (s *ActionsService) UpdateSelfHostedRunnersSettingsInOrganization(ctx context.Context, org string, body SelfHostedRunnersSettingsOrganizationOpt) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SelfHostedRunnersAllowedRepos struct {
	TotalCount   int           `json:"total_count"`
	Repositories []*Repository `json:"repositories"`
}

//meta:operation GET /orgs/{org}/actions/permissions/self-hosted-runners/repositories
func (s *ActionsService) ListRepositoriesSelfHostedRunnersAllowedInOrganization(ctx context.Context, org string, opts *ListOptions) (*SelfHostedRunnersAllowedRepos, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/self-hosted-runners/repositories
func (s *ActionsService) SetRepositoriesSelfHostedRunnersAllowedInOrganization(ctx context.Context, org string, repositoryIDs []int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/self-hosted-runners/repositories/{repository_id}
func (s *ActionsService) AddRepositorySelfHostedRunnersAllowedInOrganization(ctx context.Context, org string, repositoryID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/permissions/self-hosted-runners/repositories/{repository_id}
func (s *ActionsService) RemoveRepositorySelfHostedRunnersAllowedInOrganization(ctx context.Context, org string, repositoryID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/actions/permissions/fork-pr-workflows-private-repos
func (s *ActionsService) GetPrivateRepoForkPRWorkflowSettingsInOrganization(ctx context.Context, org string) (*WorkflowsPermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/fork-pr-workflows-private-repos
func (s *ActionsService) UpdatePrivateRepoForkPRWorkflowSettingsInOrganization(ctx context.Context, org string, body *WorkflowsPermissionsOpt) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/actions/permissions/fork-pr-contributor-approval
func (s *ActionsService) GetOrganizationForkPRContributorApprovalPermissions(ctx context.Context, org string) (*ContributorApprovalPermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/fork-pr-contributor-approval
func (s *ActionsService) UpdateOrganizationForkPRContributorApprovalPermissions(ctx context.Context, org string, body ContributorApprovalPermissions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

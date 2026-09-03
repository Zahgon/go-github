package github

import (
	"context"
)

type ActionsEnabledOnEnterpriseRepos struct {
	TotalCount    int             `json:"total_count"`
	Organizations []*Organization `json:"organizations"`
}

type ActionsPermissionsEnterprise struct {
	EnabledOrganizations *string `json:"enabled_organizations,omitempty"`
	AllowedActions       *string `json:"allowed_actions,omitempty"`
	SelectedActionsURL   *string `json:"selected_actions_url,omitempty"`
}

func (a ActionsPermissionsEnterprise) String() string { _ = "STUB: not implemented"; return "" }

type DefaultWorkflowPermissionEnterprise struct {
	DefaultWorkflowPermissions   *string `json:"default_workflow_permissions,omitempty"`
	CanApprovePullRequestReviews *bool   `json:"can_approve_pull_request_reviews,omitempty"`
}

type SelfHostRunnerPermissionsEnterprise struct {
	DisableSelfHostedRunnersForAllOrgs *bool `json:"disable_self_hosted_runners_for_all_orgs,omitempty"`
}

func (a SelfHostRunnerPermissionsEnterprise) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /enterprises/{enterprise}/actions/permissions
func (s *ActionsService) GetActionsPermissionsInEnterprise(ctx context.Context, enterprise string) (*ActionsPermissionsEnterprise, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/permissions
func (s *ActionsService) UpdateActionsPermissionsInEnterprise(ctx context.Context, enterprise string, body ActionsPermissionsEnterprise) (*ActionsPermissionsEnterprise, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/permissions/organizations
func (s *ActionsService) ListEnabledOrgsInEnterprise(ctx context.Context, owner string, opts *ListOptions) (*ActionsEnabledOnEnterpriseRepos, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/permissions/organizations
func (s *ActionsService) SetEnabledOrgsInEnterprise(ctx context.Context, owner string, organizationIDs []int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/permissions/organizations/{org_id}
func (s *ActionsService) AddEnabledOrgInEnterprise(ctx context.Context, owner string, organizationID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/actions/permissions/organizations/{org_id}
func (s *ActionsService) RemoveEnabledOrgInEnterprise(ctx context.Context, owner string, organizationID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/permissions/selected-actions
func (s *ActionsService) GetActionsAllowedInEnterprise(ctx context.Context, enterprise string) (*ActionsAllowed, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/permissions/selected-actions
func (s *ActionsService) UpdateActionsAllowedInEnterprise(ctx context.Context, enterprise string, body ActionsAllowed) (*ActionsAllowed, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/permissions/workflow
func (s *ActionsService) GetDefaultWorkflowPermissionsInEnterprise(ctx context.Context, enterprise string) (*DefaultWorkflowPermissionEnterprise, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/permissions/workflow
func (s *ActionsService) UpdateDefaultWorkflowPermissionsInEnterprise(ctx context.Context, enterprise string, body DefaultWorkflowPermissionEnterprise) (*DefaultWorkflowPermissionEnterprise, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/permissions/artifact-and-log-retention
func (s *ActionsService) GetArtifactAndLogRetentionPeriodInEnterprise(ctx context.Context, enterprise string) (*ArtifactPeriod, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/permissions/artifact-and-log-retention
func (s *ActionsService) UpdateArtifactAndLogRetentionPeriodInEnterprise(ctx context.Context, enterprise string, body ArtifactPeriodOpt) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/permissions/self-hosted-runners
func (s *ActionsService) GetSelfHostedRunnerPermissionsInEnterprise(ctx context.Context, enterprise string) (*SelfHostRunnerPermissionsEnterprise, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/permissions/self-hosted-runners
func (s *ActionsService) UpdateSelfHostedRunnerPermissionsInEnterprise(ctx context.Context, enterprise string, body SelfHostRunnerPermissionsEnterprise) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/permissions/fork-pr-workflows-private-repos
func (s *ActionsService) GetPrivateRepoForkPRWorkflowSettingsInEnterprise(ctx context.Context, enterprise string) (*WorkflowsPermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/permissions/fork-pr-workflows-private-repos
func (s *ActionsService) UpdatePrivateRepoForkPRWorkflowSettingsInEnterprise(ctx context.Context, enterprise string, body *WorkflowsPermissionsOpt) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/permissions/fork-pr-contributor-approval
func (s *ActionsService) GetEnterpriseForkPRContributorApprovalPermissions(ctx context.Context, enterprise string) (*ContributorApprovalPermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/actions/permissions/fork-pr-contributor-approval
func (s *ActionsService) UpdateEnterpriseForkPRContributorApprovalPermissions(ctx context.Context, enterprise string, body ContributorApprovalPermissions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

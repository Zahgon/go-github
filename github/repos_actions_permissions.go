package github

import (
	"context"
)

type ActionsPermissionsRepository struct {
	Enabled            *bool   `json:"enabled,omitempty"`
	AllowedActions     *string `json:"allowed_actions,omitempty"`
	SelectedActionsURL *string `json:"selected_actions_url,omitempty"`
	SHAPinningRequired *bool   `json:"sha_pinning_required,omitempty"`
}

func (a ActionsPermissionsRepository) String() string { _ = "STUB: not implemented"; return "" }

type DefaultWorkflowPermissionRepository struct {
	DefaultWorkflowPermissions   *string `json:"default_workflow_permissions,omitempty"`
	CanApprovePullRequestReviews *bool   `json:"can_approve_pull_request_reviews,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/actions/permissions
func (s *RepositoriesService) GetActionsPermissions(ctx context.Context, owner, repo string) (*ActionsPermissionsRepository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/permissions
func (s *RepositoriesService) UpdateActionsPermissions(ctx context.Context, owner, repo string, body ActionsPermissionsRepository) (*ActionsPermissionsRepository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/permissions/workflow
func (s *RepositoriesService) GetDefaultWorkflowPermissions(ctx context.Context, owner, repo string) (*DefaultWorkflowPermissionRepository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/permissions/workflow
func (s *RepositoriesService) UpdateDefaultWorkflowPermissions(ctx context.Context, owner, repo string, body DefaultWorkflowPermissionRepository) (*DefaultWorkflowPermissionRepository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/permissions/artifact-and-log-retention
func (s *RepositoriesService) GetArtifactAndLogRetentionPeriod(ctx context.Context, owner, repo string) (*ArtifactPeriod, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/permissions/artifact-and-log-retention
func (s *RepositoriesService) UpdateArtifactAndLogRetentionPeriod(ctx context.Context, owner, repo string, body ArtifactPeriodOpt) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/permissions/fork-pr-workflows-private-repos
func (s *RepositoriesService) GetPrivateRepoForkPRWorkflowSettings(ctx context.Context, owner, repo string) (*WorkflowsPermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/permissions/fork-pr-workflows-private-repos
func (s *RepositoriesService) UpdatePrivateRepoForkPRWorkflowSettings(ctx context.Context, owner, repo string, body *WorkflowsPermissionsOpt) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/permissions/fork-pr-contributor-approval
func (s *ActionsService) GetForkPRContributorApprovalPermissions(ctx context.Context, owner, repo string) (*ContributorApprovalPermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/permissions/fork-pr-contributor-approval
func (s *ActionsService) UpdateForkPRContributorApprovalPermissions(ctx context.Context, owner, repo string, body ContributorApprovalPermissions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

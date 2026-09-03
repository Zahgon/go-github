package github

import (
	"context"
)

type Environment struct {
	Owner                  *string         `json:"owner,omitempty"`
	Repo                   *string         `json:"repo,omitempty"`
	EnvironmentName        *string         `json:"environment_name,omitempty"`
	WaitTimer              *int            `json:"wait_timer,omitempty"`
	Reviewers              []*EnvReviewers `json:"reviewers,omitempty"`
	DeploymentBranchPolicy *BranchPolicy   `json:"deployment_branch_policy,omitempty"`

	ID              *int64            `json:"id,omitempty"`
	NodeID          *string           `json:"node_id,omitempty"`
	Name            *string           `json:"name,omitempty"`
	URL             *string           `json:"url,omitempty"`
	HTMLURL         *string           `json:"html_url,omitempty"`
	CreatedAt       *Timestamp        `json:"created_at,omitempty"`
	UpdatedAt       *Timestamp        `json:"updated_at,omitempty"`
	CanAdminsBypass *bool             `json:"can_admins_bypass,omitempty"`
	ProtectionRules []*ProtectionRule `json:"protection_rules,omitempty"`
}

type EnvReviewers struct {
	Type *string `json:"type,omitempty"`
	ID   *int64  `json:"id,omitempty"`
}

type BranchPolicy struct {
	ProtectedBranches    *bool `json:"protected_branches,omitempty"`
	CustomBranchPolicies *bool `json:"custom_branch_policies,omitempty"`
}

type EnvResponse struct {
	TotalCount   *int           `json:"total_count,omitempty"`
	Environments []*Environment `json:"environments,omitempty"`
}

type ProtectionRule struct {
	ID                *int64              `json:"id,omitempty"`
	NodeID            *string             `json:"node_id,omitempty"`
	PreventSelfReview *bool               `json:"prevent_self_review,omitempty"`
	Type              *string             `json:"type,omitempty"`
	WaitTimer         *int                `json:"wait_timer,omitempty"`
	Reviewers         []*RequiredReviewer `json:"reviewers,omitempty"`
}

type RequiredReviewer struct {
	Type     *string `json:"type,omitempty"`
	Reviewer any     `json:"reviewer,omitempty"`
}

type EnvironmentListOptions struct {
	ListOptions
}

func (r *RequiredReviewer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

//meta:operation GET /repos/{owner}/{repo}/environments
func (s *RepositoriesService) ListEnvironments(ctx context.Context, owner, repo string, opts *EnvironmentListOptions) (*EnvResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}
func (s *RepositoriesService) GetEnvironment(ctx context.Context, owner, repo, name string) (*Environment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c CreateUpdateEnvironment) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CreateUpdateEnvironment struct {
	WaitTimer              *int            `json:"wait_timer"`
	Reviewers              []*EnvReviewers `json:"reviewers"`
	CanAdminsBypass        *bool           `json:"can_admins_bypass"`
	DeploymentBranchPolicy *BranchPolicy   `json:"deployment_branch_policy"`
	PreventSelfReview      *bool           `json:"prevent_self_review,omitempty"`
}

type createUpdateEnvironmentNoEnterprise struct {
	DeploymentBranchPolicy *BranchPolicy `json:"deployment_branch_policy"`
}

//meta:operation PUT /repos/{owner}/{repo}/environments/{environment_name}
func (s *RepositoriesService) CreateUpdateEnvironment(ctx context.Context, owner, repo, name string, body *CreateUpdateEnvironment) (*Environment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *RepositoriesService) createNewEnvNoEnterprise(ctx context.Context, u string, environment *CreateUpdateEnvironment) (*Environment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/environments/{environment_name}
func (s *RepositoriesService) DeleteEnvironment(ctx context.Context, owner, repo, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

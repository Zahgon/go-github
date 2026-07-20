package github

import (
	"context"
)

type DeploymentBranchPolicy struct {
	Name   *string `json:"name,omitempty"`
	ID     *int64  `json:"id,omitempty"`
	NodeID *string `json:"node_id,omitempty"`
	Type   *string `json:"type,omitempty"`
}

type DeploymentBranchPolicyResponse struct {
	TotalCount     *int                      `json:"total_count,omitempty"`
	BranchPolicies []*DeploymentBranchPolicy `json:"branch_policies,omitempty"`
}

type CreateDeploymentBranchPolicyRequest struct {
	Name string  `json:"name"`
	Type *string `json:"type,omitempty"`
}

type UpdateDeploymentBranchPolicyRequest struct {
	Name string `json:"name"`
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies
func (s *RepositoriesService) ListDeploymentBranchPolicies(ctx context.Context, owner, repo, environment string, opts *ListOptions) (*DeploymentBranchPolicyResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies/{branch_policy_id}
func (s *RepositoriesService) GetDeploymentBranchPolicy(ctx context.Context, owner, repo, environment string, branchPolicyID int64) (*DeploymentBranchPolicy, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies
func (s *RepositoriesService) CreateDeploymentBranchPolicy(ctx context.Context, owner, repo, environment string, body CreateDeploymentBranchPolicyRequest) (*DeploymentBranchPolicy, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies/{branch_policy_id}
func (s *RepositoriesService) UpdateDeploymentBranchPolicy(ctx context.Context, owner, repo, environment string, branchPolicyID int64, body UpdateDeploymentBranchPolicyRequest) (*DeploymentBranchPolicy, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies/{branch_policy_id}
func (s *RepositoriesService) DeleteDeploymentBranchPolicy(ctx context.Context, owner, repo, environment string, branchPolicyID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

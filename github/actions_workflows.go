package github

import (
	"context"
)

type Workflow struct {
	ID        *int64     `json:"id,omitempty"`
	NodeID    *string    `json:"node_id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Path      *string    `json:"path,omitempty"`
	State     *string    `json:"state,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
	URL       *string    `json:"url,omitempty"`
	HTMLURL   *string    `json:"html_url,omitempty"`
	BadgeURL  *string    `json:"badge_url,omitempty"`
}

type Workflows struct {
	TotalCount *int        `json:"total_count,omitempty"`
	Workflows  []*Workflow `json:"workflows,omitempty"`
}

type WorkflowUsage struct {
	Billable *WorkflowBillMap `json:"billable,omitempty"`
}

type WorkflowBillMap map[string]*WorkflowBill

type WorkflowBill struct {
	TotalMS *int64 `json:"total_ms,omitempty"`
}

type CreateWorkflowDispatchEventRequest struct {
	Ref string `json:"ref"`

	Inputs map[string]any `json:"inputs,omitempty"`

	ReturnRunDetails *bool `json:"return_run_details,omitempty"`
}

type WorkflowDispatchRunDetails struct {
	WorkflowRunID *int64  `json:"workflow_run_id,omitempty"`
	RunURL        *string `json:"run_url,omitempty"`
	HTMLURL       *string `json:"html_url,omitempty"`
}

type WorkflowsPermissions struct {
	RunWorkflowsFromForkPullRequests  *bool `json:"run_workflows_from_fork_pull_requests,omitempty"`
	SendWriteTokensToWorkflows        *bool `json:"send_write_tokens_to_workflows,omitempty"`
	SendSecretsAndVariables           *bool `json:"send_secrets_and_variables,omitempty"`
	RequireApprovalForForkPRWorkflows *bool `json:"require_approval_for_fork_pr_workflows,omitempty"`
}

func (w WorkflowsPermissions) String() string { _ = "STUB: not implemented"; return "" }

type WorkflowsPermissionsOpt struct {
	RunWorkflowsFromForkPullRequests  bool  `json:"run_workflows_from_fork_pull_requests"`
	SendWriteTokensToWorkflows        *bool `json:"send_write_tokens_to_workflows,omitempty"`
	SendSecretsAndVariables           *bool `json:"send_secrets_and_variables,omitempty"`
	RequireApprovalForForkPRWorkflows *bool `json:"require_approval_for_fork_pr_workflows,omitempty"`
}

type ContributorApprovalPermissions struct {
	ApprovalPolicy string `json:"approval_policy"`
}

func (p ContributorApprovalPermissions) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/actions/workflows
func (s *ActionsService) ListWorkflows(ctx context.Context, owner, repo string, opts *ListOptions) (*Workflows, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/workflows/{workflow_id}
func (s *ActionsService) GetWorkflowByID(ctx context.Context, owner, repo string, workflowID int64) (*Workflow, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/workflows/{workflow_id}
func (s *ActionsService) GetWorkflowByFileName(ctx context.Context, owner, repo, workflowFileName string) (*Workflow, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) getWorkflow(ctx context.Context, url string) (*Workflow, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/workflows/{workflow_id}/timing
func (s *ActionsService) GetWorkflowUsageByID(ctx context.Context, owner, repo string, workflowID int64) (*WorkflowUsage, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/workflows/{workflow_id}/timing
func (s *ActionsService) GetWorkflowUsageByFileName(ctx context.Context, owner, repo, workflowFileName string) (*WorkflowUsage, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) getWorkflowUsage(ctx context.Context, url string) (*WorkflowUsage, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/actions/workflows/{workflow_id}/dispatches
func (s *ActionsService) CreateWorkflowDispatchEventByID(ctx context.Context, owner, repo string, workflowID int64, body CreateWorkflowDispatchEventRequest) (*WorkflowDispatchRunDetails, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/actions/workflows/{workflow_id}/dispatches
func (s *ActionsService) CreateWorkflowDispatchEventByFileName(ctx context.Context, owner, repo, workflowFileName string, body CreateWorkflowDispatchEventRequest) (*WorkflowDispatchRunDetails, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) createWorkflowDispatchEvent(ctx context.Context, url string, body *CreateWorkflowDispatchEventRequest) (*WorkflowDispatchRunDetails, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/workflows/{workflow_id}/enable
func (s *ActionsService) EnableWorkflowByID(ctx context.Context, owner, repo string, workflowID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/workflows/{workflow_id}/enable
func (s *ActionsService) EnableWorkflowByFileName(ctx context.Context, owner, repo, workflowFileName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/workflows/{workflow_id}/disable
func (s *ActionsService) DisableWorkflowByID(ctx context.Context, owner, repo string, workflowID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/workflows/{workflow_id}/disable
func (s *ActionsService) DisableWorkflowByFileName(ctx context.Context, owner, repo, workflowFileName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ActionsService) doNewPutRequest(ctx context.Context, url string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

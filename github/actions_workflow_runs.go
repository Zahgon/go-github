package github

import (
	"context"
	"net/url"
)

type WorkflowRun struct {
	ID                  *int64                `json:"id,omitempty"`
	Name                *string               `json:"name,omitempty"`
	NodeID              *string               `json:"node_id,omitempty"`
	HeadBranch          *string               `json:"head_branch,omitempty"`
	HeadSHA             *string               `json:"head_sha,omitempty"`
	Path                *string               `json:"path,omitempty"`
	RunNumber           *int                  `json:"run_number,omitempty"`
	RunAttempt          *int                  `json:"run_attempt,omitempty"`
	Event               *string               `json:"event,omitempty"`
	DisplayTitle        *string               `json:"display_title,omitempty"`
	Status              *string               `json:"status,omitempty"`
	Conclusion          *string               `json:"conclusion,omitempty"`
	WorkflowID          *int64                `json:"workflow_id,omitempty"`
	CheckSuiteID        *int64                `json:"check_suite_id,omitempty"`
	CheckSuiteNodeID    *string               `json:"check_suite_node_id,omitempty"`
	URL                 *string               `json:"url,omitempty"`
	HTMLURL             *string               `json:"html_url,omitempty"`
	PullRequests        []*PullRequest        `json:"pull_requests,omitempty"`
	CreatedAt           *Timestamp            `json:"created_at,omitempty"`
	UpdatedAt           *Timestamp            `json:"updated_at,omitempty"`
	RunStartedAt        *Timestamp            `json:"run_started_at,omitempty"`
	JobsURL             *string               `json:"jobs_url,omitempty"`
	LogsURL             *string               `json:"logs_url,omitempty"`
	CheckSuiteURL       *string               `json:"check_suite_url,omitempty"`
	ArtifactsURL        *string               `json:"artifacts_url,omitempty"`
	CancelURL           *string               `json:"cancel_url,omitempty"`
	RerunURL            *string               `json:"rerun_url,omitempty"`
	PreviousAttemptURL  *string               `json:"previous_attempt_url,omitempty"`
	HeadCommit          *HeadCommit           `json:"head_commit,omitempty"`
	WorkflowURL         *string               `json:"workflow_url,omitempty"`
	Repository          *Repository           `json:"repository,omitempty"`
	HeadRepository      *Repository           `json:"head_repository,omitempty"`
	Actor               *User                 `json:"actor,omitempty"`
	TriggeringActor     *User                 `json:"triggering_actor,omitempty"`
	ReferencedWorkflows []*ReferencedWorkflow `json:"referenced_workflows,omitempty"`
}

type WorkflowRuns struct {
	TotalCount   *int           `json:"total_count,omitempty"`
	WorkflowRuns []*WorkflowRun `json:"workflow_runs,omitempty"`
}

type ListWorkflowRunsOptions struct {
	Actor               string `url:"actor,omitempty"`
	Branch              string `url:"branch,omitempty"`
	Event               string `url:"event,omitempty"`
	Status              string `url:"status,omitempty"`
	Created             string `url:"created,omitempty"`
	HeadSHA             string `url:"head_sha,omitempty"`
	ExcludePullRequests bool   `url:"exclude_pull_requests,omitempty"`
	CheckSuiteID        int64  `url:"check_suite_id,omitempty"`
	ListOptions
}

type WorkflowRunUsage struct {
	Billable      *WorkflowRunBillMap `json:"billable,omitempty"`
	RunDurationMS *int64              `json:"run_duration_ms,omitempty"`
}

type WorkflowRunBillMap map[string]*WorkflowRunBill

type WorkflowRunBill struct {
	TotalMS *int64               `json:"total_ms,omitempty"`
	Jobs    *int                 `json:"jobs,omitempty"`
	JobRuns []*WorkflowRunJobRun `json:"job_runs,omitempty"`
}

type WorkflowRunJobRun struct {
	JobID      *int   `json:"job_id,omitempty"`
	DurationMS *int64 `json:"duration_ms,omitempty"`
}

type WorkflowRunAttemptOptions struct {
	ExcludePullRequests *bool `url:"exclude_pull_requests,omitempty"`
}

type PendingDeploymentsRequest struct {
	EnvironmentIDs []int64 `json:"environment_ids"`

	State   string `json:"state"`
	Comment string `json:"comment"`
}

type ReferencedWorkflow struct {
	Path *string `json:"path,omitempty"`
	SHA  *string `json:"sha,omitempty"`
	Ref  *string `json:"ref,omitempty"`
}

type PendingDeployment struct {
	Environment           *PendingDeploymentEnvironment `json:"environment,omitempty"`
	WaitTimer             *int64                        `json:"wait_timer,omitempty"`
	WaitTimerStartedAt    *Timestamp                    `json:"wait_timer_started_at,omitempty"`
	CurrentUserCanApprove *bool                         `json:"current_user_can_approve,omitempty"`
	Reviewers             []*RequiredReviewer           `json:"reviewers,omitempty"`
}

type PendingDeploymentEnvironment struct {
	ID      *int64  `json:"id,omitempty"`
	NodeID  *string `json:"node_id,omitempty"`
	Name    *string `json:"name,omitempty"`
	URL     *string `json:"url,omitempty"`
	HTMLURL *string `json:"html_url,omitempty"`
}

type ReviewCustomDeploymentProtectionRuleRequest struct {
	EnvironmentName string `json:"environment_name"`
	State           string `json:"state"`
	Comment         string `json:"comment"`
}

func (s *ActionsService) listWorkflowRuns(ctx context.Context, endpoint string, opts *ListWorkflowRunsOptions) (*WorkflowRuns, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/workflows/{workflow_id}/runs
func (s *ActionsService) ListWorkflowRunsByID(ctx context.Context, owner, repo string, workflowID int64, opts *ListWorkflowRunsOptions) (*WorkflowRuns, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/workflows/{workflow_id}/runs
func (s *ActionsService) ListWorkflowRunsByFileName(ctx context.Context, owner, repo, workflowFileName string, opts *ListWorkflowRunsOptions) (*WorkflowRuns, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/runs
func (s *ActionsService) ListRepositoryWorkflowRuns(ctx context.Context, owner, repo string, opts *ListWorkflowRunsOptions) (*WorkflowRuns, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/runs/{run_id}
func (s *ActionsService) GetWorkflowRunByID(ctx context.Context, owner, repo string, runID int64) (*WorkflowRun, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}
func (s *ActionsService) GetWorkflowRunAttempt(ctx context.Context, owner, repo string, runID int64, attemptNumber int, opts *WorkflowRunAttemptOptions) (*WorkflowRun, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/logs
func (s *ActionsService) GetWorkflowRunAttemptLogs(ctx context.Context, owner, repo string, runID int64, attemptNumber, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) getWorkflowRunAttemptLogsWithoutRateLimit(ctx context.Context, u string, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) getWorkflowRunAttemptLogsWithRateLimit(ctx context.Context, u string, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/actions/runs/{run_id}/rerun
func (s *ActionsService) RerunWorkflowByID(ctx context.Context, owner, repo string, runID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/actions/runs/{run_id}/rerun-failed-jobs
func (s *ActionsService) RerunFailedJobsByID(ctx context.Context, owner, repo string, runID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/actions/jobs/{job_id}/rerun
func (s *ActionsService) RerunJobByID(ctx context.Context, owner, repo string, jobID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/actions/runs/{run_id}/cancel
func (s *ActionsService) CancelWorkflowRunByID(ctx context.Context, owner, repo string, runID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/runs/{run_id}/logs
func (s *ActionsService) GetWorkflowRunLogs(ctx context.Context, owner, repo string, runID int64, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) getWorkflowRunLogsWithoutRateLimit(ctx context.Context, u string, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) getWorkflowRunLogsWithRateLimit(ctx context.Context, u string, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/actions/runs/{run_id}
func (s *ActionsService) DeleteWorkflowRun(ctx context.Context, owner, repo string, runID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/actions/runs/{run_id}/logs
func (s *ActionsService) DeleteWorkflowRunLogs(ctx context.Context, owner, repo string, runID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/runs/{run_id}/timing
func (s *ActionsService) GetWorkflowRunUsageByID(ctx context.Context, owner, repo string, runID int64) (*WorkflowRunUsage, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments
func (s *ActionsService) GetPendingDeployments(ctx context.Context, owner, repo string, runID int64) ([]*PendingDeployment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments
func (s *ActionsService) PendingDeployments(ctx context.Context, owner, repo string, runID int64, body PendingDeploymentsRequest) ([]*Deployment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/actions/runs/{run_id}/deployment_protection_rule
func (s *ActionsService) ReviewCustomDeploymentProtectionRule(ctx context.Context, owner, repo string, runID int64, body ReviewCustomDeploymentProtectionRuleRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

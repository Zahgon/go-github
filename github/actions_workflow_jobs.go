package github

import (
	"context"
	"net/url"
)

type TaskStep struct {
	Name        *string    `json:"name,omitempty"`
	Status      *string    `json:"status,omitempty"`
	Conclusion  *string    `json:"conclusion,omitempty"`
	Number      *int64     `json:"number,omitempty"`
	StartedAt   *Timestamp `json:"started_at,omitempty"`
	CompletedAt *Timestamp `json:"completed_at,omitempty"`
}

type WorkflowJob struct {
	ID          *int64      `json:"id,omitempty"`
	RunID       *int64      `json:"run_id,omitempty"`
	RunURL      *string     `json:"run_url,omitempty"`
	NodeID      *string     `json:"node_id,omitempty"`
	HeadBranch  *string     `json:"head_branch,omitempty"`
	HeadSHA     *string     `json:"head_sha,omitempty"`
	URL         *string     `json:"url,omitempty"`
	HTMLURL     *string     `json:"html_url,omitempty"`
	Status      *string     `json:"status,omitempty"`
	Conclusion  *string     `json:"conclusion,omitempty"`
	CreatedAt   *Timestamp  `json:"created_at,omitempty"`
	StartedAt   *Timestamp  `json:"started_at,omitempty"`
	CompletedAt *Timestamp  `json:"completed_at,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Steps       []*TaskStep `json:"steps,omitempty"`
	CheckRunURL *string     `json:"check_run_url,omitempty"`

	Labels          []string `json:"labels,omitempty"`
	RunnerID        *int64   `json:"runner_id,omitempty"`
	RunnerName      *string  `json:"runner_name,omitempty"`
	RunnerGroupID   *int64   `json:"runner_group_id,omitempty"`
	RunnerGroupName *string  `json:"runner_group_name,omitempty"`
	RunAttempt      *int64   `json:"run_attempt,omitempty"`
	WorkflowName    *string  `json:"workflow_name,omitempty"`
}

type Jobs struct {
	TotalCount *int           `json:"total_count,omitempty"`
	Jobs       []*WorkflowJob `json:"jobs,omitempty"`
}

type ListWorkflowJobsOptions struct {
	Filter string `url:"filter,omitempty"`
	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/actions/runs/{run_id}/jobs
func (s *ActionsService) ListWorkflowJobs(ctx context.Context, owner, repo string, runID int64, opts *ListWorkflowJobsOptions) (*Jobs, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/jobs
func (s *ActionsService) ListWorkflowJobsAttempt(ctx context.Context, owner, repo string, runID, attemptNumber int64, opts *ListOptions) (*Jobs, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/jobs/{job_id}
func (s *ActionsService) GetWorkflowJobByID(ctx context.Context, owner, repo string, jobID int64) (*WorkflowJob, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/jobs/{job_id}/logs
func (s *ActionsService) GetWorkflowJobLogs(ctx context.Context, owner, repo string, jobID int64, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) getWorkflowJobLogsWithoutRateLimit(ctx context.Context, u string, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) getWorkflowJobLogsWithRateLimit(ctx context.Context, u string, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

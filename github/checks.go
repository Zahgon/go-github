package github

import (
	"context"
)

type ChecksService service

type CheckRun struct {
	ID           *int64          `json:"id,omitempty"`
	NodeID       *string         `json:"node_id,omitempty"`
	HeadSHA      *string         `json:"head_sha,omitempty"`
	ExternalID   *string         `json:"external_id,omitempty"`
	URL          *string         `json:"url,omitempty"`
	HTMLURL      *string         `json:"html_url,omitempty"`
	DetailsURL   *string         `json:"details_url,omitempty"`
	Status       *string         `json:"status,omitempty"`
	Conclusion   *string         `json:"conclusion,omitempty"`
	StartedAt    *Timestamp      `json:"started_at,omitempty"`
	CompletedAt  *Timestamp      `json:"completed_at,omitempty"`
	Output       *CheckRunOutput `json:"output,omitempty"`
	Name         *string         `json:"name,omitempty"`
	CheckSuite   *CheckSuite     `json:"check_suite,omitempty"`
	App          *App            `json:"app,omitempty"`
	PullRequests []*PullRequest  `json:"pull_requests,omitempty"`
}

type CheckRunOutput struct {
	Title            *string               `json:"title,omitempty"`
	Summary          *string               `json:"summary,omitempty"`
	Text             *string               `json:"text,omitempty"`
	AnnotationsCount *int                  `json:"annotations_count,omitempty"`
	AnnotationsURL   *string               `json:"annotations_url,omitempty"`
	Annotations      []*CheckRunAnnotation `json:"annotations,omitempty"`
	Images           []*CheckRunImage      `json:"images,omitempty"`
}

type CheckRunAnnotation struct {
	Path            *string `json:"path,omitempty"`
	StartLine       *int    `json:"start_line,omitempty"`
	EndLine         *int    `json:"end_line,omitempty"`
	StartColumn     *int    `json:"start_column,omitempty"`
	EndColumn       *int    `json:"end_column,omitempty"`
	AnnotationLevel *string `json:"annotation_level,omitempty"`
	Message         *string `json:"message,omitempty"`
	Title           *string `json:"title,omitempty"`
	RawDetails      *string `json:"raw_details,omitempty"`
}

type CheckRunImage struct {
	Alt      *string `json:"alt,omitempty"`
	ImageURL *string `json:"image_url,omitempty"`
	Caption  *string `json:"caption,omitempty"`
}

type CheckSuite struct {
	ID           *int64         `json:"id,omitempty"`
	NodeID       *string        `json:"node_id,omitempty"`
	HeadBranch   *string        `json:"head_branch,omitempty"`
	HeadSHA      *string        `json:"head_sha,omitempty"`
	URL          *string        `json:"url,omitempty"`
	BeforeSHA    *string        `json:"before,omitempty"`
	AfterSHA     *string        `json:"after,omitempty"`
	Status       *string        `json:"status,omitempty"`
	Conclusion   *string        `json:"conclusion,omitempty"`
	CreatedAt    *Timestamp     `json:"created_at,omitempty"`
	UpdatedAt    *Timestamp     `json:"updated_at,omitempty"`
	App          *App           `json:"app,omitempty"`
	Repository   *Repository    `json:"repository,omitempty"`
	PullRequests []*PullRequest `json:"pull_requests,omitempty"`

	HeadCommit           *Commit `json:"head_commit,omitempty"`
	LatestCheckRunsCount *int64  `json:"latest_check_runs_count,omitempty"`
	Rerequestable        *bool   `json:"rerequestable,omitempty"`
	RunsRerequestable    *bool   `json:"runs_rerequestable,omitempty"`
}

func (c CheckRun) String() string { _ = "STUB: not implemented"; return "" }

func (c CheckSuite) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/check-runs/{check_run_id}
func (s *ChecksService) GetCheckRun(ctx context.Context, owner, repo string, checkRunID int64) (*CheckRun, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/check-suites/{check_suite_id}
func (s *ChecksService) GetCheckSuite(ctx context.Context, owner, repo string, checkSuiteID int64) (*CheckSuite, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CreateCheckRunOptions struct {
	Name        string            `json:"name"`
	HeadSHA     string            `json:"head_sha"`
	DetailsURL  *string           `json:"details_url,omitempty"`
	ExternalID  *string           `json:"external_id,omitempty"`
	Status      *string           `json:"status,omitempty"`
	Conclusion  *string           `json:"conclusion,omitempty"`
	StartedAt   *Timestamp        `json:"started_at,omitempty"`
	CompletedAt *Timestamp        `json:"completed_at,omitempty"`
	Output      *CheckRunOutput   `json:"output,omitempty"`
	Actions     []*CheckRunAction `json:"actions,omitempty"`
}

type CheckRunAction struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Identifier  string `json:"identifier"`
}

//meta:operation POST /repos/{owner}/{repo}/check-runs
func (s *ChecksService) CreateCheckRun(ctx context.Context, owner, repo string, body CreateCheckRunOptions) (*CheckRun, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type UpdateCheckRunOptions struct {
	Name        string            `json:"name"`
	DetailsURL  *string           `json:"details_url,omitempty"`
	ExternalID  *string           `json:"external_id,omitempty"`
	Status      *string           `json:"status,omitempty"`
	Conclusion  *string           `json:"conclusion,omitempty"`
	CompletedAt *Timestamp        `json:"completed_at,omitempty"`
	Output      *CheckRunOutput   `json:"output,omitempty"`
	Actions     []*CheckRunAction `json:"actions,omitempty"`
}

//meta:operation PATCH /repos/{owner}/{repo}/check-runs/{check_run_id}
func (s *ChecksService) UpdateCheckRun(ctx context.Context, owner, repo string, checkRunID int64, body UpdateCheckRunOptions) (*CheckRun, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/check-runs/{check_run_id}/annotations
func (s *ChecksService) ListCheckRunAnnotations(ctx context.Context, owner, repo string, checkRunID int64, opts *ListOptions) ([]*CheckRunAnnotation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ListCheckRunsOptions struct {
	CheckName *string `url:"check_name,omitempty"`
	Status    *string `url:"status,omitempty"`
	Filter    *string `url:"filter,omitempty"`
	AppID     *int64  `url:"app_id,omitempty"`

	ListOptions
}

type ListCheckRunsResults struct {
	Total     *int        `json:"total_count,omitempty"`
	CheckRuns []*CheckRun `json:"check_runs,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/commits/{ref}/check-runs
func (s *ChecksService) ListCheckRunsForRef(ctx context.Context, owner, repo, ref string, opts *ListCheckRunsOptions) (*ListCheckRunsResults, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/check-suites/{check_suite_id}/check-runs
func (s *ChecksService) ListCheckRunsCheckSuite(ctx context.Context, owner, repo string, checkSuiteID int64, opts *ListCheckRunsOptions) (*ListCheckRunsResults, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/check-runs/{check_run_id}/rerequest
func (s *ChecksService) ReRequestCheckRun(ctx context.Context, owner, repo string, checkRunID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ListCheckSuiteOptions struct {
	CheckName *string `url:"check_name,omitempty"`
	AppID     *int64  `url:"app_id,omitempty"`

	ListOptions
}

type ListCheckSuiteResults struct {
	Total       *int          `json:"total_count,omitempty"`
	CheckSuites []*CheckSuite `json:"check_suites,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/commits/{ref}/check-suites
func (s *ChecksService) ListCheckSuitesForRef(ctx context.Context, owner, repo, ref string, opts *ListCheckSuiteOptions) (*ListCheckSuiteResults, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type AutoTriggerCheck struct {
	AppID   *int64 `json:"app_id,omitempty"`
	Setting *bool  `json:"setting,omitempty"`
}

type CheckSuitePreferenceOptions struct {
	AutoTriggerChecks []*AutoTriggerCheck `json:"auto_trigger_checks,omitempty"`
}

type CheckSuitePreferenceResults struct {
	Preferences *PreferenceList `json:"preferences,omitempty"`
	Repository  *Repository     `json:"repository,omitempty"`
}

type PreferenceList struct {
	AutoTriggerChecks []*AutoTriggerCheck `json:"auto_trigger_checks,omitempty"`
}

//meta:operation PATCH /repos/{owner}/{repo}/check-suites/preferences
func (s *ChecksService) SetCheckSuitePreferences(ctx context.Context, owner, repo string, body CheckSuitePreferenceOptions) (*CheckSuitePreferenceResults, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CreateCheckSuiteOptions struct {
	HeadSHA    string  `json:"head_sha"`
	HeadBranch *string `json:"head_branch,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/check-suites
func (s *ChecksService) CreateCheckSuite(ctx context.Context, owner, repo string, body CreateCheckSuiteOptions) (*CheckSuite, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/check-suites/{check_suite_id}/rerequest
func (s *ChecksService) ReRequestCheckSuite(ctx context.Context, owner, repo string, checkSuiteID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

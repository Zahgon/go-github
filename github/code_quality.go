package github

import (
	"context"
)

type CodeQualityFindingRule struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Help        *string `json:"help,omitempty"`
	Severity    string  `json:"severity"`
	Category    string  `json:"category"`
}

type CodeQualityFindingLocation struct {
	Path        string `json:"path"`
	StartLine   *int   `json:"start_line,omitempty"`
	EndLine     *int   `json:"end_line,omitempty"`
	StartColumn *int   `json:"start_column,omitempty"`
	EndColumn   *int   `json:"end_column,omitempty"`
}

type CodeQualityFindingMessage struct {
	Text     string `json:"text"`
	Markdown string `json:"markdown"`
}

type CodeQualityFinding struct {
	Number    int                        `json:"number"`
	State     string                     `json:"state"`
	URL       string                     `json:"url"`
	Rule      CodeQualityFindingRule     `json:"rule"`
	Location  CodeQualityFindingLocation `json:"location"`
	Message   CodeQualityFindingMessage  `json:"message"`
	CreatedAt *Timestamp                 `json:"created_at,omitempty"`
}

type ListCodeQualityFindingsOptions struct {
	State     string `url:"state,omitempty"`
	Direction string `url:"direction,omitempty"`

	ListCursorOptions
}

type CodeQualityService service

type CodeQualitySetupConfiguration struct {
	State       *string    `json:"state,omitempty"`
	Languages   []string   `json:"languages,omitempty"`
	RunnerType  *string    `json:"runner_type,omitempty"`
	RunnerLabel *string    `json:"runner_label,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
	Schedule    *string    `json:"schedule,omitempty"`
}

type CodeQualityUpdateSetupRequest struct {
	State       *string  `json:"state,omitempty"`
	RunnerType  *string  `json:"runner_type,omitempty"`
	RunnerLabel *string  `json:"runner_label,omitempty"`
	Languages   []string `json:"languages,omitempty"`
}

type CodeQualityUpdateSetupResponse struct {
	RunID  *int64  `json:"run_id,omitempty"`
	RunURL *string `json:"run_url,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/code-quality/setup
func (s *CodeQualityService) GetSetup(ctx context.Context, owner, repo string) (*CodeQualitySetupConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/code-quality/setup
func (s *CodeQualityService) UpdateSetup(ctx context.Context, owner, repo string, body CodeQualityUpdateSetupRequest) (*CodeQualityUpdateSetupResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/code-quality/findings
func (s *CodeQualityService) ListFindings(ctx context.Context, owner, repo string, opts *ListCodeQualityFindingsOptions) ([]*CodeQualityFinding, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/code-quality/findings/{finding_number}
func (s *CodeQualityService) GetFinding(ctx context.Context, owner, repo string, findingNumber int) (*CodeQualityFinding, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

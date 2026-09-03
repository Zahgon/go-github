package github

import (
	"context"
)

type Label struct {
	ID          int64   `json:"id"`
	URL         string  `json:"url"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Description *string `json:"description"`
	Default     bool    `json:"default"`
	NodeID      string  `json:"node_id"`
}

func (l Label) String() string { _ = "STUB: not implemented"; return "" }

type CreateIssueLabelRequest struct {
	Name        string  `json:"name"`
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateIssueLabelRequest struct {
	NewName     *string `json:"new_name,omitempty"`
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/labels
func (s *IssuesService) ListLabels(ctx context.Context, owner, repo string, opts *ListOptions) ([]*Label, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/labels/{name}
func (s *IssuesService) GetLabel(ctx context.Context, owner, repo, name string) (*Label, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/labels
func (s *IssuesService) CreateLabel(ctx context.Context, owner, repo string, body CreateIssueLabelRequest) (*Label, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/labels/{name}
func (s *IssuesService) UpdateLabel(ctx context.Context, owner, repo, name string, body UpdateIssueLabelRequest) (*Label, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/labels/{name}
func (s *IssuesService) DeleteLabel(ctx context.Context, owner, repo, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}/labels
func (s *IssuesService) ListLabelsByIssue(ctx context.Context, owner, repo string, number int, opts *ListOptions) ([]*Label, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/issues/{issue_number}/labels
func (s *IssuesService) AddLabelsToIssue(ctx context.Context, owner, repo string, number int, body []string) ([]*Label, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/{issue_number}/labels/{name}
func (s *IssuesService) RemoveLabelForIssue(ctx context.Context, owner, repo string, number int, label string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/issues/{issue_number}/labels
func (s *IssuesService) ReplaceLabelsForIssue(ctx context.Context, owner, repo string, number int, body []string) ([]*Label, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/{issue_number}/labels
func (s *IssuesService) RemoveLabelsForIssue(ctx context.Context, owner, repo string, number int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/milestones/{milestone_number}/labels
func (s *IssuesService) ListLabelsForMilestone(ctx context.Context, owner, repo string, number int, opts *ListOptions) ([]*Label, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

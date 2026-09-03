package github

import (
	"context"
)

type SubIssueService service

type SubIssue Issue

func (i SubIssue) String() string { _ = "STUB: not implemented"; return "" }

type SubIssueListByIssueOptions struct {
	IssueListByRepoOptions
}

type SubIssueRequest struct {
	SubIssueID    int64  `json:"sub_issue_id"`
	AfterID       *int64 `json:"after_id,omitempty"`
	BeforeID      *int64 `json:"before_id,omitempty"`
	ReplaceParent *bool  `json:"replace_parent,omitempty"`
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/{issue_number}/sub_issue
func (s *SubIssueService) Remove(ctx context.Context, owner, repo string, issueNumber int64, subIssue SubIssueRequest) (*SubIssue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}/sub_issues
func (s *SubIssueService) ListByIssue(ctx context.Context, owner, repo string, issueNumber int64, opts *ListOptions) ([]*SubIssue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/issues/{issue_number}/sub_issues
func (s *SubIssueService) Add(ctx context.Context, owner, repo string, issueNumber int64, body SubIssueRequest) (*SubIssue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/issues/{issue_number}/sub_issues/priority
func (s *SubIssueService) Reprioritize(ctx context.Context, owner, repo string, issueNumber int64, body SubIssueRequest) (*SubIssue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}/parent
func (s *SubIssueService) GetParentIssue(ctx context.Context, owner, repo string, subIssueNumber int64) (*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

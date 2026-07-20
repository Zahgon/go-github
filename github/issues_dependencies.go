package github

import (
	"context"
)

type IssueDependencyRequest struct {
	IssueID int64 `json:"issue_id"`
}

//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}/dependencies/blocked_by
func (s *IssuesService) ListBlockedBy(ctx context.Context, owner, repo string, issueNumber int64, opts *ListOptions) ([]*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/issues/{issue_number}/dependencies/blocked_by
func (s *IssuesService) AddBlockedBy(ctx context.Context, owner, repo string, issueNumber int64, body IssueDependencyRequest) (*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/{issue_number}/dependencies/blocked_by/{issue_id}
func (s *IssuesService) RemoveBlockedBy(ctx context.Context, owner, repo string, issueNumber, blockingIssueID int64) (*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}/dependencies/blocking
func (s *IssuesService) ListBlocking(ctx context.Context, owner, repo string, issueNumber int64, opts *ListOptions) ([]*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

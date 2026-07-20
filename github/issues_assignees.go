package github

import (
	"context"
)

//meta:operation GET /repos/{owner}/{repo}/assignees
func (s *IssuesService) ListAssignees(ctx context.Context, owner, repo string, opts *ListOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/assignees/{assignee}
func (s *IssuesService) IsAssignee(ctx context.Context, owner, repo, user string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/issues/{issue_number}/assignees
func (s *IssuesService) AddAssignees(ctx context.Context, owner, repo string, number int, assignees []string) (*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/{issue_number}/assignees
func (s *IssuesService) RemoveAssignees(ctx context.Context, owner, repo string, number int, assignees []string) (*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

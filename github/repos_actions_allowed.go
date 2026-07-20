package github

import (
	"context"
)

//meta:operation GET /repos/{owner}/{repo}/actions/permissions/selected-actions
func (s *RepositoriesService) GetActionsAllowed(ctx context.Context, org, repo string) (*ActionsAllowed, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/permissions/selected-actions
func (s *RepositoriesService) EditActionsAllowed(ctx context.Context, org, repo string, body ActionsAllowed) (*ActionsAllowed, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

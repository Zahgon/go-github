package github

import (
	"context"
)

type RepositoryActionsAccessLevel struct {
	AccessLevel *string `json:"access_level,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/actions/permissions/access
func (s *RepositoriesService) GetActionsAccessLevel(ctx context.Context, owner, repo string) (*RepositoryActionsAccessLevel, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/permissions/access
func (s *RepositoriesService) EditActionsAccessLevel(ctx context.Context, owner, repo string, body RepositoryActionsAccessLevel) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

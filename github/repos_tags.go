package github

import (
	"context"
)

type TagProtection struct {
	ID      *int64  `json:"id"`
	Pattern *string `json:"pattern"`
}

type tagProtectionRequest struct {
	Pattern string `json:"pattern"`
}

//meta:operation GET /repos/{owner}/{repo}/tags/protection
func (s *RepositoriesService) ListTagProtection(ctx context.Context, owner, repo string) ([]*TagProtection, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/tags/protection
func (s *RepositoriesService) CreateTagProtection(ctx context.Context, owner, repo, pattern string) (*TagProtection, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/tags/protection/{tag_protection_id}
func (s *RepositoriesService) DeleteTagProtection(ctx context.Context, owner, repo string, tagProtectionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

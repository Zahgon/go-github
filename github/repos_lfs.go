package github

import (
	"context"
)

//meta:operation PUT /repos/{owner}/{repo}/lfs
func (s *RepositoriesService) EnableLFS(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/lfs
func (s *RepositoriesService) DisableLFS(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

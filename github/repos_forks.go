package github

import (
	"context"
)

type RepositoryListForksOptions struct {
	Sort string `url:"sort,omitempty"`

	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/forks
func (s *RepositoriesService) ListForks(ctx context.Context, owner, repo string, opts *RepositoryListForksOptions) ([]*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type RepositoryCreateForkOptions struct {
	Organization      string `json:"organization,omitempty"`
	Name              string `json:"name,omitempty"`
	DefaultBranchOnly bool   `json:"default_branch_only,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/forks
func (s *RepositoriesService) CreateFork(ctx context.Context, owner, repo string, body *RepositoryCreateForkOptions) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

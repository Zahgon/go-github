package github

import (
	"context"
)

type RepoImmutableReleasesStatus struct {
	Enabled         *bool `json:"enabled,omitempty"`
	EnforcedByOwner *bool `json:"enforced_by_owner,omitempty"`
}

//meta:operation PUT /repos/{owner}/{repo}/immutable-releases
func (s *RepositoriesService) EnableImmutableReleases(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/immutable-releases
func (s *RepositoriesService) DisableImmutableReleases(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/immutable-releases
func (s *RepositoriesService) AreImmutableReleasesEnabled(ctx context.Context, owner, repo string) (*RepoImmutableReleasesStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

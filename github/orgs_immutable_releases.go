package github

import (
	"context"
)

type ImmutableReleaseSettings struct {
	EnforcedRepositories *string `json:"enforced_repositories,omitempty"`

	SelectedRepositoriesURL *string `json:"selected_repositories_url,omitempty"`
}

type ImmutableReleasePolicy struct {
	EnforcedRepositories *string `json:"enforced_repositories,omitempty"`

	SelectedRepositoryIDs []int64 `json:"selected_repository_ids,omitempty"`
}

type setImmutableReleasesRepositoriesOptions struct {
	SelectedRepositoryIDs []int64 `json:"selected_repository_ids"`
}

//meta:operation GET /orgs/{org}/settings/immutable-releases
func (s *OrganizationsService) GetImmutableReleasesSettings(ctx context.Context, org string) (*ImmutableReleaseSettings, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/settings/immutable-releases
func (s *OrganizationsService) UpdateImmutableReleasesSettings(ctx context.Context, org string, body ImmutableReleasePolicy) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/settings/immutable-releases/repositories
func (s *OrganizationsService) ListImmutableReleaseRepositories(ctx context.Context, org string, opts *ListOptions) (*ListRepositories, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/settings/immutable-releases/repositories
func (s *OrganizationsService) SetImmutableReleaseRepositories(ctx context.Context, org string, repositoryIDs []int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/settings/immutable-releases/repositories/{repository_id}
func (s *OrganizationsService) EnableRepositoryForImmutableRelease(ctx context.Context, org string, repoID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/settings/immutable-releases/repositories/{repository_id}
func (s *OrganizationsService) DisableRepositoryForImmutableRelease(ctx context.Context, org string, repoID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

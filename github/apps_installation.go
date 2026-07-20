package github

import (
	"context"
)

type ListRepositories struct {
	TotalCount   *int          `json:"total_count,omitempty"`
	Repositories []*Repository `json:"repositories"`
}

//meta:operation GET /installation/repositories
func (s *AppsService) ListRepos(ctx context.Context, opts *ListOptions) (*ListRepositories, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/installations/{installation_id}/repositories
func (s *AppsService) ListUserRepos(ctx context.Context, id int64, opts *ListOptions) (*ListRepositories, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /user/installations/{installation_id}/repositories/{repository_id}
func (s *AppsService) AddRepository(ctx context.Context, instID, repoID int64) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /user/installations/{installation_id}/repositories/{repository_id}
func (s *AppsService) RemoveRepository(ctx context.Context, instID, repoID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /installation/token
func (s *AppsService) RevokeInstallationToken(ctx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

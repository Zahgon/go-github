package github

import (
	"context"
)

type MigrationService service

type Migration struct {
	ID   *int64  `json:"id,omitempty"`
	GUID *string `json:"guid,omitempty"`

	State *string `json:"state,omitempty"`

	LockRepositories *bool `json:"lock_repositories,omitempty"`

	ExcludeAttachments *bool         `json:"exclude_attachments,omitempty"`
	URL                *string       `json:"url,omitempty"`
	CreatedAt          *string       `json:"created_at,omitempty"`
	UpdatedAt          *string       `json:"updated_at,omitempty"`
	Repositories       []*Repository `json:"repositories,omitempty"`
}

func (m Migration) String() string { _ = "STUB: not implemented"; return "" }

type MigrationOptions struct {
	LockRepositories bool

	ExcludeAttachments bool

	ExcludeReleases bool

	Exclude []string
}

type startMigration struct {
	Repositories []string `json:"repositories,omitempty"`

	LockRepositories *bool `json:"lock_repositories,omitempty"`

	ExcludeAttachments *bool `json:"exclude_attachments,omitempty"`

	ExcludeReleases *bool `json:"exclude_releases,omitempty"`

	Exclude []string `json:"exclude,omitempty"`
}

//meta:operation POST /orgs/{org}/migrations
func (s *MigrationService) StartMigration(ctx context.Context, org string, repos []string, opts *MigrationOptions) (*Migration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/migrations
func (s *MigrationService) ListMigrations(ctx context.Context, org string, opts *ListOptions) ([]*Migration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/migrations/{migration_id}
func (s *MigrationService) MigrationStatus(ctx context.Context, org string, id int64) (*Migration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/migrations/{migration_id}/archive
func (s *MigrationService) MigrationArchiveURL(ctx context.Context, org string, id int64) (url string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

//meta:operation DELETE /orgs/{org}/migrations/{migration_id}/archive
func (s *MigrationService) DeleteMigration(ctx context.Context, org string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/migrations/{migration_id}/repos/{repo_name}/lock
func (s *MigrationService) UnlockRepo(ctx context.Context, org string, id int64, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

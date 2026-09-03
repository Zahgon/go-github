package github

import (
	"context"
)

type UserMigration struct {
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

func (m UserMigration) String() string { _ = "STUB: not implemented"; return "" }

type UserMigrationOptions struct {
	LockRepositories bool

	ExcludeAttachments bool
}

type startUserMigration struct {
	Repositories []string `json:"repositories,omitempty"`

	LockRepositories *bool `json:"lock_repositories,omitempty"`

	ExcludeAttachments *bool `json:"exclude_attachments,omitempty"`
}

//meta:operation POST /user/migrations
func (s *MigrationService) StartUserMigration(ctx context.Context, repos []string, opts *UserMigrationOptions) (*UserMigration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/migrations
func (s *MigrationService) ListUserMigrations(ctx context.Context, opts *ListOptions) ([]*UserMigration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/migrations/{migration_id}
func (s *MigrationService) UserMigrationStatus(ctx context.Context, id int64) (*UserMigration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/migrations/{migration_id}/archive
func (s *MigrationService) UserMigrationArchiveURL(ctx context.Context, id int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//meta:operation DELETE /user/migrations/{migration_id}/archive
func (s *MigrationService) DeleteUserMigration(ctx context.Context, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /user/migrations/{migration_id}/repos/{repo_name}/lock
func (s *MigrationService) UnlockUserRepo(ctx context.Context, id int64, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

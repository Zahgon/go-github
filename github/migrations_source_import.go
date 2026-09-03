package github

import (
	"context"
)

type Import struct {
	VCSURL *string `json:"vcs_url,omitempty"`

	VCS *string `json:"vcs,omitempty"`

	VCSUsername *string `json:"vcs_username,omitempty"`
	VCSPassword *string `json:"vcs_password,omitempty"`

	TFVCProject *string `json:"tfvc_project,omitempty"`

	UseLFS *string `json:"use_lfs,omitempty"`

	HasLargeFiles *bool `json:"has_large_files,omitempty"`

	LargeFilesSize *int `json:"large_files_size,omitempty"`

	LargeFilesCount *int `json:"large_files_count,omitempty"`

	Status        *string `json:"status,omitempty"`
	CommitCount   *int    `json:"commit_count,omitempty"`
	StatusText    *string `json:"status_text,omitempty"`
	AuthorsCount  *int    `json:"authors_count,omitempty"`
	Percent       *int    `json:"percent,omitempty"`
	PushPercent   *int    `json:"push_percent,omitempty"`
	URL           *string `json:"url,omitempty"`
	HTMLURL       *string `json:"html_url,omitempty"`
	AuthorsURL    *string `json:"authors_url,omitempty"`
	RepositoryURL *string `json:"repository_url,omitempty"`
	Message       *string `json:"message,omitempty"`
	FailedStep    *string `json:"failed_step,omitempty"`

	HumanName *string `json:"human_name,omitempty"`

	ProjectChoices []*Import `json:"project_choices,omitempty"`
}

func (i Import) String() string { _ = "STUB: not implemented"; return "" }

type SourceImportAuthor struct {
	ID         *int64  `json:"id,omitempty"`
	RemoteID   *string `json:"remote_id,omitempty"`
	RemoteName *string `json:"remote_name,omitempty"`
	Email      *string `json:"email,omitempty"`
	Name       *string `json:"name,omitempty"`
	URL        *string `json:"url,omitempty"`
	ImportURL  *string `json:"import_url,omitempty"`
}

func (a SourceImportAuthor) String() string { _ = "STUB: not implemented"; return "" }

type LargeFile struct {
	RefName *string `json:"ref_name,omitempty"`
	Path    *string `json:"path,omitempty"`
	OID     *string `json:"oid,omitempty"`
	Size    *int    `json:"size,omitempty"`
}

func (f LargeFile) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation PUT /repos/{owner}/{repo}/import
func (s *MigrationService) StartImport(ctx context.Context, owner, repo string, body *Import) (*Import, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/import
func (s *MigrationService) ImportProgress(ctx context.Context, owner, repo string) (*Import, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/import
func (s *MigrationService) UpdateImport(ctx context.Context, owner, repo string, body *Import) (*Import, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/import/authors
func (s *MigrationService) CommitAuthors(ctx context.Context, owner, repo string) ([]*SourceImportAuthor, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/import/authors/{author_id}
func (s *MigrationService) MapCommitAuthor(ctx context.Context, owner, repo string, id int64, body *SourceImportAuthor) (*SourceImportAuthor, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/import/lfs
func (s *MigrationService) SetLFSPreference(ctx context.Context, owner, repo string, body *Import) (*Import, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/import/large_files
func (s *MigrationService) LargeFiles(ctx context.Context, owner, repo string) ([]*LargeFile, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/import
func (s *MigrationService) CancelImport(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

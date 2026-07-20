package github

import (
	"context"
	"time"
)

type GistsService service

type Gist struct {
	ID          *string                   `json:"id,omitempty"`
	Description *string                   `json:"description,omitempty"`
	Public      *bool                     `json:"public,omitempty"`
	Owner       *User                     `json:"owner,omitempty"`
	Files       map[GistFilename]GistFile `json:"files,omitempty"`
	Comments    *int                      `json:"comments,omitempty"`
	HTMLURL     *string                   `json:"html_url,omitempty"`
	GitPullURL  *string                   `json:"git_pull_url,omitempty"`
	GitPushURL  *string                   `json:"git_push_url,omitempty"`
	CreatedAt   *Timestamp                `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp                `json:"updated_at,omitempty"`
	NodeID      *string                   `json:"node_id,omitempty"`
}

func (g Gist) String() string { _ = "STUB: not implemented"; return "" }

type CreateGistRequest struct {
	Description *string `json:"description,omitempty"`
	Public      *bool   `json:"public,omitempty"`

	Files map[GistFilename]*CreateGistFile `json:"files"`
}

type UpdateGistRequest struct {
	Description *string `json:"description,omitempty"`

	Files map[GistFilename]*UpdateGistFile `json:"files,omitempty"`
}

type GistFilename string

type GistFile struct {
	Size     *int    `json:"size,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Language *string `json:"language,omitempty"`
	Type     *string `json:"type,omitempty"`
	RawURL   *string `json:"raw_url,omitempty"`
	Content  *string `json:"content,omitempty"`
}

func (g GistFile) String() string { _ = "STUB: not implemented"; return "" }

type CreateGistFile struct {
	Content string `json:"content"`
}

type UpdateGistFile struct {
	Content *string `json:"content,omitempty"`

	Filename *string `json:"filename,omitempty"`
}

type GistCommit struct {
	URL          *string      `json:"url,omitempty"`
	Version      *string      `json:"version,omitempty"`
	User         *User        `json:"user,omitempty"`
	ChangeStatus *CommitStats `json:"change_status,omitempty"`
	CommittedAt  *Timestamp   `json:"committed_at,omitempty"`
	NodeID       *string      `json:"node_id,omitempty"`
}

func (gc GistCommit) String() string { _ = "STUB: not implemented"; return "" }

type GistFork struct {
	URL       *string    `json:"url,omitempty"`
	User      *User      `json:"user,omitempty"`
	ID        *string    `json:"id,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
	NodeID    *string    `json:"node_id,omitempty"`
}

func (gf GistFork) String() string { _ = "STUB: not implemented"; return "" }

type GistListOptions struct {
	Since time.Time `url:"since,omitempty"`

	ListOptions
}

//meta:operation GET /gists
//meta:operation GET /users/{username}/gists
func (s *GistsService) List(ctx context.Context, user string, opts *GistListOptions) ([]*Gist, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /gists/public
func (s *GistsService) ListAll(ctx context.Context, opts *GistListOptions) ([]*Gist, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /gists/starred
func (s *GistsService) ListStarred(ctx context.Context, opts *GistListOptions) ([]*Gist, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /gists/{gist_id}
func (s *GistsService) Get(ctx context.Context, id string) (*Gist, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /gists/{gist_id}/{sha}
func (s *GistsService) GetRevision(ctx context.Context, id, sha string) (*Gist, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /gists
func (s *GistsService) Create(ctx context.Context, body CreateGistRequest) (*Gist, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /gists/{gist_id}
func (s *GistsService) Update(ctx context.Context, id string, body UpdateGistRequest) (*Gist, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /gists/{gist_id}/commits
func (s *GistsService) ListCommits(ctx context.Context, id string, opts *ListOptions) ([]*GistCommit, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /gists/{gist_id}
func (s *GistsService) Delete(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /gists/{gist_id}/star
func (s *GistsService) Star(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /gists/{gist_id}/star
func (s *GistsService) Unstar(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /gists/{gist_id}/star
func (s *GistsService) IsStarred(ctx context.Context, id string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

//meta:operation POST /gists/{gist_id}/forks
func (s *GistsService) Fork(ctx context.Context, id string) (*Gist, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /gists/{gist_id}/forks
func (s *GistsService) ListForks(ctx context.Context, id string, opts *ListOptions) ([]*GistFork, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

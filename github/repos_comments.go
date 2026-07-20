package github

import (
	"context"
)

type RepositoryComment struct {
	HTMLURL   *string    `json:"html_url,omitempty"`
	URL       *string    `json:"url,omitempty"`
	ID        *int64     `json:"id,omitempty"`
	NodeID    *string    `json:"node_id,omitempty"`
	CommitID  *string    `json:"commit_id,omitempty"`
	User      *User      `json:"user,omitempty"`
	Reactions *Reactions `json:"reactions,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`

	Body *string `json:"body"`

	Path     *string `json:"path,omitempty"`
	Position *int    `json:"position,omitempty"`
}

func (r RepositoryComment) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/comments
func (s *RepositoriesService) ListComments(ctx context.Context, owner, repo string, opts *ListOptions) ([]*RepositoryComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/commits/{commit_sha}/comments
func (s *RepositoriesService) ListCommitComments(ctx context.Context, owner, repo, sha string, opts *ListOptions) ([]*RepositoryComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/commits/{commit_sha}/comments
func (s *RepositoriesService) CreateComment(ctx context.Context, owner, repo, sha string, body *RepositoryComment) (*RepositoryComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/comments/{comment_id}
func (s *RepositoriesService) GetComment(ctx context.Context, owner, repo string, id int64) (*RepositoryComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/comments/{comment_id}
func (s *RepositoriesService) UpdateComment(ctx context.Context, owner, repo string, id int64, body *RepositoryComment) (*RepositoryComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/comments/{comment_id}
func (s *RepositoriesService) DeleteComment(ctx context.Context, owner, repo string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

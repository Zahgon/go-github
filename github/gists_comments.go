package github

import (
	"context"
)

type GistComment struct {
	ID        *int64     `json:"id,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Body      *string    `json:"body,omitempty"`
	User      *User      `json:"user,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
}

func (g GistComment) String() string { _ = "STUB: not implemented"; return "" }

type CreateGistCommentRequest struct {
	Body string `json:"body"`
}

type UpdateGistCommentRequest struct {
	Body string `json:"body"`
}

//meta:operation GET /gists/{gist_id}/comments
func (s *GistsService) ListComments(ctx context.Context, gistID string, opts *ListOptions) ([]*GistComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /gists/{gist_id}/comments/{comment_id}
func (s *GistsService) GetComment(ctx context.Context, gistID string, commentID int64) (*GistComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /gists/{gist_id}/comments
func (s *GistsService) CreateComment(ctx context.Context, gistID string, body CreateGistCommentRequest) (*GistComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /gists/{gist_id}/comments/{comment_id}
func (s *GistsService) UpdateComment(ctx context.Context, gistID string, commentID int64, body UpdateGistCommentRequest) (*GistComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /gists/{gist_id}/comments/{comment_id}
func (s *GistsService) DeleteComment(ctx context.Context, gistID string, commentID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

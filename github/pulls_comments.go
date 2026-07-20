package github

import (
	"context"
	"time"
)

type PullRequestComment struct {
	ID                  *int64     `json:"id,omitempty"`
	NodeID              *string    `json:"node_id,omitempty"`
	InReplyTo           *int64     `json:"in_reply_to_id,omitempty"`
	Body                *string    `json:"body,omitempty"`
	Path                *string    `json:"path,omitempty"`
	DiffHunk            *string    `json:"diff_hunk,omitempty"`
	PullRequestReviewID *int64     `json:"pull_request_review_id,omitempty"`
	Position            *int       `json:"position,omitempty"`
	OriginalPosition    *int       `json:"original_position,omitempty"`
	StartLine           *int       `json:"start_line,omitempty"`
	Line                *int       `json:"line,omitempty"`
	OriginalLine        *int       `json:"original_line,omitempty"`
	OriginalStartLine   *int       `json:"original_start_line,omitempty"`
	Side                *string    `json:"side,omitempty"`
	StartSide           *string    `json:"start_side,omitempty"`
	CommitID            *string    `json:"commit_id,omitempty"`
	OriginalCommitID    *string    `json:"original_commit_id,omitempty"`
	User                *User      `json:"user,omitempty"`
	Reactions           *Reactions `json:"reactions,omitempty"`
	CreatedAt           *Timestamp `json:"created_at,omitempty"`
	UpdatedAt           *Timestamp `json:"updated_at,omitempty"`

	AuthorAssociation *string `json:"author_association,omitempty"`
	URL               *string `json:"url,omitempty"`
	HTMLURL           *string `json:"html_url,omitempty"`
	PullRequestURL    *string `json:"pull_request_url,omitempty"`

	SubjectType *string `json:"subject_type,omitempty"`
}

func (p PullRequestComment) String() string { _ = "STUB: not implemented"; return "" }

type PullRequestListCommentsOptions struct {
	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	Since time.Time `url:"since,omitempty"`

	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/pulls/comments
//meta:operation GET /repos/{owner}/{repo}/pulls/{pull_number}/comments
func (s *PullRequestsService) ListComments(ctx context.Context, owner, repo string, number int, opts *PullRequestListCommentsOptions) ([]*PullRequestComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pulls/comments/{comment_id}
func (s *PullRequestsService) GetComment(ctx context.Context, owner, repo string, commentID int64) (*PullRequestComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/pulls/{pull_number}/comments
func (s *PullRequestsService) CreateComment(ctx context.Context, owner, repo string, number int, body *PullRequestComment) (*PullRequestComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/pulls/{pull_number}/comments
func (s *PullRequestsService) CreateCommentInReplyTo(ctx context.Context, owner, repo string, number int, body string, commentID int64) (*PullRequestComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/pulls/comments/{comment_id}
func (s *PullRequestsService) EditComment(ctx context.Context, owner, repo string, commentID int64, body *PullRequestComment) (*PullRequestComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/pulls/comments/{comment_id}
func (s *PullRequestsService) DeleteComment(ctx context.Context, owner, repo string, commentID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

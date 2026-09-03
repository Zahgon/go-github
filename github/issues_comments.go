package github

import (
	"context"
	"time"
)

type IssueComment struct {
	ID        *int64     `json:"id,omitempty"`
	NodeID    *string    `json:"node_id,omitempty"`
	Body      *string    `json:"body,omitempty"`
	User      *User      `json:"user,omitempty"`
	Reactions *Reactions `json:"reactions,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`

	AuthorAssociation     *string                `json:"author_association,omitempty"`
	PerformedViaGithubApp *App                   `json:"performed_via_github_app,omitempty"`
	Pin                   *PinnedIssueComment    `json:"pin,omitempty"`
	Minimized             *MinimizedIssueComment `json:"minimized,omitempty"`
	URL                   *string                `json:"url,omitempty"`
	HTMLURL               *string                `json:"html_url,omitempty"`
	IssueURL              *string                `json:"issue_url,omitempty"`
}

type PinnedIssueComment struct {
	PinnedAt *Timestamp `json:"pinned_at,omitempty"`
	PinnedBy *User      `json:"pinned_by,omitempty"`
}

type MinimizedIssueComment struct {
	Reason *string `json:"reason,omitempty"`
}

func (i IssueComment) String() string { _ = "STUB: not implemented"; return "" }

type IssueCommentRequest struct {
	Body string `json:"body"`
}

type IssueListCommentsOptions struct {
	Sort *string `url:"sort,omitempty"`

	Direction *string `url:"direction,omitempty"`

	Since *time.Time `url:"since,omitempty"`

	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/issues/comments
//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}/comments
func (s *IssuesService) ListComments(ctx context.Context, owner, repo string, number int, opts *IssueListCommentsOptions) ([]*IssueComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/comments/{comment_id}
func (s *IssuesService) GetComment(ctx context.Context, owner, repo string, commentID int64) (*IssueComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/issues/{issue_number}/comments
func (s *IssuesService) CreateComment(ctx context.Context, owner, repo string, number int, body IssueCommentRequest) (*IssueComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/issues/comments/{comment_id}
func (s *IssuesService) UpdateComment(ctx context.Context, owner, repo string, commentID int64, body IssueCommentRequest) (*IssueComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/comments/{comment_id}
func (s *IssuesService) DeleteComment(ctx context.Context, owner, repo string, commentID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

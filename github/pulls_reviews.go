package github

import (
	"context"
	"errors"
)

var ErrMixedCommentStyles = errors.New("cannot use both position and side/line form comments")

type PullRequestReview struct {
	ID             *int64     `json:"id,omitempty"`
	NodeID         *string    `json:"node_id,omitempty"`
	User           *User      `json:"user,omitempty"`
	Body           *string    `json:"body,omitempty"`
	SubmittedAt    *Timestamp `json:"submitted_at,omitempty"`
	CommitID       *string    `json:"commit_id,omitempty"`
	HTMLURL        *string    `json:"html_url,omitempty"`
	PullRequestURL *string    `json:"pull_request_url,omitempty"`
	State          *string    `json:"state,omitempty"`

	AuthorAssociation *string `json:"author_association,omitempty"`
}

func (p PullRequestReview) String() string { _ = "STUB: not implemented"; return "" }

type DraftReviewComment struct {
	Path     *string `json:"path,omitempty"`
	Position *int    `json:"position,omitempty"`
	Body     *string `json:"body,omitempty"`

	StartSide *string `json:"start_side,omitempty"`
	Side      *string `json:"side,omitempty"`
	StartLine *int    `json:"start_line,omitempty"`
	Line      *int    `json:"line,omitempty"`
}

func (c DraftReviewComment) String() string { _ = "STUB: not implemented"; return "" }

type PullRequestReviewRequest struct {
	NodeID   *string               `json:"node_id,omitempty"`
	CommitID *string               `json:"commit_id,omitempty"`
	Body     *string               `json:"body,omitempty"`
	Event    *string               `json:"event,omitempty"`
	Comments []*DraftReviewComment `json:"comments,omitempty"`
}

func (r PullRequestReviewRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *PullRequestReviewRequest) isComfortFadePreview() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type PullRequestReviewDismissalRequest struct {
	Message *string `json:"message,omitempty"`
}

func (r PullRequestReviewDismissalRequest) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/pulls/{pull_number}/reviews
func (s *PullRequestsService) ListReviews(ctx context.Context, owner, repo string, number int, opts *ListOptions) ([]*PullRequestReview, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}
func (s *PullRequestsService) GetReview(ctx context.Context, owner, repo string, number int, reviewID int64) (*PullRequestReview, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}
func (s *PullRequestsService) DeletePendingReview(ctx context.Context, owner, repo string, number int, reviewID int64) (*PullRequestReview, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/comments
func (s *PullRequestsService) ListReviewComments(ctx context.Context, owner, repo string, number int, reviewID int64, opts *ListOptions) ([]*PullRequestComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/pulls/{pull_number}/reviews
func (s *PullRequestsService) CreateReview(ctx context.Context, owner, repo string, number int, body *PullRequestReviewRequest) (*PullRequestReview, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}
func (s *PullRequestsService) UpdateReview(ctx context.Context, owner, repo string, number int, reviewID int64, body string) (*PullRequestReview, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/events
func (s *PullRequestsService) SubmitReview(ctx context.Context, owner, repo string, number int, reviewID int64, body *PullRequestReviewRequest) (*PullRequestReview, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/dismissals
func (s *PullRequestsService) DismissReview(ctx context.Context, owner, repo string, number int, reviewID int64, body *PullRequestReviewDismissalRequest) (*PullRequestReview, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

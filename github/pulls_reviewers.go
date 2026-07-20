package github

import (
	"context"
)

type ReviewersRequest struct {
	NodeID        *string  `json:"node_id,omitempty"`
	Reviewers     []string `json:"reviewers,omitempty"`
	TeamReviewers []string `json:"team_reviewers,omitempty"`
}

type Reviewers struct {
	Users []*User `json:"users,omitempty"`
	Teams []*Team `json:"teams,omitempty"`
}

type removeReviewersRequest struct {
	NodeID *string `json:"node_id,omitempty"`

	Reviewers     []string `json:"reviewers"`
	TeamReviewers []string `json:"team_reviewers,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/pulls/{pull_number}/requested_reviewers
func (s *PullRequestsService) RequestReviewers(ctx context.Context, owner, repo string, number int, reviewers ReviewersRequest) (*PullRequest, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pulls/{pull_number}/requested_reviewers
func (s *PullRequestsService) ListReviewers(ctx context.Context, owner, repo string, number int) (*Reviewers, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/pulls/{pull_number}/requested_reviewers
func (s *PullRequestsService) RemoveReviewers(ctx context.Context, owner, repo string, number int, reviewers ReviewersRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package github

import (
	"context"
)

type IssueEvent struct {
	ID  *int64  `json:"id,omitempty"`
	URL *string `json:"url,omitempty"`

	Actor *User `json:"actor,omitempty"`

	Action string `json:"action,omitempty"`

	Event *string `json:"event,omitempty"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`
	Issue     *Issue     `json:"issue,omitempty"`

	Repository            *Repository      `json:"repository,omitempty"`
	Assignee              *User            `json:"assignee,omitempty"`
	Assigner              *User            `json:"assigner,omitempty"`
	CommitID              *string          `json:"commit_id,omitempty"`
	Milestone             *Milestone       `json:"milestone,omitempty"`
	Label                 *Label           `json:"label,omitempty"`
	Rename                *Rename          `json:"rename,omitempty"`
	LockReason            *string          `json:"lock_reason,omitempty"`
	DismissedReview       *DismissedReview `json:"dismissed_review,omitempty"`
	RequestedReviewer     *User            `json:"requested_reviewer,omitempty"`
	RequestedTeam         *Team            `json:"requested_team,omitempty"`
	ReviewRequester       *User            `json:"review_requester,omitempty"`
	PerformedViaGithubApp *App             `json:"performed_via_github_app,omitempty"`
}

type DismissedReview struct {
	State             *string `json:"state,omitempty"`
	ReviewID          *int64  `json:"review_id,omitempty"`
	DismissalMessage  *string `json:"dismissal_message,omitempty"`
	DismissalCommitID *string `json:"dismissal_commit_id,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}/events
func (s *IssuesService) ListIssueEvents(ctx context.Context, owner, repo string, number int, opts *ListOptions) ([]*IssueEvent, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/events
func (s *IssuesService) ListRepositoryEvents(ctx context.Context, owner, repo string, opts *ListOptions) ([]*IssueEvent, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/events/{event_id}
func (s *IssuesService) GetEvent(ctx context.Context, owner, repo string, id int64) (*IssueEvent, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type Rename struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

func (r Rename) String() string { _ = "STUB: not implemented"; return "" }

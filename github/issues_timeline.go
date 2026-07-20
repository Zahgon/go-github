package github

import (
	"context"
)

type Timeline struct {
	ID        *int64  `json:"id,omitempty"`
	URL       *string `json:"url,omitempty"`
	CommitURL *string `json:"commit_url,omitempty"`

	Actor *User `json:"actor,omitempty"`

	User *User `json:"user,omitempty"`

	Author *CommitAuthor `json:"author,omitempty"`

	Committer *CommitAuthor `json:"committer,omitempty"`

	SHA *string `json:"sha,omitempty"`

	Message *string `json:"message,omitempty"`

	Parents []*Commit `json:"parents,omitempty"`

	Event *string `json:"event,omitempty"`

	CommitID *string `json:"commit_id,omitempty"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`

	Label *Label `json:"label,omitempty"`

	Assignee *User `json:"assignee,omitempty"`
	Assigner *User `json:"assigner,omitempty"`

	Milestone *Milestone `json:"milestone,omitempty"`

	Source *Source `json:"source,omitempty"`

	Rename *Rename `json:"rename,omitempty"`

	State *string `json:"state,omitempty"`

	Reviewer *User `json:"requested_reviewer,omitempty"`

	RequestedTeam *Team `json:"requested_team,omitempty"`

	Requester *User `json:"review_requester,omitempty"`

	Body        *string    `json:"body,omitempty"`
	SubmittedAt *Timestamp `json:"submitted_at,omitempty"`

	PerformedViaGithubApp *App `json:"performed_via_github_app,omitempty"`
}

type Source struct {
	ID    *int64  `json:"id,omitempty"`
	URL   *string `json:"url,omitempty"`
	Actor *User   `json:"actor,omitempty"`
	Type  *string `json:"type,omitempty"`
	Issue *Issue  `json:"issue,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}/timeline
func (s *IssuesService) ListIssueTimeline(ctx context.Context, owner, repo string, number int, opts *ListOptions) ([]*Timeline, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

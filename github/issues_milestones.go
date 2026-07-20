package github

import (
	"context"
)

type Milestone struct {
	URL          *string    `json:"url,omitempty"`
	HTMLURL      *string    `json:"html_url,omitempty"`
	LabelsURL    *string    `json:"labels_url,omitempty"`
	ID           *int64     `json:"id,omitempty"`
	Number       *int       `json:"number,omitempty"`
	State        *string    `json:"state,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Creator      *User      `json:"creator,omitempty"`
	OpenIssues   *int       `json:"open_issues,omitempty"`
	ClosedIssues *int       `json:"closed_issues,omitempty"`
	CreatedAt    *Timestamp `json:"created_at,omitempty"`
	UpdatedAt    *Timestamp `json:"updated_at,omitempty"`
	ClosedAt     *Timestamp `json:"closed_at,omitempty"`
	DueOn        *Timestamp `json:"due_on,omitempty"`
	NodeID       *string    `json:"node_id,omitempty"`
}

func (m Milestone) String() string { _ = "STUB: not implemented"; return "" }

type MilestoneListOptions struct {
	State string `url:"state,omitempty"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/milestones
func (s *IssuesService) ListMilestones(ctx context.Context, owner, repo string, opts *MilestoneListOptions) ([]*Milestone, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/milestones/{milestone_number}
func (s *IssuesService) GetMilestone(ctx context.Context, owner, repo string, number int) (*Milestone, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/milestones
func (s *IssuesService) CreateMilestone(ctx context.Context, owner, repo string, body *Milestone) (*Milestone, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/milestones/{milestone_number}
func (s *IssuesService) EditMilestone(ctx context.Context, owner, repo string, number int, body *Milestone) (*Milestone, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/milestones/{milestone_number}
func (s *IssuesService) DeleteMilestone(ctx context.Context, owner, repo string, number int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

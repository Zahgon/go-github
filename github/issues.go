package github

import (
	"context"
	"time"
)

type IssuesService service

type IssueDependenciesSummary struct {
	BlockedBy      *int `json:"blocked_by,omitempty"`
	Blocking       *int `json:"blocking,omitempty"`
	TotalBlockedBy *int `json:"total_blocked_by,omitempty"`
	TotalBlocking  *int `json:"total_blocking,omitempty"`
}

type SubIssuesSummary struct {
	Total            *int `json:"total,omitempty"`
	Completed        *int `json:"completed,omitempty"`
	PercentCompleted *int `json:"percent_completed,omitempty"`
}

type Issue struct {
	ID     *int64  `json:"id,omitempty"`
	Number *int    `json:"number,omitempty"`
	State  *string `json:"state,omitempty"`

	StateReason *string `json:"state_reason,omitempty"`
	Locked      *bool   `json:"locked,omitempty"`
	Title       *string `json:"title,omitempty"`
	Body        *string `json:"body,omitempty"`

	AuthorAssociation        *string                   `json:"author_association,omitempty"`
	User                     *User                     `json:"user,omitempty"`
	Labels                   []*Label                  `json:"labels,omitempty"`
	Assignee                 *User                     `json:"assignee,omitempty"`
	Comments                 *int                      `json:"comments,omitempty"`
	ClosedAt                 *Timestamp                `json:"closed_at,omitempty"`
	CreatedAt                *Timestamp                `json:"created_at,omitempty"`
	UpdatedAt                *Timestamp                `json:"updated_at,omitempty"`
	ClosedBy                 *User                     `json:"closed_by,omitempty"`
	URL                      *string                   `json:"url,omitempty"`
	HTMLURL                  *string                   `json:"html_url,omitempty"`
	CommentsURL              *string                   `json:"comments_url,omitempty"`
	EventsURL                *string                   `json:"events_url,omitempty"`
	LabelsURL                *string                   `json:"labels_url,omitempty"`
	RepositoryURL            *string                   `json:"repository_url,omitempty"`
	ParentIssueURL           *string                   `json:"parent_issue_url,omitempty"`
	Milestone                *Milestone                `json:"milestone,omitempty"`
	PullRequestLinks         *PullRequestLinks         `json:"pull_request,omitempty"`
	Repository               *Repository               `json:"repository,omitempty"`
	Reactions                *Reactions                `json:"reactions,omitempty"`
	Assignees                []*User                   `json:"assignees,omitempty"`
	NodeID                   *string                   `json:"node_id,omitempty"`
	Draft                    *bool                     `json:"draft,omitempty"`
	Type                     *IssueType                `json:"type,omitempty"`
	PinnedComment            *IssueComment             `json:"pinned_comment,omitempty"`
	PerformedViaGithubApp    *App                      `json:"performed_via_github_app,omitempty"`
	IssueDependenciesSummary *IssueDependenciesSummary `json:"issue_dependencies_summary,omitempty"`
	SubIssuesSummary         *SubIssuesSummary         `json:"sub_issues_summary,omitempty"`
	IssueFieldValues         []*IssueFieldValue        `json:"issue_field_values,omitempty"`

	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	ActiveLockReason *string `json:"active_lock_reason,omitempty"`
}

func (i Issue) String() string { _ = "STUB: not implemented"; return "" }

func (i Issue) IsPullRequest() bool { _ = "STUB: not implemented"; return false }

type IssueRequest struct {
	Title    *string   `json:"title,omitempty"`
	Body     *string   `json:"body,omitempty"`
	Labels   *[]string `json:"labels,omitempty"`
	Assignee *string   `json:"assignee,omitempty"`
	State    *string   `json:"state,omitempty"`

	StateReason      *string                   `json:"state_reason,omitempty"`
	Milestone        *int                      `json:"milestone,omitempty"`
	Assignees        *[]string                 `json:"assignees,omitempty"`
	Type             *string                   `json:"type,omitempty"`
	IssueFieldValues []*IssueRequestFieldValue `json:"issue_field_values,omitempty"`
}

type PullRequestLinks struct {
	URL      *string    `json:"url,omitempty"`
	HTMLURL  *string    `json:"html_url,omitempty"`
	DiffURL  *string    `json:"diff_url,omitempty"`
	PatchURL *string    `json:"patch_url,omitempty"`
	MergedAt *Timestamp `json:"merged_at,omitempty"`
}

type IssueType struct {
	ID          *int64     `json:"id,omitempty"`
	NodeID      *string    `json:"node_id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Color       *string    `json:"color,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
}

type IssueFieldValueSingleSelectOption struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type IssueRequestFieldValue struct {
	FieldID int64 `json:"field_id"`
	Value   any   `json:"value"`
}

type IssueFieldValue struct {
	IssueFieldID       int64                              `json:"issue_field_id"`
	NodeID             string                             `json:"node_id"`
	DataType           string                             `json:"data_type"`
	Value              any                                `json:"value"`
	SingleSelectOption *IssueFieldValueSingleSelectOption `json:"single_select_option,omitempty"`
}

type ListAllIssuesOptions struct {
	Filter string `url:"filter,omitempty"`

	State string `url:"state,omitempty"`

	Labels []string `url:"labels,comma,omitempty"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	Since time.Time `url:"since,omitempty"`

	Collab bool `url:"collab,omitempty"`
	Orgs   bool `url:"orgs,omitempty"`
	Owned  bool `url:"owned,omitempty"`
	Pulls  bool `url:"pulls,omitempty"`

	ListOptions
}

//meta:operation GET /issues
func (s *IssuesService) ListAllIssues(ctx context.Context, opts *ListAllIssuesOptions) ([]*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ListUserIssuesOptions struct {
	Filter string `url:"filter,omitempty"`

	State string `url:"state,omitempty"`

	Labels []string `url:"labels,comma,omitempty"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	Since time.Time `url:"since,omitempty"`

	ListOptions
}

//meta:operation GET /user/issues
func (s *IssuesService) ListUserIssues(ctx context.Context, opts *ListUserIssuesOptions) ([]*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type IssueListByOrgOptions struct {
	Filter string `url:"filter,omitempty"`

	State string `url:"state,omitempty"`

	Labels []string `url:"labels,comma,omitempty"`

	Type string `url:"type,omitempty"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	Since time.Time `url:"since,omitempty"`

	ListOptions
}

//meta:operation GET /orgs/{org}/issues
func (s *IssuesService) ListByOrg(ctx context.Context, org string, opts *IssueListByOrgOptions) ([]*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type IssueListByRepoOptions struct {
	Milestone string `url:"milestone,omitempty"`

	State string `url:"state,omitempty"`

	Assignee string `url:"assignee,omitempty"`

	Type string `url:"type,omitempty"`

	Creator string `url:"creator,omitempty"`

	Mentioned string `url:"mentioned,omitempty"`

	Labels []string `url:"labels,omitempty,comma"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	Since time.Time `url:"since,omitempty"`

	ListCursorOptions

	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/issues
func (s *IssuesService) ListByRepo(ctx context.Context, owner, repo string, opts *IssueListByRepoOptions) ([]*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}
func (s *IssuesService) Get(ctx context.Context, owner, repo string, number int) (*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/issues
func (s *IssuesService) Create(ctx context.Context, owner, repo string, body *IssueRequest) (*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/issues/{issue_number}
func (s *IssuesService) Edit(ctx context.Context, owner, repo string, number int, body *IssueRequest) (*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/issues/{issue_number}
func (s *IssuesService) RemoveMilestone(ctx context.Context, owner, repo string, issueNumber int) (*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type LockIssueOptions struct {
	LockReason string `json:"lock_reason,omitempty"`
}

//meta:operation PUT /repos/{owner}/{repo}/issues/{issue_number}/lock
func (s *IssuesService) Lock(ctx context.Context, owner, repo string, number int, body *LockIssueOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/{issue_number}/lock
func (s *IssuesService) Unlock(ctx context.Context, owner, repo string, number int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

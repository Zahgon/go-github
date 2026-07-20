package github

import (
	"context"
)

type PullRequestsService service

type PullRequestAutoMerge struct {
	EnabledBy     *User   `json:"enabled_by,omitempty"`
	MergeMethod   *string `json:"merge_method,omitempty"`
	CommitTitle   *string `json:"commit_title,omitempty"`
	CommitMessage *string `json:"commit_message,omitempty"`
}

type PullRequest struct {
	ID                *int64     `json:"id,omitempty"`
	Number            *int       `json:"number,omitempty"`
	State             *string    `json:"state,omitempty"`
	Locked            *bool      `json:"locked,omitempty"`
	Title             *string    `json:"title,omitempty"`
	Body              *string    `json:"body,omitempty"`
	CreatedAt         *Timestamp `json:"created_at,omitempty"`
	UpdatedAt         *Timestamp `json:"updated_at,omitempty"`
	ClosedAt          *Timestamp `json:"closed_at,omitempty"`
	MergedAt          *Timestamp `json:"merged_at,omitempty"`
	Labels            []*Label   `json:"labels,omitempty"`
	User              *User      `json:"user,omitempty"`
	Draft             *bool      `json:"draft,omitempty"`
	URL               *string    `json:"url,omitempty"`
	HTMLURL           *string    `json:"html_url,omitempty"`
	IssueURL          *string    `json:"issue_url,omitempty"`
	StatusesURL       *string    `json:"statuses_url,omitempty"`
	DiffURL           *string    `json:"diff_url,omitempty"`
	PatchURL          *string    `json:"patch_url,omitempty"`
	CommitsURL        *string    `json:"commits_url,omitempty"`
	CommentsURL       *string    `json:"comments_url,omitempty"`
	ReviewCommentsURL *string    `json:"review_comments_url,omitempty"`
	ReviewCommentURL  *string    `json:"review_comment_url,omitempty"`
	Assignee          *User      `json:"assignee,omitempty"`
	Assignees         []*User    `json:"assignees,omitempty"`
	Milestone         *Milestone `json:"milestone,omitempty"`

	AuthorAssociation  *string               `json:"author_association,omitempty"`
	NodeID             *string               `json:"node_id,omitempty"`
	RequestedReviewers []*User               `json:"requested_reviewers,omitempty"`
	AutoMerge          *PullRequestAutoMerge `json:"auto_merge,omitempty"`

	Merged              *bool   `json:"merged,omitempty"`
	Mergeable           *bool   `json:"mergeable,omitempty"`
	MergeableState      *string `json:"mergeable_state,omitempty"`
	Rebaseable          *bool   `json:"rebaseable,omitempty"`
	MergedBy            *User   `json:"merged_by,omitempty"`
	MergeCommitSHA      *string `json:"merge_commit_sha,omitempty"`
	Comments            *int    `json:"comments,omitempty"`
	Commits             *int    `json:"commits,omitempty"`
	Additions           *int    `json:"additions,omitempty"`
	Deletions           *int    `json:"deletions,omitempty"`
	ChangedFiles        *int    `json:"changed_files,omitempty"`
	MaintainerCanModify *bool   `json:"maintainer_can_modify,omitempty"`
	ReviewComments      *int    `json:"review_comments,omitempty"`

	RequestedTeams []*Team `json:"requested_teams,omitempty"`

	Links *PRLinks           `json:"_links,omitempty"`
	Head  *PullRequestBranch `json:"head,omitempty"`
	Base  *PullRequestBranch `json:"base,omitempty"`

	ActiveLockReason *string `json:"active_lock_reason,omitempty"`
}

func (p PullRequest) String() string { _ = "STUB: not implemented"; return "" }

type PRLink struct {
	HRef *string `json:"href,omitempty"`
}

type PRLinks struct {
	Self           *PRLink `json:"self,omitempty"`
	HTML           *PRLink `json:"html,omitempty"`
	Issue          *PRLink `json:"issue,omitempty"`
	Comments       *PRLink `json:"comments,omitempty"`
	ReviewComments *PRLink `json:"review_comments,omitempty"`
	ReviewComment  *PRLink `json:"review_comment,omitempty"`
	Commits        *PRLink `json:"commits,omitempty"`
	Statuses       *PRLink `json:"statuses,omitempty"`
}

type PullRequestBranch struct {
	Label *string     `json:"label,omitempty"`
	Ref   *string     `json:"ref,omitempty"`
	SHA   *string     `json:"sha,omitempty"`
	Repo  *Repository `json:"repo,omitempty"`
	User  *User       `json:"user,omitempty"`
}

type PullRequestListOptions struct {
	State string `url:"state,omitempty"`

	Head string `url:"head,omitempty"`

	Base string `url:"base,omitempty"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/pulls
func (s *PullRequestsService) List(ctx context.Context, owner, repo string, opts *PullRequestListOptions) ([]*PullRequest, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/commits/{commit_sha}/pulls
func (s *PullRequestsService) ListPullRequestsWithCommit(ctx context.Context, owner, repo, sha string, opts *ListOptions) ([]*PullRequest, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pulls/{pull_number}
func (s *PullRequestsService) Get(ctx context.Context, owner, repo string, number int) (*PullRequest, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pulls/{pull_number}
func (s *PullRequestsService) GetRaw(ctx context.Context, owner, repo string, number int, opts RawOptions) (string, *Response, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

type NewPullRequest struct {
	Title *string `json:"title,omitempty"`

	Head     *string `json:"head,omitempty"`
	HeadRepo *string `json:"head_repo,omitempty"`

	Base                *string `json:"base,omitempty"`
	Body                *string `json:"body,omitempty"`
	Issue               *int    `json:"issue,omitempty"`
	MaintainerCanModify *bool   `json:"maintainer_can_modify,omitempty"`
	Draft               *bool   `json:"draft,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/pulls
func (s *PullRequestsService) Create(ctx context.Context, owner, repo string, body *NewPullRequest) (*PullRequest, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type PullRequestBranchUpdateOptions struct {
	ExpectedHeadSHA *string `json:"expected_head_sha,omitempty"`
}

type PullRequestBranchUpdateResponse struct {
	Message *string `json:"message,omitempty"`
	URL     *string `json:"url,omitempty"`
}

//meta:operation PUT /repos/{owner}/{repo}/pulls/{pull_number}/update-branch
func (s *PullRequestsService) UpdateBranch(ctx context.Context, owner, repo string, number int, body *PullRequestBranchUpdateOptions) (*PullRequestBranchUpdateResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type pullRequestUpdate struct {
	Title               *string `json:"title,omitempty"`
	Body                *string `json:"body,omitempty"`
	State               *string `json:"state,omitempty"`
	Base                *string `json:"base,omitempty"`
	MaintainerCanModify *bool   `json:"maintainer_can_modify,omitempty"`
}

//meta:operation PATCH /repos/{owner}/{repo}/pulls/{pull_number}
func (s *PullRequestsService) Edit(ctx context.Context, owner, repo string, number int, pull *PullRequest) (*PullRequest, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pulls/{pull_number}/commits
func (s *PullRequestsService) ListCommits(ctx context.Context, owner, repo string, number int, opts *ListOptions) ([]*RepositoryCommit, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pulls/{pull_number}/files
func (s *PullRequestsService) ListFiles(ctx context.Context, owner, repo string, number int, opts *ListOptions) ([]*CommitFile, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pulls/{pull_number}/merge
func (s *PullRequestsService) IsMerged(ctx context.Context, owner, repo string, number int) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

type PullRequestMergeResult struct {
	SHA     *string `json:"sha,omitempty"`
	Merged  *bool   `json:"merged,omitempty"`
	Message *string `json:"message,omitempty"`
}

type PullRequestOptions struct {
	CommitTitle string
	SHA         string

	MergeMethod string

	DontDefaultIfBlank bool
}

type pullRequestMergeRequest struct {
	CommitMessage *string `json:"commit_message,omitempty"`
	CommitTitle   string  `json:"commit_title,omitempty"`
	MergeMethod   string  `json:"merge_method,omitempty"`
	SHA           string  `json:"sha,omitempty"`
}

//meta:operation PUT /repos/{owner}/{repo}/pulls/{pull_number}/merge
func (s *PullRequestsService) Merge(ctx context.Context, owner, repo string, number int, commitMessage string, options *PullRequestOptions) (*PullRequestMergeResult, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

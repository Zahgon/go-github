package github

import (
	"context"
	"time"
)

type RepositoryCommit struct {
	NodeID      *string   `json:"node_id,omitempty"`
	SHA         *string   `json:"sha,omitempty"`
	Commit      *Commit   `json:"commit,omitempty"`
	Author      *User     `json:"author,omitempty"`
	Committer   *User     `json:"committer,omitempty"`
	Parents     []*Commit `json:"parents,omitempty"`
	HTMLURL     *string   `json:"html_url,omitempty"`
	URL         *string   `json:"url,omitempty"`
	CommentsURL *string   `json:"comments_url,omitempty"`

	Stats *CommitStats `json:"stats,omitempty"`

	Files []*CommitFile `json:"files,omitempty"`
}

func (r RepositoryCommit) String() string { _ = "STUB: not implemented"; return "" }

type CommitStats struct {
	Additions *int `json:"additions,omitempty"`
	Deletions *int `json:"deletions,omitempty"`
	Total     *int `json:"total,omitempty"`
}

func (c CommitStats) String() string { _ = "STUB: not implemented"; return "" }

type CommitFile struct {
	SHA              *string `json:"sha,omitempty"`
	Filename         *string `json:"filename,omitempty"`
	Additions        *int    `json:"additions,omitempty"`
	Deletions        *int    `json:"deletions,omitempty"`
	Changes          *int    `json:"changes,omitempty"`
	Status           *string `json:"status,omitempty"`
	Patch            *string `json:"patch,omitempty"`
	BlobURL          *string `json:"blob_url,omitempty"`
	RawURL           *string `json:"raw_url,omitempty"`
	ContentsURL      *string `json:"contents_url,omitempty"`
	PreviousFilename *string `json:"previous_filename,omitempty"`
}

func (c CommitFile) String() string { _ = "STUB: not implemented"; return "" }

type CommitsComparison struct {
	BaseCommit      *RepositoryCommit `json:"base_commit,omitempty"`
	MergeBaseCommit *RepositoryCommit `json:"merge_base_commit,omitempty"`

	Status       *string `json:"status,omitempty"`
	AheadBy      *int    `json:"ahead_by,omitempty"`
	BehindBy     *int    `json:"behind_by,omitempty"`
	TotalCommits *int    `json:"total_commits,omitempty"`

	Commits []*RepositoryCommit `json:"commits,omitempty"`

	Files []*CommitFile `json:"files,omitempty"`

	HTMLURL      *string `json:"html_url,omitempty"`
	PermalinkURL *string `json:"permalink_url,omitempty"`
	DiffURL      *string `json:"diff_url,omitempty"`
	PatchURL     *string `json:"patch_url,omitempty"`
	URL          *string `json:"url,omitempty"`
}

func (c CommitsComparison) String() string { _ = "STUB: not implemented"; return "" }

type CommitsListOptions struct {
	SHA string `url:"sha,omitempty"`

	Path string `url:"path,omitempty"`

	Author string `url:"author,omitempty"`

	Since time.Time `url:"since,omitempty"`

	Until time.Time `url:"until,omitempty"`

	ListOptions
}

type BranchCommit struct {
	Name      *string `json:"name,omitempty"`
	Commit    *Commit `json:"commit,omitempty"`
	Protected *bool   `json:"protected,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/commits
func (s *RepositoriesService) ListCommits(ctx context.Context, owner, repo string, opts *CommitsListOptions) ([]*RepositoryCommit, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/commits/{ref}
func (s *RepositoriesService) GetCommit(ctx context.Context, owner, repo, sha string, opts *ListOptions) (*RepositoryCommit, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/commits/{ref}
func (s *RepositoriesService) GetCommitRaw(ctx context.Context, owner, repo, sha string, opts RawOptions) (string, *Response, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/commits/{ref}
func (s *RepositoriesService) GetCommitSHA1(ctx context.Context, owner, repo, ref, lastSHA string) (string, *Response, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/compare/{basehead}
func (s *RepositoriesService) CompareCommits(ctx context.Context, owner, repo, base, head string, opts *ListOptions) (*CommitsComparison, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/compare/{basehead}
func (s *RepositoriesService) CompareCommitsRaw(ctx context.Context, owner, repo, base, head string, opts RawOptions) (string, *Response, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/commits/{commit_sha}/branches-where-head
func (s *RepositoriesService) ListBranchesHeadCommit(ctx context.Context, owner, repo, sha string) ([]*BranchCommit, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

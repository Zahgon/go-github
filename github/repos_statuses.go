package github

import (
	"context"
)

type RepoStatus struct {
	ID     *int64  `json:"id,omitempty"`
	NodeID *string `json:"node_id,omitempty"`
	URL    *string `json:"url,omitempty"`

	State *string `json:"state,omitempty"`

	TargetURL *string `json:"target_url,omitempty"`

	Description *string `json:"description,omitempty"`

	Context *string `json:"context,omitempty"`

	AvatarURL *string `json:"avatar_url,omitempty"`

	Creator   *User      `json:"creator,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
}

func (r RepoStatus) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/commits/{ref}/statuses
func (s *RepositoriesService) ListStatuses(ctx context.Context, owner, repo, ref string, opts *ListOptions) ([]*RepoStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/statuses/{sha}
func (s *RepositoriesService) CreateStatus(ctx context.Context, owner, repo, ref string, status RepoStatus) (*RepoStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CombinedStatus struct {
	State *string `json:"state,omitempty"`

	Name       *string       `json:"name,omitempty"`
	SHA        *string       `json:"sha,omitempty"`
	TotalCount *int          `json:"total_count,omitempty"`
	Statuses   []*RepoStatus `json:"statuses,omitempty"`

	CommitURL     *string `json:"commit_url,omitempty"`
	RepositoryURL *string `json:"repository_url,omitempty"`
}

func (s CombinedStatus) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/commits/{ref}/status
func (s *RepositoriesService) GetCombinedStatus(ctx context.Context, owner, repo, ref string, opts *ListOptions) (*CombinedStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

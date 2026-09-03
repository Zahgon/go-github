package github

import (
	"context"
)

type StarredRepository struct {
	StarredAt  *Timestamp  `json:"starred_at,omitempty"`
	Repository *Repository `json:"repo,omitempty"`
}

type Stargazer struct {
	StarredAt *Timestamp `json:"starred_at,omitempty"`
	User      *User      `json:"user,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/stargazers
func (s *ActivityService) ListStargazers(ctx context.Context, owner, repo string, opts *ListOptions) ([]*Stargazer, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ActivityListStarredOptions struct {
	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	ListOptions
}

//meta:operation GET /user/starred
//meta:operation GET /users/{username}/starred
func (s *ActivityService) ListStarred(ctx context.Context, user string, opts *ActivityListStarredOptions) ([]*StarredRepository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/starred/{owner}/{repo}
func (s *ActivityService) IsStarred(ctx context.Context, owner, repo string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

//meta:operation PUT /user/starred/{owner}/{repo}
func (s *ActivityService) Star(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /user/starred/{owner}/{repo}
func (s *ActivityService) Unstar(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

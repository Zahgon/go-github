package github

import (
	"context"
)

type CreateAutolinkRequest struct {
	KeyPrefix      string `json:"key_prefix"`
	URLTemplate    string `json:"url_template"`
	IsAlphanumeric *bool  `json:"is_alphanumeric,omitempty"`
}

type Autolink struct {
	ID             *int64  `json:"id,omitempty"`
	KeyPrefix      *string `json:"key_prefix,omitempty"`
	URLTemplate    *string `json:"url_template,omitempty"`
	IsAlphanumeric *bool   `json:"is_alphanumeric,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/autolinks
func (s *RepositoriesService) ListAutolinks(ctx context.Context, owner, repo string) ([]*Autolink, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/autolinks
func (s *RepositoriesService) CreateAutolink(ctx context.Context, owner, repo string, body CreateAutolinkRequest) (*Autolink, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/autolinks/{autolink_id}
func (s *RepositoriesService) GetAutolink(ctx context.Context, owner, repo string, id int64) (*Autolink, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/autolinks/{autolink_id}
func (s *RepositoriesService) DeleteAutolink(ctx context.Context, owner, repo string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

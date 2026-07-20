package github

import (
	"context"
)

type AutolinkOptions struct {
	KeyPrefix      *string `json:"key_prefix,omitempty"`
	URLTemplate    *string `json:"url_template,omitempty"`
	IsAlphanumeric *bool   `json:"is_alphanumeric,omitempty"`
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
func (s *RepositoriesService) AddAutolink(ctx context.Context, owner, repo string, body *AutolinkOptions) (*Autolink, *Response, error) {
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

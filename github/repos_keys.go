package github

import (
	"context"
)

//meta:operation GET /repos/{owner}/{repo}/keys
func (s *RepositoriesService) ListKeys(ctx context.Context, owner, repo string, opts *ListOptions) ([]*Key, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/keys/{key_id}
func (s *RepositoriesService) GetKey(ctx context.Context, owner, repo string, id int64) (*Key, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CreateDeployKeyRequest struct {
	Title    *string `json:"title,omitempty"`
	Key      string  `json:"key"`
	ReadOnly *bool   `json:"read_only,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/keys
func (s *RepositoriesService) CreateKey(ctx context.Context, owner, repo string, body CreateDeployKeyRequest) (*Key, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/keys/{key_id}
func (s *RepositoriesService) DeleteKey(ctx context.Context, owner, repo string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

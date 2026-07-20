package github

import (
	"context"
)

type PreReceiveHook struct {
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Enforcement *string `json:"enforcement,omitempty"`
	ConfigURL   *string `json:"configuration_url,omitempty"`
}

func (p PreReceiveHook) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/pre-receive-hooks
func (s *RepositoriesService) ListPreReceiveHooks(ctx context.Context, owner, repo string, opts *ListOptions) ([]*PreReceiveHook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pre-receive-hooks/{pre_receive_hook_id}
func (s *RepositoriesService) GetPreReceiveHook(ctx context.Context, owner, repo string, id int64) (*PreReceiveHook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/pre-receive-hooks/{pre_receive_hook_id}
func (s *RepositoriesService) UpdatePreReceiveHook(ctx context.Context, owner, repo string, id int64, body *PreReceiveHook) (*PreReceiveHook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/pre-receive-hooks/{pre_receive_hook_id}
func (s *RepositoriesService) DeletePreReceiveHook(ctx context.Context, owner, repo string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package github

import (
	"context"
)

type HookConfig struct {
	ContentType *string `json:"content_type,omitempty"`

	InsecureSSL *string `json:"insecure_ssl,omitempty"`
	URL         *string `json:"url,omitempty"`

	Secret *string `json:"secret,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/hooks/{hook_id}/config
func (s *RepositoriesService) GetHookConfiguration(ctx context.Context, owner, repo string, id int64) (*HookConfig, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/hooks/{hook_id}/config
func (s *RepositoriesService) UpdateHookConfiguration(ctx context.Context, owner, repo string, id int64, body HookConfig) (*HookConfig, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

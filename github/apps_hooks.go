package github

import (
	"context"
)

//meta:operation GET /app/hook/config
func (s *AppsService) GetHookConfig(ctx context.Context) (*HookConfig, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /app/hook/config
func (s *AppsService) UpdateHookConfig(ctx context.Context, body HookConfig) (*HookConfig, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
